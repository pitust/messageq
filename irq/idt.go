package irq

import (
	"unsafe"
)

func ep(e error) {
	if e != nil {
		panic(e)
	}
}
func CreateIDT(tgd uintptr) {
	i := 0
	for i < 256 {
		ptr := GetISR(uint64(i))
		*(*uint16)(unsafe.Pointer(tgd + 0x00)) = uint16(ptr)
		*(*uint16)(unsafe.Pointer(tgd + 0x02)) = uint16(8)
		*(*uint8)(unsafe.Pointer(tgd + 0x04)) = uint8(0)
		*(*uint8)(unsafe.Pointer(tgd + 0x05)) = uint8(0x8e)
		*(*uint16)(unsafe.Pointer(tgd + 0x06)) = uint16(ptr >> 16)
		*(*uintptr)(unsafe.Pointer(tgd + 0x08)) = ptr >> 32
		tgd += 16
		i++
	}
}