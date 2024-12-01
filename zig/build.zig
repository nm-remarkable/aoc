const std = @import("std");
const fmt = @import("std").fmt;

// Although this function looks imperative, note that its job is to
// declaratively construct a build graph that will be executed by an external
// runner.
pub fn build(b: *std.Build) !void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    // Run different executables based on the day or year
    const options = b.addOptions();
    const day = b.option([]const u8, "day", "advent day") orelse "01";
    const year = b.option([]const u8, "year", "advent year") orelse "2023";
    options.addOption([]const u8, "day", day);
    options.addOption([]const u8, "year", year);

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    var exePath: []const u8 = try allocator.alloc(u8, 24);
    exePath = try fmt.allocPrint(
        allocator,
        "src/{s}/{s}.zig",
        .{ year, day },
    );
    defer allocator.free(exePath);

    // Advent library
    const lib2 = b.addStaticLibrary(.{
        .name = "lib",
        .root_source_file = b.path("lib/lib.zig"),
        .target = target,
        .optimize = optimize,
    });
    lib2.root_module.addOptions("options", options);
    b.installArtifact(lib2);

    // build system
    const advent = b.addExecutable(.{
        .name = "advent",
        .root_source_file = b.path(exePath),
        .target = target,
        .optimize = optimize,
    });
    advent.root_module.addImport("lib", &lib2.root_module);
    b.installArtifact(advent);

    const run_cmd = b.addRunArtifact(advent);
    run_cmd.step.dependOn(b.getInstallStep());
    if (b.args) |args| {
        run_cmd.addArgs(args);
    }
    const run_step = b.step("run", "Run the app");
    run_step.dependOn(&run_cmd.step);
    const advent_unit_tests = b.addTest(.{
        .root_source_file = b.path("src/main.zig"),
        .target = target,
        .optimize = optimize,
    });
    advent.root_module.addImport("lib", &lib2.root_module);

    const run_exe_unit_tests = b.addRunArtifact(advent_unit_tests);
    const test_step = b.step("test", "Run unit tests");
    test_step.dependOn(&run_exe_unit_tests.step);
}
