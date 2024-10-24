#pragma once

#include <iostream>
#include <format>
#include "format.cpp"

constexpr std::string_view reset = "\x1b[0m";
constexpr std::string_view red = "\x1b[31m";
constexpr std::string_view green = "\x1b[32m";
constexpr std::string_view yellow = "\x1b[33m";
constexpr std::string_view blue = "\x1b[34m";
constexpr std::string_view cyan = "\x1b[36m";

struct lg
{
    enum LEVEL
    {
        TRACE,
        DEBUG,
        INFO,
        WARN,
        ERROR,
    };
    void setLevel(lg::LEVEL lv)
    {
        level = lv;
    }
    lg::LEVEL level = lg::LEVEL::INFO;
};

static std::unique_ptr<lg> logger = std::make_unique<lg>();

template <typename... Args>
inline void SPDLOG_INFO(const std::string_view &str, Args &&...args)
{
    if (logger->level > lg::LEVEL::INFO)
    {
        return;
    }
    fmt::println("{}[info] {} {}", green, reset, std::vformat(str, std::make_format_args(args...)));
}

template <typename... Args>
inline void SPDLOG_DEBUG(const std::string_view &str, Args &&...args)
{
    if (logger->level > lg::LEVEL::DEBUG)
    {
        return;
    }
    fmt::println("{}[debug]{} {}", cyan, reset, std::vformat(str, std::make_format_args(args...)));
}

template <typename... Args>
inline void SPDLOG_TRACE(const std::string_view &str, Args &&...args)
{
    if (logger->level > lg::LEVEL::TRACE)
    {
        return;
    }
    fmt::println("{}[trace]{} {}", blue, reset, std::vformat(str, std::make_format_args(args...)));
}

template <typename... Args>
inline void SPDLOG_WARN(const std::string_view &str, Args &&...args)
{
    if (logger->level > lg::LEVEL::WARN)
    {
        return;
    }
    fmt::println("{}[warn] {} {}", yellow, reset, std::vformat(str, std::make_format_args(args...)));
}

template <typename... Args>
inline void SPDLOG_ERROR(const std::string_view &str, Args &&...args)
{
    if (logger->level > lg::LEVEL::ERROR)
    {
        return;
    }
    fmt::println("{}[error]{} {}", red, reset, std::vformat(str, std::make_format_args(args...)));
}
