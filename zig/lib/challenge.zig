const std = @import("std");
const options = @import("options");

pub fn openResourceFile() !std.fs.File {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    const relative_path = try std.fmt.allocPrint(
        allocator,
        "../resources/{s}/{s}/input.txt",
        .{ options.year, options.day },
    );
    const file = try std.fs.cwd().openFile(relative_path, .{});
    defer allocator.free(relative_path);

    return file;
}

pub const Challenge = struct {
    allocator: std.mem.Allocator,
    firstPart: *const fn (std.mem.Allocator, std.fs.File) anyerror!void,
    secondPart: *const fn (std.mem.Allocator, std.fs.File) anyerror!void,

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
                const part = args.next() orelse break;
                if (std.mem.eql(u8, part, "2")) {
                    part1 = false;
                    break;
                }
            }
        }
        const file = try openResourceFile();
        defer file.close();

        if (part1) {
            try self.firstPart(allocator, file);
            return;
        }
        try self.secondPart(allocator, file);
    }
};
