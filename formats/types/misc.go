/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

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
