#s 写一个爬取全国天气预告的爬虫
import requests
from bs4 import BeautifulSoup
import pandas as pd
import time
import os
import sys
import json
import csv
import datetime
from matplotlib.ticker import MultipleLocator, FormatStrFormatter

#s 1.获取网页源代码
def get_html(url):
    try:
        r = requests.get(url, timeout=30)
        r.raise_for_status()
        r.encoding = r.apparent_encoding
        return r.text
    except:
        return "产生异常"

#s 2.解析网页源代码
def parse_html(html):
    final = []
    soup = BeautifulSoup(html, 'html.parser')
    #s 获取所有的tr标签
    trs = soup.find_all('tr', {'class': ''})
    for tr in trs:
        z = []
        #s 获取所有的td标签
        tds = tr.find_all('td')
        for td in tds:
            z.append(td.string)
        final.append(z)
    return final

#s 3.保存数据
def save_data(data, path):
    #s 1.创建文件对象
    with open(path, 'w', newline='') as f:
        #s 2.创建csv写入对象
        writer = csv.writer(f)
        #s 3.写入数据
        for i in data:
            writer.writerow(i)
        #s 4.关闭文件
        f.close()
    #s 2.创建csv写入对象
    #s 5.返回结果

#s 4.主函数
def main():
    #s 1.获取网页源代码
    url = 'http://www.tianqihoubao.com/lishi/zhengzhou/month/201901.html'
    html = get_html(url)
    #s 2.解析网页源代码
    data = parse_html(html)
    #s 3.保存数据
    path = 'C:/Users/lenovo/Desktop/zhengzhou.csv'
    save_data(data, path)
#s 5.调用主函数
if __name__ == '__main__':
    main()
#s 6.运行程序
#s 7.查看结果
#s 8.提交代码

