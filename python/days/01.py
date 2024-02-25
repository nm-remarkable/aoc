import re
import sys
from pathlib import Path

from utils.logger import logger

numbers = r"one|two|three|four|five|six|seven|eight|nine"
digit = r"\d" + r"|"

map_to_int = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}


def intify(number: str) -> int:
    if number in map_to_int:
        return map_to_int[number]
    return int(number)


def main() -> int:
    input_file = Path(__file__).parents[2] / "resources" / "01" / "input.txt"
    if not input_file.exists():
        sys.exit(f"Could not find resource file: {input_file.as_posix()}")

    solution = 0
    with input_file.open("r") as f:
        for index, line in enumerate(f):
            first_match = re.search(
                digit + numbers,
                line,
            ).group(0)

            # [::-1] reverses a list or a string
            second_match = re.search(
                digit + numbers[::-1],
                line[::-1],
            ).group(0)[::-1]

            add = int(f"{intify(first_match)}{intify(second_match)}")
            logger.debug(f"line: {index+1}, {add=}")
            solution += add
    return solution
