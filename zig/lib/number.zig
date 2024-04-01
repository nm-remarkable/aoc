const std = @import("std");

pub fn digit(ch: u8) ?u8 {
    if (ch >= '0' and ch <= '9') {
        return ch - '0';
    }
    return null;
}
