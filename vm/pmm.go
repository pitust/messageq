package vm

import "unsafe"

type MemoryRegion struct {
	Base uintptr
	Len  uintptr
}

var mem []MemoryRegion = []MemoryRegion{}

func PhysAlloc() unsafe.Pointer {
	if len(mem) == 0 {
		panic("OOM")
	}
	tregion := &mem[0]
	for tregion.Len < 4096 {
		mem = mem[1:]
		tregion = &mem[0]
	}
	tgbase := unsafe.Pointer(tregion.Base)
	if tregion.Len > 4096 {
		tregion.Base += 4096
		tregion.Len -= 4096
		return tgbase
	}
	mem = mem[1:]
	return tgbase
}

func PhysFree(ptr unsafe.Pointer) {
	PhysAddRegion(MemoryRegion{Base: uintptr(ptr), Len: 4096})
}

func PhysAddRegion(r MemoryRegion) {
	mem = append(mem, r)
}
