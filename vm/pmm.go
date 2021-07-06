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
package vm

import "unsafe"

type MemoryRegion struct {
	Base uintptr
	Len  uintptr
}

var mem []MemoryRegion = []MemoryRegion{}

func PhysAlloc() unsafe.Pointer {
	if len(mem) == 0 {
		panic("OOM")
	}
	tregion := &mem[0]
	for tregion.Len < 4096 {
		mem = mem[1:]
		tregion = &mem[0]
	}
	tgbase := unsafe.Pointer(tregion.Base)
	if tregion.Len > 4096 {
		tregion.Base += 4096
		tregion.Len -= 4096
		return tgbase
	}
	mem = mem[1:]
	return tgbase
}

func PhysFree(ptr unsafe.Pointer) {
	PhysAddRegion(MemoryRegion{Base: uintptr(ptr), Len: 4096})
}

func PhysAddRegion(r MemoryRegion) {
	mem = append(mem, r)
}
