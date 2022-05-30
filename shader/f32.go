package shader

import (
	"bytes"
	"encoding/binary"
)

func f32Bytes(values ...float32) []byte {

	/* b := make([]byte, 4*len(values)) */
	/* for i, v := range values { */
	/* u := math.Float32bits(v) */
	/* b[4*i+0] = byte(u >> 0) */
	/* b[4*i+1] = byte(u >> 8) */
	/* b[4*i+2] = byte(u >> 16) */
	/* b[4*i+3] = byte(u >> 24) */
	/* } */

	var buf bytes.Buffer
	for _, v := range values {
		binary.Write(&buf, binary.LittleEndian, v)
	}
	return buf.Bytes()
}

func bytesF32(b []byte) []float32 {
	_ = b[3]
	r := []float32{}

	for i := 0; i < len(b); i += 4 {
		// f := uint32(b[i]) | uint32(b[i+1])<<8 | uint32(b[i+2])<<16 | uint32(b[i+3])<<24
		f := binary.LittleEndian.Uint32(b[i : i+3])
		r = append(r, float32(f))
	}

	return r
}
