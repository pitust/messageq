// The messageQ operating system
// Copyright (C) 2021 pitust

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
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