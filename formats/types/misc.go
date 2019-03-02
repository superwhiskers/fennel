/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package types

import (
	"crypto/sha256"
	"encoding/binary"
)

// HashPassword hashes an nnid password based using the password and the pid
func HashPassword(upid uint32, password string) []byte {

	pid := make([]byte, 4)
	binary.LittleEndian.PutUint32(pid, upid)

	data := sha256.Sum256(append(pid, append([]byte{0x02, 0x65, 0x43, 0x46}, []byte(password)...)...))

	return data[:]

}
