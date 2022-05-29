import logging
from handle_excel import ExcelHandle
import time
import traceback


class CustomExcelHandle(ExcelHandle):

    def __init__(self, name, upload_file):
        super().__init__(name, upload_file)
        self.custom_name = "梵迄（无痕）"
        self.barcode_to_price = {
            "6973601560041": 60,
            "300030001": 109,
            "3000300040": 158,
            "6973601560034": 43,
            "300030005022": 75,
            "0010073": 97,
            "6973601560317": 34.5,
            "200010050": 54,
            "0010057": 75,
            "6973601560836": 33,
            "0010414": 66,
            "6973601560829": 9,
            "0010423": 18.5,
            "9421900027570": 219,
            "6973601560706": 132,
            "0010222": 220,
            "6973601560720": 132,
            "0010221": 220,
            "6973601560768": 132,
            "6973601560751": 132,
            "0010518": 220,
            "6973601560485": 65,
            "0010106": 97,
            "0010125": 108,
            "6973601560805": 18,
            "0010432": 30,
            "6973601560799": 18,
            "0010315": 30,
            "6973601560492": 99,
            "20010109": 148,
            "6973601560782": 71,
            "0010362": 130,
            "0010348": 183,
            "6972138271079": 18,
            "0010144": 30,
            "0010162": 40,
            "6973601560195": 49,
            "0010355": 76,
            "200010079": 99,
            "0010369": 160,
            "6973601560324": 32,
            "0010052": 57,
            "6973601560355": 29,
            "0010053": 56,
            "0010054": 75,
            "6973601560867": 27,
            "0010449": 45,
            "0010450": 58,
            "0010451": 72,
            "6973601560652": 30,
            "0010468": 53,
            "0010469": 59,
            "6973601560553": 41,
            "0010193": 77,
            "0010177": 108,
            "6973601560362盒": 23,
            "0010357": 41,
            "0010358": 59,
            "6973601560379": 23,
            "0010056": 41,
            "0010148": 59,
            "6973601560409": 54,
            "10010071": 98,
            "6973601560072": 8,
            "300091001": 24,
            "6973601560096": 8,
            "300091002": 24,
            "6973601560065": 11,
            "300083002": 24,
            "6973601560027": 19,
            "300087004": 34,
            "6973601560089": 15,
            "300083001036": 24,
            "6973601560119": 19,
            "0010041": 32,
            "6973601560003": 14,
            "0010117": 24,
            "6949805914188": 11,
            "0010036": 23,
            "0010037": 35,
            "6970869081288": 16,
            "6970869080946": 16,
            "6970869081295": 16,
            "200050039": 31,
            "6939713003579": 16,
            "6939713005085": 16,
            "6939713005092": 16,
            "0010139": 36,
            "6923601907377": 13,
            "0010163": 21,
            "6923601900675": 19,
            "0010291": 30,
            "0010180": 17,
            "0010178": 21.5,
        }

    def start_handle_excel(self):
        logging.info("{} start handle excel".format(self.name))
        wb, ws, items = self.init_read_excel()
        try:
            for index, row in enumerate(items):
                sn = row[2].value
                shop_sn = row[2].value
                receive_people = row[8].value
                phone = row[9].value
                province = ""
                city = ""
                county = ""
                address = row[10].value
                sales_channel_name = self.custom_name
                product_name = row[3].value
                barcode = row[7].value
                unit_price = self.barcode_to_price.get(barcode)
                numbers = row[5].value
                ws.append([sn, shop_sn, "", "", "", "", self.custom_name, "", "", receive_people, phone, "", "",
                           province, city, county, address, "", "", "", "", "", "", "", "", sales_channel_name, "",
                           product_name, barcode, "", "", "", numbers, unit_price, "", "", "", "", "", "", "", "", "",
                           "", "", "", "", "", "", "", ""])
            result_file_path = "./result/shanghaifanqi/梵迄{0}.xlsx".format("".join(str(time.time()).split(".")))
            wb.save(result_file_path)
            logging.info("{0} handle excel file {1} success".format(self.name, self.upload_file))
        except Exception as e:
            exc = traceback.format_exc()
            logging.error("error %s" % exc)
        wb.close()
        self.delete_success_file()
