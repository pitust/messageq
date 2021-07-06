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

// virtual memory subsystem
package vm

import "github.com/pitust/messageq/v2/misc"

type UserVM struct {
	vm [256]uintptr
}

func (this *UserVM) Clone() UserVM {
	return UserVM{
		vm: this.vm,
	}
}

func (this *UserVM) Use() {
	page_table := misc.CR3()
	for i := range misc.Iter(256) {
		*misc.UsizePtr(page_table + (uintptr(i) << 3)) = this.vm[i]
	}
}
func (this *UserVM) Save() {
	page_table := misc.CR3()
	for i := range misc.Iter(256) {
		this.vm[i] = *misc.UsizePtr(page_table + (uintptr(i) << 3))
	}
}
func (this *UserVM) With(f func()) {
	pvm := CaptureUserVM()
	this.Use()
	f()
	this.Save()
	pvm.Use()
}

func CaptureUserVM() UserVM {
	var vm [256]uintptr
	page_table := misc.CR3()
	for i := range misc.Iter(256) {
		vm[i] = *misc.UsizePtr(page_table + (uintptr(i) << 3))
	}
	return UserVM{vm}
}

var KernelVM UserVM