import os
import logging


class ExcelHandle(object):

    def __init__(self, name, upload_file):
        self.name = name
        self.upload_file = upload_file
        self.result_excel_header = ["导入编号", "网店订单号", "客户账号", "客户名销售渠道名称称", "收货人", "手机", "省份",
                                    "市（区）", "区（县）", "收货地址", "销售渠道名称", "货品名称", "条码", "数量", "单价"]

    def delete_success_file(self):
        if os.path.exists(self.upload_file):
            os.remove(self.upload_file)
            logging.info("delete upload file {0} success".format(self.upload_file))
        else:
            logging.warning("not found upload file {0}", self.upload_file)

