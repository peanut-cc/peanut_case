# 监听每个客户的 uploads 目录 用于通知其他开始处理
import os

from uuid import uuid4
from watchdog.events import FileSystemEventHandler
import logging


class FileMonitor(FileSystemEventHandler):

    def __init__(self, callback):
        self.callback = callback

    def on_moved(self, event):
        super(FileMonitor, self).on_moved(event)
        what = 'directory' if event.is_directory else 'file'
        # 后面用到可以添加对应的回调

    def on_deleted(self, event):
        super(FileMonitor, self).on_deleted(event)
        what = 'directory' if event.is_directory else 'file'
        # 后面用到可以添加对应的回调

    def on_modified(self, event):
        super(FileMonitor, self).on_modified(event)
        what = 'directory' if event.is_directory else 'file'
        # 后面用到可以添加对应的回调

    def on_created(self, event):
        super(FileMonitor, self).on_moved(event)
        if not event.is_directory:
            file_path = os.path.dirname(event.src_path)
            dir_name = os.path.basename(file_path)
            file_name = os.path.basename(event.src_path)
            if ".csv" in file_name:
                new_name = "{}/{}.csv".format(file_path, uuid4())
            else:
                new_name = "{}/{}.xlsx".format(file_path, uuid4())
            logging.info("custom {0} upload file {1}".format(dir_name,event.src_path))
            os.rename(event.src_path, new_name)
            self.callback(dir_name, new_name)

