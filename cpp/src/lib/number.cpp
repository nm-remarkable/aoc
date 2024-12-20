#pragma once

#include <optional>
#include <string>
#include <regex>
#include <unordered_map>
#include "format.cpp"

namespace advent::number
{
    std::optional<int> digit(char ch)
    {
        if (ch >= '1' && ch <= '9')
        {
            return {ch - '0'};
        }
        return {};
    }

    std::string findDigits(const std::string &s)
    {
        int startPos = 0, left = 0, right = 0;
        const int len = s.length();
        while (startPos < len)
        {
            const std::optional<int> num = digit(s[startPos]);
            if (num)
            {
                left = *num;
                break;
            }
            startPos++;
        }
        startPos = len - 1;
        while (startPos >= 0)
        {
            const std::optional<int> num = digit(s[startPos]);
            if (num)
            {
                right = *num;
                break;
            }
            startPos--;
        }
        return fmt::format("{}{}", left, right);
    }

    int atoi(std::string s)
    {
        int startPos = 0, n = 0, sign = 1;
        while (s[startPos] == ' ')
        {
            startPos++;
        }
        if (s[startPos] == '-' || s[startPos] == '+')
        {
            sign = s[startPos] == '+' ? 1 : -1;
            startPos++;
        }

        const int stringLength = s.length();
        for (int i = startPos; i < stringLength; i++)
        {
            const std::optional<int> num = advent::number::digit(s[i]);
            if (!num)
            {
                break;
            }
            if (n > (INT_MAX - *num) / 10)
            {
                return sign == 1 ? INT_MAX : INT_MIN;
            }
            n = n * 10 + *num;
        }
        return n * sign;
    }

    std::string findNumbers(const std::string &s)
    {
        std::unordered_map<std::string, int> digits = {
            {"one", 1},
            {"two", 2},
            {"three", 3},
            {"four", 4},
            {"five", 5},
            {"six", 6},
            {"seven", 7},
            {"eight", 8},
            {"nine", 9}};
        int startPos = 0, left = 0, right = 0;
        const int len = s.length();
        while (startPos < len)
        {
            const std::optional<int> num = digit(s[startPos]);
            if (num)
            {
                left = *num;
                goto leftSideLoopEnd;
            }
            for (const auto &[word, d] : digits)
            {
                const std::string window = s.substr(startPos, word.length());
                if (window == word)
                {
                    left = d;
                    goto leftSideLoopEnd;
                }
            }
            startPos++;
        }
    leftSideLoopEnd:
        startPos = len - 1;
        while (startPos >= 0)
        {
            const std::optional<int> num = digit(s[startPos]);
            if (num)
            {
                right = *num;
                goto rightSideLoopEnd;
            }
            for (const auto &[word, d] : digits)
            {
                const std::string window = s.substr(startPos, word.length());
                if (window == word)
                {
                    right = d;
                    goto rightSideLoopEnd;
                }
            }
            startPos--;
        }
    rightSideLoopEnd:
        return fmt::format("{}{}", left, right);
    }
}
