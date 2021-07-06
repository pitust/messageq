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

// LAPIC configuration
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
	apicBase := msr.Read(/* IA32_APIC_BASE */ 0x1b)
	lapic := uintptr(apicBase) & 0xffff_ffff_ffff_f000
	*misc.Uint32Ptr(lapic + 0x380) = 0xfffff
}
func EOI() {
	apicBase := msr.Read(/* IA32_APIC_BASE */ 0x1b)
	lapic := uintptr(apicBase) & 0xffff_ffff_ffff_f000
	*misc.Uint32Ptr(lapic + 0xB0) = 0
}