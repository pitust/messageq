package misc

//export tlb_flush
func FlushTLB()

//export read_cr3
func CR3() uintptr