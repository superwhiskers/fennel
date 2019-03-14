mod fennel;
use libc;
#[link(name = "fennel", kind = "static")]

extern fn fennel_newAccountServerClient(
    accountServer: *const u8,
    certificatePath: *const u8,
    keyPath: *const u8,
) -> *mut libc::c_void;
