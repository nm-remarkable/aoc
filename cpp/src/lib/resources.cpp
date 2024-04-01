#pragma once

#include <filesystem>
#include <fstream>
#include <optional>
#include "format.cpp"

namespace fs = std::filesystem;
namespace advent::resources
{

    std::optional<fs::path> getResource(const int day)
    {
        fs::path filePath = fs::current_path().string() + fmt::format("../../resources/0{}/input.txt", day);
        if (fs::exists(filePath))
        {
            return {filePath};
        }
        else
        {
            fmt::print("File {} not found.", filePath.string());
        }
        return {};
    }
}
