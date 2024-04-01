const std = @import("std");
const Challenge = @import("advent::challenge").Challenge;
const number = @import("advent::number");

const Day = struct {
    fn firstPart(file: std.fs.File) ![]const u8 {
        var buf_reader = std.io.bufferedReader(file.reader());
        var stream = buf_reader.reader();

        var buf: [1024]u8 = undefined;
        while (try stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
            for (line) |ch| {
                const d = number.digit(ch);
                if (d) |dd| {
                    std.log.warn("dd: {d}", .{dd});
                }
            }
        }
        return "bazinga!";
    }

    fn secondPart(file: std.fs.File) ![]const u8 {
        _ = file;
        return "sad trombone...";
    }

    fn challenge(self: *Day) Challenge {
        return .{
            .ptr = self,
            .src = @src(),
            .firstPart = firstPart,
            .secondPart = secondPart,
        };
    }
};

pub fn main() !void {
    var day = Day{};
    const chall = day.challenge();
    try chall.solve();
}
