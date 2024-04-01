#pragma once

#include <iostream>
#include <format>

namespace fmt
{
    template <typename... Args>
    inline void print(const std::string_view str, Args &&...args)
    {
        std::cout << std::vformat(str, std::make_format_args(args...));
    }

    template <typename... Args>
    inline std::string format(const std::string str, Args &&...args)
    {
        return std::vformat(str, std::make_format_args(args...));
    }

    void println(const std::string &message)
    {
        fmt::print("{}\n", message);
    }
}
