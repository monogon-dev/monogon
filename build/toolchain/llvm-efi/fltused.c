// This is a marker symbol emitted by MSVC-ABI compatible compilers. Its presence indicates that the linked binary
// contains instructions working with floating-point registers. Since we do not have a standard library which consumes
// it we can just define it as zero.
// See https://github.com/rust-lang/rust/issues/62785#issuecomment-531186089 for more discussion.
// Since building static libraries is not possible with Bazel this is compiled and checked in.
int _fltused __attribute__((used)) = 0;