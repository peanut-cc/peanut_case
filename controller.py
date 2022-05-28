import logging
import importlib
from watchdog.observers import Observer

from monitor import FileMonitor


class Controller(object):

    def __init__(self, file_path):
        self.file_path = file_path
        self.upload_monitor = FileMonitor(self.handle_upload_file)
        self.custom_class = "CustomExcelHandle"
        self.sep = ":"

    def plugin_load(self, plugin_name, package, name, upload_file):
        m, _, c = plugin_name.partition(self.sep)
        mod = importlib.import_module(m, package)
        cls = getattr(mod, c)
        return cls(name, upload_file)

    def handle_upload_file(self, custom_name, file):
        logging.info("custom %s upload file %s", custom_name, file)
        plugin = ".{0}:{1}".format(custom_name, self.custom_class)
        package = "plugins.{}".format(custom_name)
        custom = self.plugin_load(plugin, package, custom_name, file)
        custom.start_handle_excel()

    def start(self):
        logging.info("start file monitor service")
        observer = Observer()
        observer.schedule(self.upload_monitor, self.file_path, recursive=True)
        observer.start()
        observer.join()
