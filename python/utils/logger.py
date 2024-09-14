"""Logger"""

import logging
import sys

FORMAT = "%(asctime)s - %(levelname)s -- %(message)s"
logging.basicConfig(stream=sys.stdout, level=logging.INFO, format=FORMAT)
logger = logging.getLogger()
