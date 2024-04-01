#include <filesystem>
#include "lib/number.cpp"
#include "lib/format.cpp"
#include "lib/resources.cpp"

int main(void)
{
    namespace fs = std::filesystem;
    namespace adv = advent::number;

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
            sum += adv::parseNumber(adv::findDigits(line));
        }
        file.close();
        fmt::print("Sum: {}", sum);
    }

    return 0;
}
