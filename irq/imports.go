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

// IRQ handling
package irq

import "unsafe"

//export do_int3
func KSVC_u64(rdi uint64, rsi uint64) uint64

//export do_int3
func KSVC_uintptr(rdi uint64, rsi uintptr) uint64

//export do_int3
func KSVC_unsafeptr(rdi uint64, rsi unsafe.Pointer) uint64

//export do_int3
func KSVC_regsptr(rdi uint64, rsi *Regs) uint64

//export do_int3
func KSVC_returns_unsafeptr(rdi uint64) unsafe.Pointer

//export get_isr
func GetISR(isr uint64) uintptr

//export do_lidt
func LoadIDT(idtr uintptr)