package signal

import (
	"github.com/pitust/messageq/v2/irq"
	"github.com/pitust/messageq/v2/vm"
)

type Signal uint64

const (
	SIGILL = Signal(1)
	SIGSEGV = Signal(2)
)

func Raise2(vm *vm.UserVM, regs *irq.Regs, sig Signal, ec0 uint64, ec1 uint64) {
	panic("Raise2: todo")
}
func Raise1(vm *vm.UserVM, regs *irq.Regs, sig Signal, ec0 uint64) {
	Raise2(vm, regs, sig, ec0, 0)
}
func Raise(vm *vm.UserVM, regs *irq.Regs, sig Signal) {
	Raise2(vm, regs, sig, 0, 0)
}