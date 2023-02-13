// Package ref provides the reference implementation for kernel command line
// parsing as present in the Linux kernel. This is a separate package and
// not part of the bootparam tests because Go does not let you use cgo in
// tests.
package ref

// Reference implementation from the kernel

/*
#include <stdlib.h>
#include <ctype.h>
#include <stddef.h>

#define _U	0x01
#define _L	0x02
#define _D	0x04
#define _C	0x08
#define _P	0x10
#define _S	0x20
#define _X	0x40
#define _SP	0x80

#define __ismask(x) (_ctype[(int)(unsigned char)(x)])
#define kisspace(c)	((__ismask(c)&(_S)) != 0)

const unsigned char _ctype[] = {
_C,_C,_C,_C,_C,_C,_C,_C,
_C,_C|_S,_C|_S,_C|_S,_C|_S,_C|_S,_C,_C,
_C,_C,_C,_C,_C,_C,_C,_C,
_C,_C,_C,_C,_C,_C,_C,_C,
_S|_SP,_P,_P,_P,_P,_P,_P,_P,
_P,_P,_P,_P,_P,_P,_P,_P,
_D,_D,_D,_D,_D,_D,_D,_D,
_D,_D,_P,_P,_P,_P,_P,_P,
_P,_U|_X,_U|_X,_U|_X,_U|_X,_U|_X,_U|_X,_U,
_U,_U,_U,_U,_U,_U,_U,_U,
_U,_U,_U,_U,_U,_U,_U,_U,
_U,_U,_U,_P,_P,_P,_P,_P,
_P,_L|_X,_L|_X,_L|_X,_L|_X,_L|_X,_L|_X,_L,
_L,_L,_L,_L,_L,_L,_L,_L,
_L,_L,_L,_L,_L,_L,_L,_L,
_L,_L,_L,_P,_P,_P,_P,_C,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
_S|_SP,_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,
_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,_P,
_U,_U,_U,_U,_U,_U,_U,_U,_U,_U,_U,_U,_U,_U,_U,_U,
_U,_U,_U,_U,_U,_U,_U,_P,_U,_U,_U,_U,_U,_U,_U,_L,
_L,_L,_L,_L,_L,_L,_L,_L,_L,_L,_L,_L,_L,_L,_L,_L,
_L,_L,_L,_L,_L,_L,_L,_P,_L,_L,_L,_L,_L,_L,_L,_L};



char *skip_spaces(const char *str)
{
	while (kisspace(*str))
		++str;
	return (char *)str;
}


// * Parse a string to get a param value pair.
// * You can use " around spaces, but can't escape ".
// * Hyphens and underscores equivalent in parameter names.
 char *next_arg(char *args, char **param, char **val)
 {
	 unsigned int i, equals = 0;
	 int in_quote = 0, quoted = 0;

	 if (*args == '"') {
		 args++;
		 in_quote = 1;
		 quoted = 1;
	 }

	 for (i = 0; args[i]; i++) {
		 if (kisspace(args[i]) && !in_quote)
			 break;
		 if (equals == 0) {
			 if (args[i] == '=')
				 equals = i;
		 }
		 if (args[i] == '"')
			 in_quote = !in_quote;
	 }

	 *param = args;
	 if (!equals)
		 *val = NULL;
	 else {
		 args[equals] = '\0';
		 *val = args + equals + 1;

		 // Don't include quotes in value.
		 if (**val == '"') {
			 (*val)++;
			 if (args[i-1] == '"')
				 args[i-1] = '\0';
		 }
	 }
	 if (quoted && i > 0 && args[i-1] == '"')
		 args[i-1] = '\0';

	 if (args[i]) {
		 args[i] = '\0';
		 args += i + 1;
	 } else
		 args += i;

	 // Chew up trailing spaces.
	 return skip_spaces(args);
 }
*/
import "C"
import (
	"unsafe"

	"source.monogon.dev/metropolis/pkg/bootparam"
)

func Parse(str string) (params bootparam.Params, rest string) {
	cs := C.CString(bootparam.TrimLeftSpace(str))
	csAllocPtr := cs
	var param, val *C.char
	for *cs != 0 {
		var p bootparam.Param
		cs = C.next_arg(cs, &param, &val)
		p.Param = C.GoString(param)
		if val != nil {
			p.Value = C.GoString(val)
		}
		if p.Param == "--" {
			rest = C.GoString(cs)
			return
		}
		params = append(params, p)
	}
	C.free(unsafe.Pointer(csAllocPtr))
	return
}
