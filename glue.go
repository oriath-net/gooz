// gooz is an interface to ooz, an open-source Kraken / Mermaid / Selkie /
// Leviathan / LZNA / BitKnit decompressor.
package gooz

// #cgo CFLAGS: -Wall
// #include "ooz.h"
import "C"
import "fmt"

// Decompress behaves similarly to copy(), but passes the data through the ooz
// decompressor.
//
// The size of the output buffer is significant to the decompressor.
func Decompress(in []byte, out []byte) (int, error) {
	i_buf := C.CBytes(in)
	o_buf := C.malloc(C.size_t(len(out) + 64)) // decoder is sloppy
	defer C.free(i_buf)
	defer C.free(o_buf)

	r_sz := C.Kraken_Decompress(
		i_buf, C.size_t(len(in)),
		o_buf, C.size_t(len(out)),
	)

	if int(r_sz) < 0 {
		return 0, fmt.Errorf("unspecified error")
	}

	if int(r_sz) != len(out) {
		return 0, fmt.Errorf("only decompressed %d/%d bytes", int(r_sz), len(out))
	}

	return copy(out, C.GoBytes(o_buf, C.int(len(out)))), nil
}
