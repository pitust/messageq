package irq

import "unsafe"

//export do_int3
func KSVC_u64(rdi uint64, rsi uint64) uint64

//export do_int3
func KSVC_uintptr(rdi uint64, rsi uintptr) uint64

//export do_int3
func KSVC_unsafeptr(rdi uint64, rsi unsafe.Pointer) uint64

//export do_int3
func KSVC_regsptr(rdi uint64, rsi *Regs) uint64

//export do_int3
func KSVC_returns_unsafeptr(rdi uint64) unsafe.Pointer

//export get_isr
func GetISR(isr uint64) uintptr

//export do_lidt
func LoadIDT(idtr uintptr)