// The messageQ operating system
// Copyright (C) 2021 pitust

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// bootboot-related functions
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
