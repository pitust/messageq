package process

import (
	"github.com/pitust/messageq/v2/descr"
	"github.com/pitust/messageq/v2/irq"
	"github.com/pitust/messageq/v2/vm"
)

type Pid uint64
type Tid uint64
type CredID uint64

type Thread struct {
	proc *Process
	id   Tid
	regs irq.Regs
}

type Cred struct {
	name string
	id   CredID
}

type Process struct {
	pid           Pid
	threads       map[Tid]*Thread
	vm            vm.UserVM
	creds         map[CredID]Cred
	has_supercred bool
	has_exited    bool
}

var ProcessTable = make(map[Pid]*Process)
var CredTable = make(map[CredID]Cred)

var pidgen = uint64(0)

func MakePID() Pid {
	pidgen++
	return Pid(pidgen)
}
func MakeTID() Tid {
	pidgen++
	return Tid(pidgen)
}

func CreateProcess() *Process {
	pid := MakePID()
	proc := &Process{
		pid:           pid,
		vm:            vm.KernelVM.Clone(),
		creds:         make(map[CredID]Cred),
		threads:       make(map[Tid]*Thread),
		has_supercred: false,
		has_exited:    false,
	}
	ProcessTable[pid] = proc
	return proc
}

func (this *Process) CreateThread() *Thread {
	if this.has_exited {
		panic("cannot create threads on an exited process")
	}
	tid := MakeTID()
	thr := &Thread{
		proc: this,
		id:   tid,
		regs: irq.Regs{
			CS: descr.GDT_USER_CODE,
			SS: descr.GDT_USER_DATA,
		},
	}
	this.threads[tid] = thr
	return thr
}
func (this *Thread) Regs() *irq.Regs {
	return &this.regs
}
func (this *Thread) With(cb func()) {
	this.proc.vm.With(cb)
}
func (this *Process) HasCred(id CredID) bool {
	_, has_cred := this.creds[id]
	return has_cred || this.has_supercred
}
func (this *Thread) VM() *vm.UserVM {
	return &this.proc.vm
}