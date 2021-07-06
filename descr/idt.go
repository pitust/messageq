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

import (
	"unsafe"
	"github.com/pitust/messageq/v2/irq"
)

var idt [4096]byte

func InitIDT() {
	idtptr := uintptr(unsafe.Pointer(&idt[0]))
	irq.CreateIDT(idtptr)
	var idtr Descriptor
	WriteDescriptor(&idtr, idtptr, 4096)
	idtpv := uintptr(unsafe.Pointer(&idtr))
	irq.LoadIDT(idtpv)
}