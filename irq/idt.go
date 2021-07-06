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