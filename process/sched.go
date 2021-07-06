package process

import (
	"github.com/pitust/messageq/v2/descr"
	"github.com/pitust/messageq/v2/irq"
	"github.com/pitust/messageq/v2/lapic"
	"github.com/pitust/messageq/v2/misc"
	"github.com/pitust/messageq/v2/signal"
)

//export syscall.dosyscall
func syscall(ps *Process, thr *Thread, sysno uint64, regs *irq.Regs)


func SchedLoop() {
	for _, proc := range ProcessTable {
		for _, thr := range proc.threads {
			thr.proc.vm.With(func() {
				thr.regs.Flgs |= /* IF */ 0x200
				if thr.regs.CS != descr.GDT_USER_CODE {
					panic("bad CS (attack prevented?)")
				}
				// lapic.EOI()
				lapic.StartDeadline()
				rv := irq.KSVC_regsptr(irq.SERVICE_SCHEDULE_USER, &thr.regs)
				if rv&irq.SERVICE_ERR == irq.SERVICE_ERR {
					panic("Error jumping to userland")
				}
				if rv&irq.SERVICE_IRQ_NOTE == irq.SERVICE_IRQ_NOTE {
					isr := rv & 0xff
					println("isr =", misc.Hex(isr))
					println("err =", misc.Hex(thr.regs.Err))
					if isr > 0x20 {
						// preempted!
						println("Thread preempted")
						if isr == 0x20 {
							return
						}
						panic("Unknown ISR")
					} else if isr == 0x0d && thr.regs.Err&3 == 2 {
						println("syscall #", misc.Hex(thr.regs.Err>>4))
						syscall(proc, thr, thr.regs.Err>>4, &thr.regs)
						thr.regs.RIP += 2
					} else {
						if isr == 0xe {
							signal.Raise2(&thr.proc.vm, &thr.regs, signal.SIGSEGV, thr.regs.Err, uint64(misc.CR2()))
						} else {
							panic("Cannot ISR -> Signal")
						}
					}
				} else {
					panic("Unknown exit cause")
				}
			})
		}
	}
}