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

// IRQ handling
package irq

const (
	SERVICE_CLASS_REGS = 0x1000

	SERVICE_STORE_REGS = SERVICE_CLASS_REGS | 1
	SERVICE_SCHEDULE_USER = SERVICE_CLASS_REGS | 2
)

const (
	SERVICE_OK = 0x1000
	SERVICE_IRQ_NOTE = 0x0100
	SERVICE_ERR = 0x2000
	SERVICE_ERR_BAD_CALL = SERVICE_ERR | 1
)