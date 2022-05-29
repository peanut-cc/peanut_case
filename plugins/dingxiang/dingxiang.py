import logging
from handle_excel import ExcelHandle
import time
import traceback


class CustomExcelHandle(ExcelHandle):

    def __init__(self, name, upload_file):
        super().__init__(name, upload_file)
        self.custom_name = "丁香（无痕）"

    def start_handle_excel(self):
        logging.info("{} start handle excel".format(self.name))
        wb, ws, items = self.init_read_excel()
        try:
            for index, row in enumerate(items):
                sn = row[2].value
                shop_sn = row[2].value
                receive_people = row[19].value
                phone = row[20].value
                province = row[21].value
                city = row[22].value
                county = row[23].value
                address = row[24].value
                sales_channel_name = self.custom_name
                product_name = row[10].value
                barcode = row[13].value
                numbers = row[15].value
                unit_price = row[18].value
                ws.append([sn, shop_sn, "", "", "", "", self.custom_name, "", "", receive_people, phone, "", "",
                           province, city, county, address, "", "", "", "", "", "", "", "", sales_channel_name, "",
                           product_name, barcode, "", "", "", numbers, unit_price, "", "", "", "", "", "", "", "", "",
                           "", "", "", "", "", "", "", ""])
            result_file_path = "./result/dingxiang/丁香{0}.xlsx".format("".join(str(time.time()).split(".")))
            wb.save(result_file_path)
            logging.info("{0} handle excel file {1} success".format(self.name, self.upload_file))

        except Exception as e:
            exc = traceback.format_exc()
            logging.error("error %s" % exc)
        wb.close()
        self.delete_success_file()











