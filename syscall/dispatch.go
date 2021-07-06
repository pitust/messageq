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
package syscall

import (
	"github.com/pitust/messageq/v2/irq"
	"github.com/pitust/messageq/v2/process"
	"github.com/pitust/messageq/v2/signal"
)

const (
	SYS_SCHED_YIELD = 20
)

//export syscall.dosyscall
func Syscall(ps *process.Process, thr *process.Thread, sysno uint64, regs *irq.Regs) {
	switch sysno {
	case SYS_SCHED_YIELD:
		regs.RAX = 0
		return
	default:
		signal.Raise1(thr.VM(), regs, signal.SIGILL, regs.Err)
		return
	}
}

// this is needed for the above fn to be built in the object file
func Init() {}