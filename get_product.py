# -*- coding: utf-8 -*-

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

import json

print(json.dumps(res, ensure_ascii=False))