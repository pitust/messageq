package bootboot

import "C"

import (
	"unsafe"
	"github.com/pitust/messageq/v2/misc"
)


type MemoryMapEntry struct {
	Base uintptr
	Len uintptr
	Ty uint8
}
type BootBoot struct {}

//go:extern bootboot
var bootboot [0]uint8

func getBootbootAddress() uintptr {
	return uintptr(unsafe.Pointer(&bootboot))
}

func Get() BootBoot {
	return BootBoot{}
}

func (self BootBoot) Size() uint32 {
	return *misc.Uint32Ptr(getBootbootAddress() + 0x04)
}

func (self BootBoot) EntryCount() uint32 {
	return (self.Size() - 128) / 16
}

func (self BootBoot) GetEntry(idx uint32) MemoryMapEntry {
	lattr := *misc.UsizePtr(getBootbootAddress() + 128 + 8 + uintptr(idx) << 4)
	return MemoryMapEntry{
		Base: *misc.UsizePtr(getBootbootAddress() + 128 + uintptr(idx) << 4),
		Len: lattr & 0xffff_ffff_ffff_fff0,
		Ty: uint8(lattr & 0x0f),
	}
}

func (self BootBoot) InitrdPointer() uint64 {
	return *misc.Uint64Ptr(getBootbootAddress() + 0x18)
}

func (self BootBoot) InitrdSize() uint64 {
	return *misc.Uint64Ptr(getBootbootAddress() + 0x20)
}
