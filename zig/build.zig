const std = @import("std");
const fmt = @import("std").fmt;

// Although this function looks imperative, note that its job is to
// declaratively construct a build graph that will be executed by an external
// runner.
pub fn build(b: *std.Build) !void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    // Run different executables based on the day
    const options = b.addOptions();
    const day = b.option(u8, "day", "advent day") orelse 1;
    options.addOption(u8, "day", day);

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    var exePath: []u8 = try allocator.alloc(u8, 12);
    if (day < 10) {
        exePath = try fmt.allocPrint(
            allocator,
            "src/0{}.zig",
            .{day},
        );
    } else {
        exePath = try fmt.allocPrint(
            allocator,
            "src/{}.zig",
            .{day},
        );
    }
    defer allocator.free(exePath);

    const lib = b.addStaticLibrary(.{
        .name = "advent",
        .root_source_file = .{ .path = "lib/challenge.zig" },
        .target = target,
        .optimize = optimize,
    });
    b.installArtifact(lib);

    // build system
    const advent = b.addExecutable(.{
        .name = "advent",
        .root_source_file = .{ .path = exePath },
        .target = target,
        .optimize = optimize,
    });
    advent.root_module.addImport("advent", b.createModule(.{
        .root_source_file = .{ .path = "lib/challenge.zig" },
    }));
    b.installArtifact(advent);

    const run_cmd = b.addRunArtifact(advent);
    run_cmd.step.dependOn(b.getInstallStep());
    if (b.args) |args| {
        run_cmd.addArgs(args);
    }
    const run_step = b.step("run", "Run the app");
    run_step.dependOn(&run_cmd.step);
    const advent_unit_tests = b.addTest(.{
        .root_source_file = .{ .path = "src/main.zig" },
        .target = target,
        .optimize = optimize,
    });

    const run_exe_unit_tests = b.addRunArtifact(advent_unit_tests);
    const test_step = b.step("test", "Run unit tests");
    test_step.dependOn(&run_exe_unit_tests.step);
}
