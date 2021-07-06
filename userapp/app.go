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

// example userland app
package main

import "github.com/pitust/messageq/v2/uapi"
import "github.com/pitust/messageq/v2/uapi/uapicreds"

func main() {
	uapi.KeGrant(uapi.KeGetPid(), uapicreds.PORT_IO)
	uapi.KePledge(uapicreds.SUPERCRED)
	uapi.KeOutByte(0xe9, 'h')
}