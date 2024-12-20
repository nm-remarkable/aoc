#include <filesystem>
#include "lib/number.cpp"
#include "lib/format.cpp"
#include "lib/logger.cpp"
#include "lib/resources.cpp"
#include "challenge.cpp"
#include <functional>

int main(int argc, char **argv)
{
    namespace fs = std::filesystem;
    namespace adv = advent::number;

    const int part = advent::challenge::parseArgs(argc, argv);
    const std::function findDigits = part == 1 ? adv::findDigits : adv::findNumbers;

    const std::optional<fs::path> filePath = advent::resources::getResource(1);
    if (filePath)
    {
        std::ifstream file(*filePath);
        if (!file.is_open())
        {
            fmt::println("Failed to open the file.");
            return 1;
        }
        // Iterate through lines of file
        std::string line;
        int sum = 0;
        while (std::getline(file, line))
        {
            sum += adv::atoi(findDigits(line));
        }
        file.close();
        fmt::println("Sum: {}", sum);
        SPDLOG_INFO("Sum: {}", sum);
        SPDLOG_DEBUG("Sum: {}", sum);
        SPDLOG_TRACE("Sum: {}", sum);
        SPDLOG_ERROR("Sum: {}", sum);
        SPDLOG_WARN("Sum: {}", sum);
    }

    return 0;
}
