# 监听每个客户的 uploads 目录 用于通知其他开始处理
import os

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
            self.callback(dir_name, event.src_path)

