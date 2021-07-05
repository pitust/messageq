package misc

import "unsafe"

func Uint8Ptr(ptr uintptr) *uint8 {
	return (*uint8)(unsafe.Pointer(ptr))
}

func Uint16Ptr(ptr uintptr) *uint16 {
	return (*uint16)(unsafe.Pointer(ptr))
}

func Uint32Ptr(ptr uintptr) *uint32 {
	return (*uint32)(unsafe.Pointer(ptr))
}

func Uint64Ptr(ptr uintptr) *uint64 {
	return (*uint64)(unsafe.Pointer(ptr))
}

func UsizePtr(ptr uintptr) *uintptr {
	return (*uintptr)(unsafe.Pointer(ptr))
}