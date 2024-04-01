#include <string>
#include "lib/format.cpp"
#include "lib/number.cpp"

namespace advent::challenge
{
    int parseArgs(int argc, char **argv)
    {
        int part = 1;
        if (argc > 1)
        {
            if (advent::number::atoi(argv[1]) == 2                                        //
                || (strcmp(argv[1], "--part") == 0 && advent::number::atoi(argv[2]) == 2) //
            )
            {
                part = 2;
            }
        }
        fmt::println("Part: {}", part);
        return part;
    }
}
