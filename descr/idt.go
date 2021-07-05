package descr

import (
	"unsafe"
	"github.com/pitust/messageq/v2/irq"
)

var idt [4096]byte

func InitIDT() {
	idtptr := uintptr(unsafe.Pointer(&idt[0]))
	irq.CreateIDT(idtptr)
	var idtr Descriptor
	WriteDescriptor(&idtr, idtptr, 4096)
	idtpv := uintptr(unsafe.Pointer(&idtr))
	irq.LoadIDT(idtpv)
}