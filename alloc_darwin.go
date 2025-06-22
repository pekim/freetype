package freetype

import (
	"modernc.org/libc"
	"modernc.org/libc/sys/types"
)

func malloc(n uint64) (r uintptr) {
	return libc.Xmalloc(nil, types.Size_t(n))
}
