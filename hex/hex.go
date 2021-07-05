package hex

func Encode(buf []byte) string {
	s := ""
	for _, b := range buf {
		s += string("0123456789abcdef"[b >> 4])
		s += string("0123456789abcdef"[b & 0x0f])
	}
	return s
}