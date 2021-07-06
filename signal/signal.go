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