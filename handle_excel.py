import os
import logging
import openpyxl
from openpyxl import Workbook


class ExcelHandle(object):

    def __init__(self, name, upload_file):
        self.name = name
        self.upload_file = upload_file
        self.result_excel_header = ["导入编号", "网店订单号", "客户账号", "客户名销售渠道名称称", "收货人", "手机", "省份",
                                    "市（区）", "区（县）", "收货地址", "销售渠道名称", "货品名称", "条码", "数量", "单价"]

    def init_read_excel(self):
        """
        初始化而excel并插入表头数据
        :return:
        """
        workbook = openpyxl.load_workbook(self.upload_file)
        # 获取第一个 sheet 表格
        sheet_name = workbook[workbook.sheetnames[0]]
        wb = Workbook()
        ws = wb.active
        ws.append(self.result_excel_header)
        return wb, ws, list(sheet_name.rows)[1:]

    def delete_success_file(self):
        """
        处理完之后删除excel
        :return:
        """
        if os.path.exists(self.upload_file):
            os.remove(self.upload_file)
            logging.info("delete upload file {0} success".format(self.upload_file))
        else:
            logging.warning("not found upload file {0}", self.upload_file)

