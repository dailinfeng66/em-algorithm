import numpy as np

import copy


def is_equal(data1, data2):
    for i1 in range(2):
        for j1 in range(5):
            if abs(data1[i][j] - data2[i][j]) > 0.005:
                return False
    return True


# 原始数据
data = {
    "T100": [1, 1, 0, 0, 1],
    "T200": [0, 1, 0, 1, 0],
    "T300": [0, 1, 1, 0, 0],
    "T400": [1, 1, 0, 1, 0],
    "T500": [1, 0, 1, 0, 0],
    "T600": [0, 1, 1, 0, 0],
    "T700": [1, 0, 1, 0, 0],
    "T800": [1, 1, 1, 0, 1],
    "T900": [1, 1, 1, 0, 0]
}

# 分类结果概率
data_r_p = {
    "T100": [1.0, 1.0],
    "T200": [1.0, 1.0],
    "T300": [1.0, 1.0],
    "T400": [1.0, 1.0],
    "T500": [1.0, 1.0],
    "T600": [1.0, 1.0],
    "T700": [1.0, 1.0],
    "T800": [1.0, 1.0],
    "T900": [1.0, 1.0],
}
# 分类结果
data_r = {
    "T100": 1,
    "T200": 1,
    "T300": 1,
    "T400": 1,
    "T500": 1,
    "T600": 1,
    "T700": 1,
    "T800": 1,
    "T900": 1,
}

# 初始 p
p1 = 0.6
p2 = 0.4

# 初始化原始概率
data_p = np.zeros(shape=(2, 5))
data_p[0][0] = 0.6
data_p[0][1] = 0.6
data_p[0][2] = 0.3
data_p[0][3] = 0.5
data_p[0][4] = 0.3

data_p[1][0] = 0.4
data_p[1][1] = 0.4
data_p[1][2] = 0.7
data_p[1][3] = 0.5
data_p[1][4] = 0.7

num = 0
while True:
    num += 1
    print("第" + str(num) + "轮")
    print("传入的p：")
    print("p1:" + str(p1))
    print("p2:" + str(p2))
    data_p_t = copy.deepcopy(data_p)
    for out_key in data:
        #  [1, 1, 0, 0, 1]
        value = data[out_key]
        # 如果是第一类的概率
        p1_t = 1 * p1
        # 如果是第一类的概率
        p2_t = 1 * p2

        for i in range(5):
            if value[i] == 1:
                p1_t *= data_p[0][i]
                p2_t *= data_p[1][i]
            else:
                p1_t *= (1 - data_p[0][i])
                p2_t *= (1 - data_p[1][i])
        # 更新分类结果概率
        data_r_p[out_key][0] = p1_t / (p1_t + p2_t)
        data_r_p[out_key][1] = p2_t / (p1_t + p2_t)
    # 更新p
    p1_add = 0
    p2_add = 0
    for out_key in data_r_p:
        value = data_r_p[out_key]
        p1_add += value[0]
        p2_add += value[1]
    p1_new = p1_add / 9
    p2_new = p2_add / 9
    # 更新data_p_t
    for i in range(2):
        for j in range(5):
            up = 0
            down = 0
            for out_key in data:
                up += data_r_p[out_key][i] * data[out_key][j]
                down += data_r_p[out_key][i]
            data_p_t[i][j] = up / down

    # p值不变进行输出分类结果

    if abs(p1_new - p1) <= 0.005 and abs(p2_new - p2) <= 0.005 and is_equal(data_p_t, data_p):
        for out_key in data_r:
            if data_r_p[out_key][0] > data_r_p[out_key][1]:
                data_r[out_key] = 1
            else:
                data_r[out_key] = 2
        break
    p1 = p1_new
    p2 = p2_new
    data_p = copy.deepcopy(data_p_t)
    # data_r = copy.deepcopy(data_r_t)
print(data_r)
