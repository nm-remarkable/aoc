const std = @import("std");
const config = @import("options");

pub fn main() !void {
    //const input =
    //    \\1abc2
    //    \\pqr3stu8vwx
    //    \\a1b2c3d4e5f
    //    \\treb7uchet
    //;
    if (config.part1) {
        std.debug.print("\nPart1\n", .{});
    }
    std.debug.print("\nIs it part 1? {}\n", .{config.part1});
}
