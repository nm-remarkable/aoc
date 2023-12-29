import re
from pathlib import Path

from utils.logger import logger


def main() -> int:
    input_file = Path(__file__).parent / "resources" / "input.01.txt"

    solution = 0
    with input_file.open("r") as f:
        for index, line in enumerate(f):
            first_number = None
            last_number = None

            """matches = re.findall(
                r"\d|one|two|three|four|five|six|seven|eight|nine",
                line,
            )
            if len(matches) == 1:
                first_number = matches.group(0)
                last_number = first_number
            else:"""

            # First number in Line
            for c in line:
                if c.isnumeric():
                    first_number = c
                    break
            # First number in Reversed Line
            for c in line[::-1]:
                if c.isnumeric():
                    last_number = c
                    break
            add = int(f"{first_number}{last_number}")
            logger.debug(f"line: {index+1}, {add=}")
            solution += add
    return solution
