const std = @import("std");

pub fn main() !void {
    //const input =
    //    \\1abc2
    //    \\pqr3stu8vwx
    //    \\a1b2c3d4e5f
    //    \\treb7uchet
    //;
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();
    var args = try std.process.argsWithAllocator(allocator);
    defer args.deinit();

    var part1 = true;
    while (args.next()) |arg| {
        // Only support "--part 1/2" for now
        if (std.mem.eql(u8, arg, "--part")) {
            const part = args.next() orelse break; // orelse part1 = true;
            if (std.mem.eql(u8, part, "2")) {
                part1 = false;
                break;
            }
        }
    }
    const partNumber: u8 = if (part1) 1 else 2;
    const filename = std.fs.path.stem(@src().file);
    std.log.info("Running advent day {s} part {}\n", .{ filename, partNumber });
}
