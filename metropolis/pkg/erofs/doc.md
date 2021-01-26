# EROFS Primer
EROFS is a relatively modern (Linux 5.3+) filesystem optimized for fast read-only use. Similar to squashfs
and cramfs EROFS filesystems have no write support in the kernel and can only be created by external tools.
Both squashfs and cramfs are extremely optimized towards achieving minimal size, to the detriment of performance.
For modern server use both of them are unacceptably slow as they support limited concurrency, make inefficient
use of the page cache and read in weird block sizes. EROFS is designed to replace them on modern, fast hardware
and generally exceeds Ext4 in performance by leveraging the fact that it is read-only. It supports compression,
but only in fixed disk-aligned chunks and using LZ4 for maximum performance.

Sadly the existing tooling to create EROFS filesystems (erofs-utils's mkfs.erofs) can only pack up single
folders which does not work in a build process as it would both require root access to get file ownership
and device nodes correct as well as a complete content copy of all relevant files which is bad as it lies
on the critical path of the image build process. Adopting mkfs.erofs for a spec-driven build process
basically amounts to a rewrite as the "library" part of it also directly reads directories and thus cannot
be used directly.

As reusing the old code proved to be more effort than it's worth, this library was born. Sadly upstream EROFS
has basically no documentation beyond a few trivial diagrams that describe how exactly the filesystem is
constructed. This document holds most knowledge I pried from `mkfs.erofs` and the kernel implementation and
should help people understand the code.

# Blocks
An EROFS filesystem consists of individual blocks, each of them 4096 bytes (4K) long. Each block can either be a metadata or a data block (but it's not possible to know just by looking at a single block). The first block is always a metadata block and the first 1024 bytes of it are occupied by padding. The next 128 bytes are a Superblock structure. The rest (2944 bytes) is available for normal metadata allocation. Blocks are numbered from zero.

# Superblock
As mentionend in the previous section about blocks the superblock is not actually a block in EROFS, but a 128 byte-sized structure 1024 bytes into the first block. Most fields don't need to be set and don't matter. The `BuildTimeSeconds` and `BuildTimeNanoseconds` fields determine the ctime, atime and mtime of all inodes which are in compact structure. This library leaves them at zero which results in all files having a creation time of 1.1.1970 (Unix zero). This is similar to what Bazel does for archives. The only fields which need be filled out are the magic bytes, the block size in bits (only 12 for 1 << 12 = 4096 is supported though) and the root `nid` which points to the inode structure of the root inode (which needs to be a directory for the EROFS to be mountable). The root inode cannot be everywhere on the disk as the integer size of the field is only 16 bits, whereas a normal nid field everywhere else is 32 bits. So in general the root directory immediately follows the superblock and thus has `nid` (1024 + 128)/32 = 36.

# Inodes
Inodes all have a common inode structure which exists in both compact (32 bytes) and extended (64 bytes) form. There's no fundamental difference, the extended form can store more metadata and has a bigger maximum file size (2^64 vs 2^32). All inode structures are aligned to 32 bytes. Through this alignment they are identifiable by a so-called `nid` which is simply their offset in bytes to the start of the filesystem divided by their alignment (32). Certain inodes (inline and compressed, a variant of inline) also store data immediately following the inode. The inode structure and its optional following data are allocated in metadata blocks. If there's no metadata block with enough free bytes to accomodate the inode, a new block is allocated and filled with that inode structure and its following data.

EROFS has three on-disk inode layouts that are in use:

## Plain Inodes
These consist only of a single inode structure (compact, 32 bytes or extended, 64 bytes) in the metadata area, and zero or more filled data blocks (empty inodes are always plain). All data blocks are consecutive and the Union value of the inode contains the block number of the first data block. The `Size` value contains the size of the content in bytes, not including the inode structure itself. The number of data blocks is determined by dividing the `Size` value of the inode by the block size and rounding up (see next paragraph why rounding up is necessary). The data blocks do not need to be adjacent to the metadata block the inode is in.

## Inodes with inline data
These are similar to plain inodes but also work for inode content sizes not neatly divisible by the block size. The leftover data at the end of the inode content that didn't fit into a whole data block is placed in the metadata area directly following the inode itself. How many bytes are appended to the inode is again determined by looking at the inode structure's `Size` value and calculating the remainder when divided by the block size (4096 bytes). The number of blocks is the result of the integer division of these numbers. As with the plain inodes the full blocks don't need to be adjacent to the metadata block.

An inline inode can thus occupy more than a whole metadata block (32 bytes inode + 4095 bytes of data that didn't fit into a full block). This special case is handled by detecting that the inode plus the inline data would exceed a full metadata block (4096) and converting to a plain inode with an additional data block which is zero-padded. This is done specifically when (inode_content_size % 4096) + inode_self_size > 4096. Thus if this special case has happened can also be determined just from the inode size value.

## Compressed inodes
EROFS supports what they call Variable-Length Extents. These are normal plain inodes or inodes with inline data, but instead of the data itself they contain a metadata structure beginning with a `MapHeader` which is mostly there for legacy reasons and always contains the same data. Then follow compressed VLE meta blocks, which contain either 2 or 16 packed 14 bit integers and a on-disk block number. For alignment reasons the first 6 VLE meta integers are always packed into the 2 integer structures. All following complete blocks of 16 VLE meta integers get packed into 16 byte together with their starting block number. Anything that's left over gets once again packed into 2 integer structures. Each integer in this compressed sequence of 14 bit integers represents 4K of uncompressed data. So a file which has an uncompressed size of 4MiB needs 1000 of these integers to be represented, independent of how well it compresses.

Note that VLE meta blocks are treated as content of plain or inline inodes. So if they exceed the maximum inline inode size there will be blocks allocated just for storing VLE meta blocks.

These VLE meta integers integers are divided into 12 lower bits and 2 upper bits. The upper bits determine what the lower 12 bits represent and also how this 4K block of uncompressed data is represented. There's three types: PLAIN, HEAD and NONHEAD. PLAIN means no compression, the block is stored as-is on disk. HEAD means this block is the start of a compressed cluster. Its 12 lower bits represent the offset of the decompressed data with regards to the uncompressed 4K block boundary. NONHEAD means this block is part of the same on-disk block as the last HEAD block when uncompressed. Its lower 12 bits represent the number of blocks until the next HEAD or PLAIN block unless at the end of a VLE meta block (2 or 16 integers), then they represent the distance from the last HEAD or PLAIN block.

Only PLAIN and HEAD blocks have actual on-disk blocks of uncompressed and compressed data respectively. NONHEAD blocks only exist to represent data that's expanded by decompressing. Thus the on-disk block number of a PLAIN or HEAD block can be determined by looking at the on-disk block number of the VLE meta block and incrementing by one for each PLAIN or HEAD block in it. So all data blocks referenced inside one VLE meta block need to be consecutive (but not adjacent to the location of the VLE meta blocks themselves).

# Unix file types and inode layout

## Directories
Directories are either plain or inline inodes. They have content which consists of 12 byte dirent structures (directoryEntryRaw). These dirent structures contain a `nid` (see Inodes) pointing to the inode structure that represents that child of the directory, a name offset and a file type. The file type is redundant (it is also stored in the child's inode) but needs to be set. Directly following the dirent structures are all names of the children. They are not terminated or aligned. The name offset stored in a dirent is relative to the start of the inode content and marks the first byte of the name for that child. The end can be determined by the name offset of the next dirent or the total size of the inode if it is the last child.

Directories always contain the `.` and `..` children, which need to point to itself and the parent inodes respectively, with the exception that the root directory's parent is defined to be itself. The individual dirents are always sorted according to their name interpreted as bytes to allow for binary searching.

## Symbolic links
Symbolic links are inline inodes. They have their literal target path as content.

## Device nodes
Device nodes are always plain inodes. Instead of a content size they have a `dev_t` integer in the `Union` inode struct value encoding the major and minor numbers. The type of device inode (block or character) is determined by the high bits of the Mode value as in standard Unix.


## Regular files
Regular files can be any of the three plain, inline or compressed inodes. The inode content is the content of the file.

## Others
FIFOs and Sockets are plain inodes with no content and no special fields. They will also be seldomly used in an EROFS.
