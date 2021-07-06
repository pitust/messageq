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

var is_sched_user = false
var kregs = Regs{}


//export putchar
func putchar(c int) int

//export do_isr_handle
func HandleISR(isr uint8, ptr unsafe.Pointer) {
	regs := RegsFromPointer(ptr)
	if regs.CS == 8 && isr == 3 {
		command := regs.RDI
		if command == SERVICE_STORE_REGS {
			args := (*Regs)(unsafe.Pointer(uintptr(regs.RSI)))
			*args = regs
			regs.RAX = SERVICE_OK
			regs.WriteTo(ptr)
			return
		}
		if command == SERVICE_SCHEDULE_USER {
			args := (*Regs)(unsafe.Pointer(uintptr(regs.RSI)))
			kregs = regs
			args.WriteTo(ptr)
			is_sched_user = true
			return
		}
		regs.RAX = SERVICE_ERR_BAD_CALL
		regs.WriteTo(ptr)
		return
	}
	if is_sched_user {
		user_regs := (*Regs)(unsafe.Pointer(uintptr(kregs.RSI)))
		is_sched_user = false
		*user_regs = regs
		regs = kregs
		is_sched_user = false
		regs.RAX = uint64(isr) | SERVICE_IRQ_NOTE | SERVICE_OK
		regs.WriteTo(ptr)
		return
	}
	for {}
}