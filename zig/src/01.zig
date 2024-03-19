const std = @import("std");
const advent = @import("advent");

const Day = struct {
    fn firstPart() []const u8 {
        return "bazinga!";
    }

    fn secondPart() []const u8 {
        return "sad trombone...";
    }

    fn challenge(self: *Day) advent.Challenge {
        return .{
            .ptr = self,
            .src = @src(),
            .firstPart = firstPart,
            .secondPart = secondPart,
        };
    }
};

pub fn main() !void {
    //const input =
    //    \\1abc2
    //    \\pqr3stu8vwx
    //    \\a1b2c3d4e5f
    //    \\treb7uchet
    //;
    var day = Day{};
    const chall = day.challenge();
    try chall.solve();
}
