# Rust Advent 2023

Important to note, I am using zig version 0.14.0-dev. There are quite a few breaking changes currently so be aware this might be broken the next time you touch it <---- It fucking was :facepalm:

Building and running zig using only compile time options
- `zig build run -Dday=01 -Dyear=2023 --summary all`

Building and running zig using args for the part
- `zig build run -Dday=01 -Dyear=2024 -- --part 2`

Target state:
- `zig build run -- --day 01 --part 2`

