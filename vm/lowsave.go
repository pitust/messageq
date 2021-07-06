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