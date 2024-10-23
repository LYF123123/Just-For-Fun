import os
import pandas as pd
import xlrd





# 定义寻找待读取内容对应头

headers=['姓名:','联系号码:', '年龄:', '性别:', '体检编号:','症状', '体温', '脉率', '呼吸频率','身高','体重','体质指数','腰围','臀围','腰臀围比','锻炼频率','每次锻炼时间','坚持锻炼时间','锻炼方式','饮食习惯','吸烟状况','日吸烟量','开始吸烟年龄','戒烟年龄','饮酒频率','日饮酒量','是否戒酒','开始饮酒年龄','开始戒酒年龄','近一年内是否曾醉酒','饮酒种类','职业病危害因素','口唇','咽部','齿列','运动功能','听力','左眼','右眼','皮肤','巩膜','淋巴结','桶状胸','呼吸音','罗音','心率','心律','杂音','压痛','肝脏','包块','脾脏','移动性浊音','下肢水肿','足背动脉搏动','胸部X线片','心电图','B 超','CA199','CEA','血清肌酐','eGFR根据CKD-EPI公式','AFP','尿素测定','空腹血糖','糖化血红蛋白','血清谷丙转氨酶','血清低密度脂蛋白胆固醇','甘油三酯','血清高密度脂蛋白胆固醇','血尿酸','血清谷草转氨酶','总胆固醇','总胆红素','脑血管疾病','心脏疾病','血管疾病','神经系统疾病','肾脏疾病','眼部疾病','其他系统疾病','平和质','阴虚质','气郁质','气虚质','湿热质','特秉质','血瘀质','阳虚质','痰湿质','异常1','异常2','异常3','异常4','异常5','异常6','健康指导:','其他指导','危险因素控制:','减重目标体重','体检结论:' ,'体检建议:']
#待读取内容对应头所在单元格
cells_to_read=[(14, 3),
        (18, 96),
        (14, 46),
        (14, 27),
        (14, 96),
        (28, 4),
        (32,4),(32,53),(32,98),
        (40,4),(40,34),(40,75),
        (44,4),(44,34),(44,75),
        (51,11),(50,75),
        (55,11),(55,75),
        (58,3),
        (61,11),(61,83),
        (65,11),(65,83),
        (68,11),(68,83),
        (72,11),(72,83),
        (75,11),(75,83),
        (78,11),
        (81,3),
        (99,10),(102,10),(99,63),
        (106,3),
        (111,3),
        (114,10),
        (117,10),
        (123,3),(123,50),(123,85),
        (127,8),(127,85),(131,8),
        (133,8),(133,56),(133,85),
        (137,8),(137,40),(137,85),
        (140,8),(140,40),
        (143,3),(146,3),
        (151,3),(153,3),(156,3),
        (206,2),(206,71),
        (208,2),(208,71),
        (211,2),(211,71),
        (213,2),(213,71),
        (215,2),(215,71),
        (219,2),(219,71),
        (222,2),(222,71),
        (224,2),(224,71),
        (228,2),(228,69),(231,2),(231,69),(234,2),(234,69),(237,2),
        (242,2),(242,37),(242,83),
        (246,2),(246,37),(246,83),
        (249,2),(249,37),(249,83),
        (255,3),(255,68),
        (259,3),(259,68),
        (263,3),(263,68),
        (268,2),
        (265,17),
        (269,2),
        (272,17),
        (270,2),
        (270,2)
        ]



output_headers=['序号', '姓名','联系电话', '年龄','性别','体检编号','症状','体温（℃）','脉率（次/分钟）','呼吸频率','身高（cm）','体重（kg）','体质指数（BMI）','腰围','臀围','腰臀围比','腰围身高比','锻炼频率','每次锻炼时间（分钟）','坚持锻炼时间（年）','锻炼方式','饮食习惯','吸烟状况','日吸烟量（支）','开始吸烟年龄（岁）','戒烟年龄','饮酒频率','日饮酒量\r\n（两）','是否戒酒','开始饮酒年龄','开始戒酒年龄','近一年内是否曾醉酒','饮酒种类','职业病危害因素','口唇','咽部','齿列','运动功能','听力','左眼视力','右眼视力','皮肤','巩膜','淋巴结','桶状胸','呼吸音','罗音','心率\r\n（次/分钟）','心律','杂音','腹部压痛','肝脏','包块','脾脏','移动性浊音','下肢水肿','足背动脉搏动','胸部X线片','心电图','B超','CA199','CEA','血清肌酐','eGFR根据CKD-EPI公式','AFP','尿素测定','空腹血糖','糖化血红蛋白','血清谷丙转氨酶','血清低密度脂蛋白胆固醇','甘油三酯','血清高密度脂蛋白胆固醇','血尿酸','血清谷草转氨酶','总胆固醇','总胆红素','脑血管疾病','心脏疾病','血管疾病','神经系统疾病','肾脏疾病','眼部疾病','其他系统疾病','中医体质分型','健康评价','健康指导','其他指导','危险因素控制','体检结论' ,'体检建议' ]



# 定义读取Excel文件的函数
# 定义读取Excel文件的函数
def read_excel_file(file_path):
    workbook = xlrd.open_workbook(file_path)
    sheet = workbook.sheet_by_index(0)  # 选择第一个sheet

    file_data = {}

    # 遍历待读取的单元格位置
    for i, (row, col) in enumerate(cells_to_read):
        header = headers[i]
        try:
            # 读取预定位置的内容
            cell_value = sheet.cell_value(row, col)

            found = False
            if isinstance(cell_value, str) and cell_value.strip() == header:  # 如果匹配标题，则向右查找
                # 先在表头所在行查找
                if cell_value.strip() != '脉率' and cell_value.strip() != '呼吸频率':
                    for shift in range(1, 35):  # 最多向右移动35列
                        if found:
                            break
                        next_value = sheet.cell_value(row, col + shift)
                        if next_value and not pd.isna(next_value):  # 避免读取空值
                            # 检查并更新file_data
                            file_data[header] = next_value.replace('\n', ' ') if isinstance(next_value, str) else next_value
                            print(f"找到 {header} 的内容: {next_value}，已更新到 file_data")
                            found = True

                # 如果未找到，再向上、向下查找
                if not found:
                    print(f"未在同一行找到 {header} 的内容，开始上下查找")
                    for row_shift in [-1, 1]:  # 先上下各浮动1行查找
                        if found:
                            break
                        for shift in range(1, 35):  # 最多向右移动35列
                            if found:
                                break
                            if 0 <= row + row_shift < sheet.nrows and 0<= col+shift < sheet.ncols:
                                next_value = sheet.cell_value(row + row_shift, col + shift)
                                if next_value and not pd.isna(next_value):  # 避免读取空值
                                    file_data[header] = next_value.replace('\n', ' ') if isinstance(next_value, str) else next_value
                                    print(f"找到 {header} 的内容: {next_value} 在 ({row + row_shift}, {col + shift})")
                                    found = True
                # 如果上下浮动1行仍未找到，再尝试上下浮动2行
                if not found:
                    for row_shift in [-2, 2]:  # 上下浮动2行查找
                        if found:
                            break
                        for shift in range(1, 35):  # 最多向右移动35列
                            if found:
                                break
                            if 0 <= row + row_shift < sheet.nrows and 0<= col+shift < sheet.ncols:
                                next_value = sheet.cell_value(row + row_shift, col + shift)
                                if next_value and not pd.isna(next_value):
                                    file_data[header] = next_value.replace('\n', ' ') if isinstance(next_value, str) else next_value
                                    print(f"找到 {header} 的内容: {next_value} 在 ({row + row_shift}, {col + shift})")
                                    found = True
            else:  # 如果不匹配，在附近查找
                print(f"表头 {header} 未在预定位置找到，开始在附近查找")
                # 在附近较大范围内查找表头
                for r_shift in range(-40, 41):  # 上下漂移40行范围
                    if found:
                        break
                    for c_shift in range(-10, 11):  # 左右漂移10列范围
                        if found:
                            break
                        if 0 <= row + r_shift < sheet.nrows and 0 <= col + c_shift < sheet.ncols:
                            nearby_value = sheet.cell_value(row + r_shift, col + c_shift)
                            if isinstance(nearby_value, str) and nearby_value.strip() == header:
                                print(f"在 ({row + r_shift}, {col + c_shift}) 找到匹配的表头 {header}")
                                if nearby_value.strip() != '脉率' and nearby_value.strip() != '呼吸频率':
                                    # 找到匹配的标题，向右侧查找
                                    for shift in range(1, 35):  # 最多向右移动35列
                                        if found:
                                            break
                                        if 0 <= row + r_shift < sheet.nrows and 0<= col+shift < sheet.ncols:
                                            next_value = sheet.cell_value(row + r_shift, col + c_shift + shift)
                                            if next_value and not pd.isna(next_value):
                                                file_data[header] = next_value.replace('\n', ' ') if isinstance(next_value, str) else next_value
                                                print(f"找到 {header} 的内容: {next_value} 在 ({row + r_shift}, {col + c_shift + shift})")
                                                found = True                              
                                # 尝试上下查找
                                for row_shift in [-1, 1]:  # 上下浮动1行查找
                                    if found:
                                        break
                                    for shift in range(1, 35):  # 最多向右移动35列
                                        if found:
                                            break
                                        if 0 <= row + r_shift + row_shift < sheet.nrows and 0<= col+c_shift+shift < sheet.ncols:
                                            next_value = sheet.cell_value(row + r_shift + row_shift, col + c_shift + shift)
                                            if next_value and not pd.isna(next_value):
                                                file_data[header] = next_value.replace('\n', ' ') if isinstance(next_value, str) else next_value
                                                print(f"找到 {header} 的内容: {next_value} 在 ({row + r_shift + row_shift}, {col + c_shift + shift})")
                                                found = True
                                for row_shift in [-2, 0, 2]:  # 上下浮动2行查找
                                    if found:
                                        break
                                    for shift in range(1, 35):  # 最多向右移动35列
                                        if found:
                                            break
                                        if 0 <= row + row_shift < sheet.nrows and 0<= col+shift < sheet.ncols:
                                            next_value = sheet.cell_value(row + r_shift + row_shift, col + c_shift + shift)
                                            if next_value and not pd.isna(next_value):
                                                file_data[header] = next_value.replace('\n', ' ') if isinstance(next_value, str) else next_value
                                                print(f"找到 {header} 的内容: {next_value} 在 ({row + r_shift + row_shift}, {col + c_shift + shift})")
                                                found = True
        except IndexError:
            # file_data[header] = None  # 如果超出范围，则设为空
            print(f"filedata[{header}]={None}") # 如果超出范围，则设为空
    return file_data
# 读取文件并保存数据
# 设置输入文件夹路径
folder_path  = 'input_folder_path'
all_data = []

for file_name in os.listdir(folder_path):
    if file_name.endswith('.xls'):
        file_path = os.path.join(folder_path, file_name)
        file_data = read_excel_file(file_path)
        all_data.append(file_data)

# 将所有数据转换为 DataFrame 并导出为Excel文件
df = pd.DataFrame(all_data)
df.to_excel('output.xlsx', index=False)

print("save to output.xlsx")