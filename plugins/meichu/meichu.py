import logging
import traceback
from handle_excel import ExcelHandle
import time


class CustomExcelHandle(ExcelHandle):

    def __init__(self, name, upload_file):
        super().__init__(name, upload_file)
        self.custom_name = "美初（无痕）"
        self.result_file_path = "./result/meichu/美初{0}.xlsx".format(time.strftime("%Y%m%d%H%M%S", time.localtime()))
        self.code_to_barcode = {
            "E211231-3": "0010414",
            "E211231-1": "6973601560836"
        }
        self.barcode_to_price = {
            "0010414": 59,
            "6973601560836": 28
        }

    def start_handle_excel(self):
        logging.info("{} start handle excel".format(self.name))
        wb, ws, items = self.init_read_excel()
        try:
            for index, row in enumerate(items):
                sn = row[0].value
                shop_sn = row[0].value
                receive_people = row[3].value
                phone = row[4].value
                province = row[5].value
                city = row[6].value
                county = row[7].value
                address = row[9].value
                sales_channel_name = self.custom_name
                product_name = row[10].value
                format_name = row[11].value
                barcode = row[12].value
                unit_price = ""
                if product_name == "仁和黄金油柑酵素" and format_name == "3盒（买二送一）" and barcode == "E211231-3":
                    barcode = self.code_to_barcode.get("E211231-3")
                    unit_price = self.barcode_to_price.get(barcode)
                elif product_name == "仁和黄金油柑酵素" and format_name == "1盒" and barcode == "E211231-1":
                    barcode = self.code_to_barcode.get("E211231-1")
                    unit_price = self.barcode_to_price.get(barcode)
                else:
                    barcode = "条码错误"
                    logging.error("handle excel {0} row {1} barcode error".format(self.upload_file, index))

                numbers = row[13].value
                ws.append([sn, shop_sn, "", self.custom_name, receive_people, phone, province, city, county, address,
                           sales_channel_name, product_name, barcode, numbers, unit_price])
            wb.save(self.result_file_path)
            logging.info("{0} handle excel file {1} success".format(self.name, self.upload_file))
            self.delete_success_file()
        except Exception as e:
            exc = traceback.format_exc()
            logging.error("error %s" % exc)
