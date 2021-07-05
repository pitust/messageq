package irq

import "unsafe"

/// The register state of an irq
type Regs struct {
	R15   uint64
	R14   uint64
	R13   uint64
	R12   uint64
	R11   uint64
	R10   uint64
	R9    uint64
	R8    uint64
	RDI   uint64
	RSI   uint64
	RDX   uint64
	RCX   uint64
	RBX   uint64
	RAX   uint64
	RBP   uint64
	Err   uint64
	RIP   uint64
	CS    uint64
	Flgs uint64
	Stack   uint64
	SS    uint64
}

func RegsFromPointer(p unsafe.Pointer) Regs {
	return Regs{
		R15:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x00))),
		R14:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x08))),
		R13:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x10))),
		R12:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x18))),
		R11:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x20))),
		R10:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x28))),
		R9:    *((*uint64)(unsafe.Pointer(uintptr(p) + 0x30))),
		R8:    *((*uint64)(unsafe.Pointer(uintptr(p) + 0x38))),
		RDI:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x40))),
		RSI:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x48))),
		RDX:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x50))),
		RCX:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x58))),
		RBX:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x60))),
		RAX:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x68))),
		RBP:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x70))),
		Err:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x78))),
		RIP:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x80))),
		CS:    *((*uint64)(unsafe.Pointer(uintptr(p) + 0x88))),
		Flgs: *((*uint64)(unsafe.Pointer(uintptr(p) + 0x90))),
		Stack:   *((*uint64)(unsafe.Pointer(uintptr(p) + 0x98))),
		SS:    *((*uint64)(unsafe.Pointer(uintptr(p) + 0xa0))),
	}
}

func (self *Regs) WriteTo(p unsafe.Pointer) {
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x00)) = self.R15
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x08)) = self.R14
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x10)) = self.R13
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x18)) = self.R12
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x20)) = self.R11
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x28)) = self.R10
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x30)) = self.R9
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x38)) = self.R8
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x40)) = self.RDI
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x48)) = self.RSI
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x50)) = self.RDX
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x58)) = self.RCX
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x60)) = self.RBX
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x68)) = self.RAX
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x70)) = self.RBP
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x78)) = self.Err
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x80)) = self.RIP
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x88)) = self.CS
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x90)) = self.Flgs
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0x98)) = self.Stack
	*(*uint64)(unsafe.Pointer(uintptr(p) + 0xa0)) = self.SS
}