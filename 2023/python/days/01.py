import re
from pathlib import Path

from utils.logger import logger

numbers = r"one|two|three|four|five|six|seven|eight|nine"
digit = r"\d" + r"|"


def intify(number: str) -> int:
    match number:
        case "one":
            return 1
        case "two":
            return 2
        case "three":
            return 3
        case "four":
            return 4
        case "five":
            return 5
        case "six":
            return 6
        case "seven":
            return 7
        case "eight":
            return 8
        case "nine":
            return 9
        case _:
            return int(number)


def main() -> int:
    input_file = Path(__file__).parent / "resources" / "input.01.txt"

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
