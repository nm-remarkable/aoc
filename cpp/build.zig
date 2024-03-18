const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    const advent = b.addExecutable(.{
        .name = "advent",
        .optimize = optimize,
        .target = target,
    });
    advent.linkLibCpp();
    advent.addCSourceFile(.{ .file = .{
        .path = "src/main.cpp",
    }, .flags = &.{
        "-std=c++23",
        "-pedantic",
        "-Wall",
        "-Wextra",
        "-Werror",
        "-Wno-missing-field-initializers",
    } });

    b.installArtifact(advent);
}
