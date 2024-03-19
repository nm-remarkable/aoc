const std = @import("std");

pub const Challenge = struct {
    ptr: *anyopaque,
    src: std.builtin.SourceLocation,

    firstPart: *const fn () []const u8,
    secondPart: *const fn () []const u8,

    pub fn day(self: Challenge) []const u8 {
        return std.fs.path.stem(self.src.file);
    }

    pub fn solve(self: Challenge) anyerror!void {
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
        std.log.info("Running advent day {s} part {}", .{ self.day(), partNumber });

        const solution: []const u8 = if (part1) self.firstPart() else self.secondPart();
        std.log.info("Solution for day {s} part {}: {s}", .{ self.day(), partNumber, solution });
    }
};
