%rename cpp_options old_cpp_options

*cpp_options:
-nostdinc %(old_cpp_options) -isystem external/musl_sysroot/include

*cc1:
%(cc1_cpu) -nostdinc -isystem external/musl_sysroot/include

*link_libgcc:
-L .%s -L external/musl_sysroot/lib

*libgcc:
libgcc.a%s %:if-exists(libgcc_eh.a%s)

*startfile:
%{!shared: external/musl_sysroot/lib/Scrt1.o} external/musl_sysroot/lib/crti.o crtbeginS.o%s

*endfile:
crtendS.o%s external/musl_sysroot/lib/crtn.o

*link:
-nostdlib -no-dynamic-linker -static %{rdynamic:-export-dynamic}

*esp_link:


*esp_options:


*esp_cpp_options:


