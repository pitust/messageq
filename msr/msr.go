package msr

//export do_rdmsr
func Read(msr uint32) uint64

//export do_wrmsr
func Write(msr uint32, val uint64)