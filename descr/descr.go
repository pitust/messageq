package descr
type Descriptor struct {
	low uintptr
	hi uintptr
}
func WriteDescriptor(dptr *Descriptor, addr uintptr, limit uint16) {
	dptr.low = (addr << 16) | uintptr(limit)
	dptr.hi = addr >> 48
}