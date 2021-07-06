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
package descr

import (
	"unsafe"
	"github.com/pitust/messageq/v2/misc"
)

const (
	GDT_NULL = 0x00
	GDT_KERNEL_CODE = 0x08
	GDT_KERNEL_DATA = 0x10
	GDT_USER_CODE = 0x18 | 3
	GDT_USER_DATA = 0x20 | 3
)

//export get_int_stack
func getIntStack() uintptr

var gdt [0x38]byte
var tss [0x6b]byte
func bits(shiftup, shiftdown, mask, val uint64) uint64 {
    return ((val >> (shiftdown - mask)) & ((1 << mask) - 1)) << shiftup
}
func InitGDT() {

	gdtptr := uintptr(unsafe.Pointer(&gdt[0]))
	tssptr := uint64(uintptr(unsafe.Pointer(&tss[0])))
	tss[0x66] = 13
	*misc.UsizePtr(uintptr(unsafe.Pointer(&tss[4]))) = getIntStack()
	*(*uint64)(unsafe.Pointer(gdtptr + 0x00)) = 0x0000000000000000
	*(*uint64)(unsafe.Pointer(gdtptr + 0x08)) = 0x00af9b000000ffff
	*(*uint64)(unsafe.Pointer(gdtptr + 0x10)) = 0x00af93000000ffff
	*(*uint64)(unsafe.Pointer(gdtptr + 0x18)) = 0x00affb000000ffff
	*(*uint64)(unsafe.Pointer(gdtptr + 0x20)) = 0x00aff3000000ffff 
	*(*uint64)(unsafe.Pointer(gdtptr + 0x28)) = bits(16, 24, 24, tssptr) | bits(56, 32, 8, tssptr) | (103 & 0xff) | ((0b1001) << 40) | ((1) << 47)
	*(*uint64)(unsafe.Pointer(gdtptr + 0x30)) = tssptr >> 32
	var gdtr Descriptor
	WriteDescriptor(&gdtr, gdtptr, 0x38)
	misc.LoadGDT(uintptr(unsafe.Pointer(&gdtr)))
}