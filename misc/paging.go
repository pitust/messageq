package misc

//export tlb_flush
func FlushTLB()

//export read_cr2
func CR2() uintptr

//export read_cr3
func CR3() uintptr