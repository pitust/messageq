package main

import "github.com/pitust/messageq/v2/uapi"
import "github.com/pitust/messageq/v2/uapi/uapicreds"

func main() {
	uapi.KeGrant(uapi.KeGetPid(), uapicreds.PORT_IO)
	uapi.KePledge(uapicreds.SUPERCRED)
	uapi.KeOutByte(0xe9, 'h')
}