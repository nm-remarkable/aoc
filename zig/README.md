# Rust Advent 2023

Important to note, I am using zig version 0.12.0-dev.3343+294f51814. There are quite a few breaking changes currently so be aware this might be broken the next time you touch it.

Building and running zig using only compile time options
- `zig build run -Dday=01 -Dpart=2 --summary all`

Building and running zig using args for the part
- `zig build run -Dday=01 -- --part 2`

Target state:
- `zig build run -- --day 01 --part 2`
