const std = @import("std");
const chal = @import("lib").challenge;
const number = @import("lib").number;

fn firstPart(file: std.fs.File) ![]const u8 {
    _ = file;
    return "Not implemented";
}

fn secondPart(file: std.fs.File) ![]const u8 {
    _ = file;
    return "sad trombone...";
}

pub fn main() !void {
    const c: chal.Challenge = .{ .firstPart = firstPart, .secondPart = secondPart };
    try c.solve();
}
