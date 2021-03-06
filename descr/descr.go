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

// Descriptors (TSS, IDT, GDT)
package descr
type Descriptor struct {
	low uintptr
	hi uintptr
}
func WriteDescriptor(dptr *Descriptor, addr uintptr, limit uint16) {
	dptr.low = (addr << 16) | uintptr(limit)
	dptr.hi = addr >> 48
}