/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

mod fennel;
use libc;
#[link(name = "fennel", kind = "static")]

extern fn fennel_newAccountServerClient(
    accountServer: *const u8,
    certificatePath: *const u8,
    keyPath: *const u8,
) -> *mut libc::c_void;
