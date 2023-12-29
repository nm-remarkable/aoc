"""Examples:
    poetry run advent 00 --debug
    python -m advent.py 01 --debug
"""

import argparse
import logging
import sys
from pathlib import Path
from pydoc import importfile

from utils.logger import logger


def parse_args(parser: argparse.ArgumentParser) -> argparse.Namespace:
    parser.add_argument(
        "n",
        type=str,
        help="Number of the challenge to run ex: 00, 01, 02",
    )
    parser.add_argument("--debug", action="store_true", help="Run logger in debug mode")
    return parser.parse_args(sys.argv[1:])


def execute() -> None:
    parser = argparse.ArgumentParser(
        formatter_class=argparse.RawTextHelpFormatter,
        description=__doc__,
    )
    args = parse_args(parser)
    logger.setLevel(logging.DEBUG if args.debug else logging.INFO)

    directory = Path(__file__).parent / "days"
    challenges = list(directory.glob("**/*"))

    solution = ""
    for c in challenges:
        if c.stem == args.n:
            module = importfile(c.as_posix())
            logger.critical(module)
            solution = module.main()

    if not solution:
        parser.exit(f"FAIL: Options for days are {[ f.stem for f in challenges]}")

    logger.info(f"Solution: {solution}")
