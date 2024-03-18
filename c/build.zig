const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    const advent = b.addExecutable(.{
        .name = "advent",
        .optimize = optimize,
        .target = target,
    });
    advent.linkLibC();
    advent.addCSourceFile(.{ .file = .{
        .path = "src/main.c",
    }, .flags = &.{
        "-std=c11",
        "-pedantic",
        "-Wall",
        "-W",
        "-Wno-missing-field-initializers",
    } });

    b.installArtifact(advent);
}
