import sys
import logging
from logging.handlers import RotatingFileHandler
from config.settings import log_config
from config.settings import customs_config

from controller import Controller


def init_log():
    rotate_handler = RotatingFileHandler(
        log_config['file_path'], "a", log_config['size'] * 1024 * 1024, log_config['backup']
    )
    rotate_handler.setLevel(logging.INFO)
    formatter = logging.Formatter(
        '[%(asctime)s] [%(filename)s:%(lineno)d] %(levelname)s %(message)s'
    )
    console = logging.StreamHandler()
    console.setFormatter(formatter)
    rotate_handler.setFormatter(formatter)
    log = logging.getLogger()
    log.addHandler(rotate_handler)
    log.addHandler(console)
    log.setLevel(logging.INFO)


if __name__ == '__main__':
    init_log()
    uploads_path = "{0}/customs".format(sys.path[0])
    logging.info(uploads_path)
    Controller(file_path=uploads_path).start()



