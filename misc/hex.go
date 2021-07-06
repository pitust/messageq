package misc

import "strconv"

func Hex(i uint64) string {
	return "0x" + strconv.FormatUint(i, 16)
}