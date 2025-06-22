package freetype

import "modernc.org/libc"

func malloc(n uint64) (r uintptr) {
	return libc.Xmalloc(nil, n)
}
