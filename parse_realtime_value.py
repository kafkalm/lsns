# -*- coding: utf-8 -*-

with open('lsns/realtime_value.table') as f:
    data = f.read()

data = data.split('\n')
data = [l.split('\t') for l in data][2:]

def get_int(v):
    try:
        return int(v)
    except Exception:
        return 0

base_profit_idx = 11
profit_offset = 6
cities = ["修格里城",
        "铁盟哨站",
        "7号自由港",
        "澄明数据中心",
        "阿妮塔战备工厂(lv40)",
        "阿妮塔能源研究所",
        "荒原站",
        "曼德矿场",
        "淘金乐园"]

product_city_profit = {}
for line in data:
    product_name = line[1]
    city_profit = {}
    for i in range(len(cities)):
        city = cities[i]
        city_profit[city] = get_int(line[base_profit_idx + profit_offset * i])
    product_city_profit[product_name] = city_profit

city_product_profit = {}
for product_name, city_profit in product_city_profit.items():
    for city_name, profit in city_profit.items():
        if city_name not in city_product_profit:
            city_product_profit[city_name] = {}
        city_product_profit[city_name][product_name] = profit


with open('lsns/product.table') as f:
    data = f.read()

data = data.split('\n')
data = [l.split('\t') for l in data]

# template = {
#     "name": "修格里城",
#     "local_product_names": ["发动机", "弹丸加速装置", "红茶", "修格里严选礼包"],
#     "inventory": {
#         "product_quantity": {
#             "发动机": 9,
#         },
#         "product_info": {
#             "item_1": {
#                 "name": "item_1",
#                 "price": 100
#             }
#         }
#     }
# }

res = []
i = 0
while i < len(data):
    row = data[i]
    t = {
        "name": row[0],
        "local_product_names": [row[1]],
        "inventory": {
            "product_quantity": {
                row[1]: int(row[2])
            },
            "product_info": {
                row[1]: {
                    "name": row[1],
                    "price": int(row[3])
                }
            }
        }
    }
    j = i+1
    while j < len(data):
        row = data[j]
        if row[0] != '':
            break
        if row[2] == 'None' or row[3] == 'None':
            j+=1
            continue
        t['local_product_names'].append(row[1])
        t['inventory']['product_quantity'][row[1]] = int(row[2])
        t['inventory']['product_info'][row[1]] = {
                    "name": row[1],
                    "price": int(row[3])
                }
        j+=1
    res.append(t)
    i = j

def find_city(cities, product_name):
    for city in cities:
        if product_name in city['local_product_names']:
            return city

for city in res:
    for product_name, profit in city_product_profit[city['name']].items():
        if product_name == '弹丸加速装置-修城' or product_name == '弹丸加速装置-哨站':
            product_name = '弹丸加速装置'
        local_city = find_city(res, product_name)
        if local_city is None:
            continue
        if local_city['inventory']['product_info'][product_name]['price'] is None:
            continue
        city['inventory']['product_info'][product_name] = {
            "name": product_name,
            "price": local_city['inventory']['product_info'][product_name]['price'] + profit
        }

import json

j = json.dumps(res, ensure_ascii=False)
with open('lsns/city.json', 'w') as f:
    f.write(j)