import logging
from watchdog.observers import Observer

from monitor import FileMonitor


class Controller(object):

    def __init__(self, file_path):
        self.file_path = file_path
        self.upload_monitor = FileMonitor(self.handle_upload_file)

    def handle_upload_file(self, custom_name, file):
        logging.info("custom %s upload file %s", custom_name, file)
        # 开始交给处理excel的线程

    def start(self):
        logging.info("start file monitor service")
        observer = Observer()
        observer.schedule(self.upload_monitor, self.file_path, recursive=True)
        observer.start()
        observer.join()
