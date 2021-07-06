package uapi

import "github.com/pitust/messageq/v2/uapi/uapicreds"

type Pid uint64
type Error uint64

const (
	EOK = Error(0)
	EACCES = Error(1)
)

//export int21
func KeGetPid() Pid

//export int22
func KeGrant(pid Pid, cred uapicreds.Cred) Error

//export int23
func KePledge(cred uapicreds.Cred)

//export int24
func KeOutByte(port uint16, val uint8)