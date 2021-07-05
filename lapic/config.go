package lapic

import (
	"unsafe"

	"github.com/pitust/messageq/v2/misc"
	"github.com/pitust/messageq/v2/msr"
)

func Config() {
	println("Configuring the LAPIC!")
	apicBase := msr.Read(/* IA32_APIC_BASE */ 0x1b)
	lapic := uintptr(apicBase) & 0xffff_ffff_ffff_f000
	println("LAPIC @", unsafe.Pointer(lapic))
	sivr := misc.Uint32Ptr(lapic + 0xF0)
	*sivr |= 0x100
	*misc.Uint32Ptr(lapic + 0x320) = 0x20
	*misc.Uint32Ptr(lapic + 0x3E0) = 0x03
	println("LAPIC enabled!")
}
func StartDeadline() {
	println("Starting deadline...")
	apicBase := msr.Read(/* IA32_APIC_BASE */ 0x1b)
	lapic := uintptr(apicBase) & 0xffff_ffff_ffff_f000
	*misc.Uint32Ptr(lapic + 0x380) = 0xfff
}