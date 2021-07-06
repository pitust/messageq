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

// The kernel entrypoint
package main

import (
	"github.com/pitust/messageq/v2/bootboot"
	"github.com/pitust/messageq/v2/descr"
	"github.com/pitust/messageq/v2/irq"
	"github.com/pitust/messageq/v2/lapic"
	"github.com/pitust/messageq/v2/misc"
	"github.com/pitust/messageq/v2/process"
	"github.com/pitust/messageq/v2/syscall"
	"github.com/pitust/messageq/v2/vm"
)

func main() {
	println("Hello, world from messageQ kernel!")

	descr.InitGDT()
	descr.InitIDT()
	syscall.Init()

	r := irq.Regs{}
	irq.KSVC_regsptr(irq.SERVICE_STORE_REGS, &r)
	r.CS = descr.GDT_USER_CODE
	r.SS = descr.GDT_USER_DATA

	const MMAP_LOW = 0x0000_020_0000_0000

	for _, i := range misc.Iter(int(bootboot.Get().EntryCount())) {
		ent := bootboot.Get().GetEntry(uint32(i))
		offset := 0x1000 - (ent.Base & 0x0fff)
		if offset == 0x1000 {
			offset = 0
		}
		ent.Base += offset
		if ent.Len < offset {
			continue
		}
		ent.Len -= offset
		if ent.Ty == 1 {
			vm.PhysAddRegion(vm.MemoryRegion{Base: ent.Base, Len: ent.Len})
		}

	}
	lapic.Config()
	vm.KernelVM = vm.CaptureUserVM()

	proc := process.CreateProcess()
	thr := proc.CreateThread()
	thr.With(func() {
		page := vm.PhysAlloc()
		vm.Map(uintptr(page), MMAP_LOW, vm.FL_USER | vm.FL_READ)
		*misc.Uint16Ptr(MMAP_LOW) = 0x14cd
		*misc.Uint16Ptr(MMAP_LOW + 2) = 0xfeeb
	})
	regs := thr.Regs()
	regs.RIP = MMAP_LOW
	process.SchedLoop()
	process.SchedLoop()
	process.SchedLoop()
	
	println("Done!")
	
	for {
	}
}
