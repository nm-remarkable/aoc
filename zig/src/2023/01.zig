const std = @import("std");
const chal = @import("lib").challenge;
const number = @import("lib").number;

fn firstPart(file: std.fs.File) ![]const u8 {
    var buf_reader = std.io.bufferedReader(file.reader());
    var stream = buf_reader.reader();

    var buf: [1024]u8 = undefined;
    while (try stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        for (line) |ch| {
            _ = number.digit(ch);
            // if (d) |dd| {}
        }
    }
    return "bazinga!";
}

fn secondPart(file: std.fs.File) ![]const u8 {
    _ = file;
    return "sad trombone...";
}

pub fn main() !void {
    const c: chal.Challenge = .{ .firstPart = firstPart, .secondPart = secondPart };
    try c.solve();
}
