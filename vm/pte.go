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

// virtual memory subsystem
package vm

import (
 	"unsafe"

	"github.com/pitust/messageq/v2/misc"
)

//export memset
func memset(p unsafe.Pointer, v uint8, len uintptr)

func getPTEPointer(va uintptr) uintptr {
    page_table := misc.CR3()
    control_pte := uintptr(0)
    va_val := va & 0x000f_ffff_ffff_f000
    offsets := []uintptr{
        ((((va_val >> 12) >> 9) >> 9) >> 9) & 0x1ff,
        (((va_val >> 12) >> 9) >> 9) & 0x1ff,
        ((va_val >> 12) >> 9) & 0x1ff,
        (va_val >> 12) & 0x1ff,
	}
    i := -1
    for _, idx := range misc.Iter(4) {
        key := offsets[idx]
        i++
        ptk := *(*uint64)(unsafe.Pointer(page_table + key * 8))
        if (ptk & 4) == 0 && (ptk & 1) == 1 && i != 3 {
            // fuck bootboot
            println("fuck bootboot: fixing bad bootboot pte code, and setting User Accessible on this god fucking damn page")
            *(*uint64)(unsafe.Pointer(page_table + key * 8)) |= 0x04
        }
        if (ptk & 0x80) == 0x80 {
            panic("Unable to map (huge page in the way)")
        }
        if (ptk & 1 == 0) && i != 3 {
			new_page_table := PhysAlloc()
            *(*uintptr)(unsafe.Pointer(page_table + key * 8)) = 0x07 | uintptr(new_page_table)
        }
        control_pte = page_table + key * 8
        page_table = (*(*uintptr)(unsafe.Pointer(page_table + key * 8))) & 0x000f_ffff_ffff_f000
    }

    return control_pte
}