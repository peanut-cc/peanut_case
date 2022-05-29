import logging
from handle_excel import ExcelHandle
import time
import traceback


class CustomExcelHandle(ExcelHandle):

    def __init__(self, name, upload_file):
        super().__init__(name, upload_file)
        self.custom_name = "钱儿（无痕）"
        self.custom_name2 = "小小包麻麻（无痕）"
        self.result_file_path = "./result/xiaoxiaobaomama/小小包麻麻{0}.xlsx".format(time.strftime("%Y%m%d%H%M%S", time.localtime()))
        self.barcode_to_price = {
            "6970869081288": 17,
            "6970869080946": 17,
            "6970869081295": 17,
            "300091001": 27.83,
            "300091002": 27.24,
            "300083005021": 26,
            "6973601560157": 31.7,
            "200010038": 54.9,
            "6949805914188": 13,
            "0010036": 27,
            "0010037": 41,
            "6939713003579": 21.73,
            "6939713005085": 21.73,
            "6939713005092": 21.73,
            "3000830015": 43.46,
            "6973601560119": 36,
            "0010041": 66.5,
            "6973601560041": 75.89,
            "300030001": 151.78,
            "3000300040": 222.67,
            "6973601560409": 58.18,
            "10010071": 103.86,
            "6973601560317": 50.7,
            "200010050": 78.9,
            "0010057": 109.1,
            "6973601560348": 25.28,
            "0010068": 46.06,
            "6973601560485": 72,
            "0010106": 108,
            "0010125": 120,
            "6973601560836": 47.6,
            "0010414": 95.2,
            "0010509": 42,
        }

    def start_handle_excel(self):
        logging.info("{} start handle excel".format(self.name))
        wb, ws, items = self.init_read_csv()
        try:
            for index, row in enumerate(items):
                sn = row[2]
                shop_sn = row[2]
                receive_people = row[5]
                phone = row[6]
                province = row[7]
                city = row[8]
                county = row[9]
                address = row[11]
                if "钱儿" in row[0]:
                    sales_channel_name = self.custom_name
                else:
                    sales_channel_name = self.custom_name2
                product_name = row[14]
                barcode = row[16]
                unit_price = self.barcode_to_price.get(row[16])
                numbers = row[18]
                ws.append([sn, shop_sn, "", "", "", "", self.custom_name, "", "", receive_people, phone, "", "",
                           province, city, county, address, "", "", "", "", "", "", "", "", sales_channel_name, "",
                           product_name, barcode, "", "", "", numbers, unit_price, "", "", "", "", "", "", "", "", "",
                           "", "", "", "", "", "", "", ""])
            result_file_path = "./result/xiaoxiaobaomama/小小包麻麻{0}.xlsx".format("".join(str(time.time()).split(".")))
            wb.save(result_file_path)
            logging.info("{0} handle excel file {1} success".format(self.name, self.upload_file))
            self.delete_success_file()
        except Exception as e:
            exc = traceback.format_exc()
            logging.error("error %s" % exc)





