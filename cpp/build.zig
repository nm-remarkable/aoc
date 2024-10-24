const std = @import("std");

pub fn build(b: *std.Build) !void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});
    const cppFlags = .{
        "-std=c++23",
        "-pedantic",
        "-Wall",
        "-Wextra",
        "-Werror",
        "-Wno-missing-field-initializers",
    };

    // Run different executables based on the day
    const options = b.addOptions();
    const day = b.option(u8, "day", "advent day") orelse 1;
    options.addOption(u8, "day", day);

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    var exePath: []u8 = try allocator.alloc(u8, 12);
    exePath = try std.fmt.allocPrint(allocator, "src/0{}.cpp", .{day});
    defer allocator.free(exePath);

    const advent = b.addExecutable(.{
        .name = "advent",
        .optimize = optimize,
        .target = target,
    });
    advent.linkLibCpp();
    advent.addCSourceFile(.{
        .file = .{ .cwd_relative = exePath },
        .flags = &cppFlags,
    });

    b.installArtifact(advent);
    const run_cmd = b.addRunArtifact(advent);
    run_cmd.step.dependOn(b.getInstallStep());
    if (b.args) |args| {
        run_cmd.addArgs(args);
    }
    const run_step = b.step("run", "Run the app");
    run_step.dependOn(&run_cmd.step);
}
