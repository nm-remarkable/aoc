const std = @import("std");
const chal = @import("lib").challenge;
const number = @import("lib").number;

fn firstPart(allocator: std.mem.Allocator, file: std.fs.File) !void {
    const stat = try file.stat();
    const fsize = stat.size;

    const file_content: [:0]u8 = try file.readToEndAllocOptions(allocator, fsize, fsize, @alignOf(u8), 0);
    defer allocator.free(file_content);

    var firstList = try std.ArrayListAligned(i32, @alignOf(i32)).initCapacity(allocator, 1000);
    var secondList = try std.ArrayListAligned(i32, @alignOf(i32)).initCapacity(allocator, 1000);
    defer firstList.deinit();
    defer secondList.deinit();

    var i: u64 = 0;
    var which: bool = true;
    var current: i32 = 0;
    while (i < fsize) {
        const char = file_content[i];
        if (number.digit(char)) |d| {
            current = current * 10 + d;
        } else {
            switch (char) {
                ' ', '\n' => {
                    if (current != 0) {
                        if (which) try firstList.append(current) else try secondList.append(current);
                        current = 0;
                        which = !which; // Flip which list to add item to
                    }
                },
                else => {},
            }
        }
        i += 1;
    }
    std.mem.sort(i32, firstList.items, {}, comptime std.sort.asc(i32));
    std.mem.sort(i32, secondList.items, {}, comptime std.sort.asc(i32));

    var result: usize = 0;
    for (firstList.items, secondList.items) |fi, si| {
        result += @abs(si - fi);

        std.log.debug("fi: {d}, si: {d}, res:{}", .{ fi, si, si - fi });
    }
    std.log.info("Result: {d}", .{result});

    return;
}

fn secondPart(allocator: std.mem.Allocator, file: std.fs.File) !void {
    const stat = try file.stat();
    const fsize = stat.size;
    const file_content: [:0]u8 = try file.readToEndAllocOptions(allocator, fsize, fsize, @alignOf(u8), 0);
    defer allocator.free(file_content);

    var firstList = try std.ArrayListAligned(i32, @alignOf(i32)).initCapacity(allocator, 1000);
    var secondList = try std.ArrayListAligned(i32, @alignOf(i32)).initCapacity(allocator, 1000);
    defer firstList.deinit();
    defer secondList.deinit();

    var i: u64 = 0;
    var which: bool = true;
    var current: i32 = 0;
    while (i < fsize) {
        const char = file_content[i];
        if (number.digit(char)) |d| {
            current = current * 10 + d;
        } else {
            switch (char) {
                ' ', '\n' => {
                    if (current != 0) {
                        if (which) try firstList.append(current) else try secondList.append(current);
                        current = 0;
                        which = !which; // Flip which list to add item to
                    }
                },
                else => {},
            }
        }
        i += 1;
    }
    std.mem.sort(i32, firstList.items, {}, comptime std.sort.asc(i32));
    std.mem.sort(i32, secondList.items, {}, comptime std.sort.asc(i32));

    var similarityList = try std.ArrayListAligned(Similar, @alignOf(Similar)).initCapacity(allocator, 1000);
    defer similarityList.deinit();

    for (firstList.items) |fi| {
        var item = Similar{ .value = fi, .times = 0 };
        for (secondList.items) |si| {
            if (si == item.value) {
                item.times += 1;
            }
        }
        if (item.times != 0) {
            try similarityList.append(item);
            std.log.debug("{any}", .{item});
        }
    }
    var result: i64 = 0;
    for (similarityList.items) |multiples| {
        result += multiples.value * multiples.times;
    }
    std.log.info("Result: {d}", .{result});
    return;
}

const Similar = struct {
    value: i32,
    times: u8,
};

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const c: chal.Challenge = .{ .allocator = arena.allocator(), .firstPart = firstPart, .secondPart = secondPart };
    try c.solve();
}
