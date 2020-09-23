// gooz is an interface to ooz, an open-source Kraken / Mermaid / Selkie /
// Leviathan / LZNA / BitKnit decompressor.
package gooz

// #cgo CFLAGS: -Wall
// #include "ooz.h"
import "C"
import "fmt"

// Decompress decompresses a buffer of compressed data and returns the
// decompressed data as a byte slice. The length of the decompressed data must
// be provided; this is often stored alongside the compressed data, or in a
// header.
func Decompress(data []byte, rawsize int) ([]byte, error) {
	i_buf := C.CBytes(data)
	o_buf := C.malloc(C.size_t(rawsize + 64)) // decoder is sloppy
	defer C.free(i_buf)
	defer C.free(o_buf)

	r_sz := C.Kraken_Decompress(
		i_buf, C.size_t(len(data)),
		o_buf, C.size_t(rawsize),
	)

	if int(r_sz) < 0 {
		return nil, fmt.Errorf("failed: unspecified error")
	}

	if int(r_sz) != rawsize {
		return nil, fmt.Errorf("only decompressed %d bytes", int(r_sz), rawsize)
	}

	return C.GoBytes(o_buf, C.int(rawsize)), nil
}
