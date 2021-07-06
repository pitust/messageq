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

import "github.com/pitust/messageq/v2/misc"

type Flags uintptr

const (
	FL_NONE = Flags(0x00)
	/// Implies FL_READ
	FL_WRITE = Flags(0x02) | FL_READ
	FL_USER = Flags(0x04)
	FL_READ = Flags(0x01)
)

/// Maps `virt` to `phys` (so *virt == *phys and tlblookup(virt) == phys), using flags from `Flags`
func Map(phys, virt uintptr, flags Flags) {
	pteptr := getPTEPointer(virt)
	*misc.UsizePtr(pteptr) = uintptr(flags) | uintptr(phys)
}

/// (internal) do the unmap of `virt` without checks
func doUnmap(virt, pteptr uintptr) {
	*misc.UsizePtr(pteptr) = 0
}

/// Unmaps a user page `virt`
func Unmap(virt uintptr) {
	pteptr := getPTEPointer(virt)
	pv := *misc.UsizePtr(pteptr)
	if pv & uintptr(FL_READ) == 0 {
		panic("Page is not mapped, cannot unmap more! (consider using EnsureUnmapped)")
	}
	if pv & uintptr(FL_USER) == 0 {
		panic("Cannot unmap a kernel page with `Unmap`, consider using KeUnmap")
	}
	doUnmap(virt, pteptr)
}

/// Unmaps kernel page at `virt`
func KeUnmap(virt uintptr) {
	pteptr := getPTEPointer(virt)
	pv := *misc.UsizePtr(pteptr)
	if pv & uintptr(FL_READ) == 0 {
		panic("Page is not mapped, cannot unmap more! (consider using `KeEnsureUnmapped`)")
	}
	if pv & uintptr(FL_USER) == uintptr(FL_USER) {
		panic("Cannot unmap a user page with `KeUnmap`, consider using `Unmap`")
	}
	doUnmap(virt, pteptr)
}

/// Ensures `virt` is unmapped (for user pages)
func EnsureUnmapped(virt uintptr) {
	pteptr := getPTEPointer(virt)
	pv := *misc.UsizePtr(pteptr)
	if pv & uintptr(FL_USER) == 0 {
		panic("Cannot unmap a kernel page with `EnsureUnmapped`, consider using `EnsureKeUnmapped`")
	}
	doUnmap(virt, pteptr)
}

/// Ensures `virt` is unmapped (for kernel pages)
func EnsureKeUnmapped(virt uintptr) {
	pteptr := getPTEPointer(virt)
	pv := *misc.UsizePtr(pteptr)
	if pv & uintptr(FL_USER) == uintptr(FL_USER) {
		panic("Cannot unmap a kernel page with `EnsureKeUnmapped`, consider using `EnsureUnmapped`")
	}
	doUnmap(virt, pteptr)
}