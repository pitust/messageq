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