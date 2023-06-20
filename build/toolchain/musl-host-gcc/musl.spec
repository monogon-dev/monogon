%rename cpp_options old_cpp_options

*cpp_options:
-nostdinc %(old_cpp_options) -isystem $SYSROOT/include

*cc1:
%(cc1_cpu) -nostdinc -isystem $SYSROOT/include

*link_libgcc:
-L .%s -L %R/lib

*libgcc:
libgcc.a%s %:if-exists(libgcc_eh.a%s)

*startfile:
%{static-pie: %R/lib/rcrt1.o; !shared: %R/lib/Scrt1.o} %R/lib/crti.o crtbeginS.o%s

*endfile:
crtendS.o%s %R/lib/crtn.o

*link:
%{static-pie: -pie} -no-dynamic-linker -nostdlib -static %{rdynamic:-export-dynamic}

*esp_link:


*esp_options:


*esp_cpp_options:


