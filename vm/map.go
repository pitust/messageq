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