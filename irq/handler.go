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
	println("isr", isr)
	println("err", regs.Err)
	println("rip", unsafe.Pointer(uintptr(regs.RIP)))
	if regs.CS == 8 && isr == 3 {
		println("Entered a kSVC ISR!")
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
	for {
		putchar(0x66)
	}
}