const std = @import("std");
const chal = @import("lib").challenge;
const number = @import("lib").number;

const fiveGuys = struct {
    five: [5]u8,
};

fn firstPart(allocator: std.mem.Allocator, file: std.fs.File) !void {
    var buf_reader = std.io.bufferedReader(file.reader());
    var input_stream = buf_reader.reader();
    var buf: [1024]u8 = undefined;
    var result: u16 = 0;
    while (try input_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        const l = line[0 .. line.len - 1]; // Kill off \r
        var parts = std.mem.tokenizeScalar(u8, l, ' ');
        var arr = std.ArrayList(i8).init(allocator);
        defer arr.deinit();

        while (parts.next()) |next| {
            try arr.append(try std.fmt.parseInt(i8, next, 10));
        }

        const slice = try arr.toOwnedSlice();
        defer allocator.free(slice);
        var directions: i8 = 0;
        for (slice[0 .. slice.len - 1], slice[1..slice.len]) |lhs, rhs| {
            const diff = @abs(lhs - rhs);
            if (diff > 3 or diff == 0) {
                directions = 20;
                break;
            }
            if (lhs > rhs) {
                directions += 1;
                continue;
            }
            if (lhs < rhs) {
                directions -= 1;
                continue;
            }
        }
        std.log.info("{d}, {d}", .{ directions, slice.len });
        if (@abs(directions) == slice.len - 1) {
            result += 1;
        }
    }
    std.log.info("Result: {}", .{result});
    return;
}

fn secondPart(allocator: std.mem.Allocator, file: std.fs.File) !void {
    var buf_reader = std.io.bufferedReader(file.reader());
    var input_stream = buf_reader.reader();
    var buf: [1024]u8 = undefined;
    const result: u16 = 0;
    while (try input_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        const l = line[0 .. line.len - 1]; // Kill off \r
        var parts = std.mem.tokenizeScalar(u8, l, ' ');
        var arr = std.ArrayList(i8).init(allocator);
        defer arr.deinit();

        while (parts.next()) |next| {
            try arr.append(try std.fmt.parseInt(i8, next, 10));
        }

        for (0..arr.items.len - 1) |index| {
            const diff = @abs(arr.items[index] - arr.items[index - 1]);
            if (diff > 3 or diff == 0) {
                break;
            }
        }
    }
    std.log.info("Result: {}", .{result});
    return;
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const c: chal.Challenge = .{ .allocator = arena.allocator(), .firstPart = firstPart, .secondPart = secondPart };
    try c.solve();
}
