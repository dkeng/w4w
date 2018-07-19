# w4w(开发中)
w4w.cc

# 环境搭建
## 使用docker搭建mysql
```
docker run --name mysql -p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=123456 \
-v /host/mysql/logs:/logs \
-v /host/mysql/data:/var/lib/mysql \
-d mysql:5.7.5
```

# Restful Api

## 创建短链接
* 请求路由：`/api/short`
* 请求方式：`POST`
* 请求参数：

|参数名|数据类型|是否必选|说明|
|-------|---------|--------|------|
|url|字符串|是|要缩短的长网址|


* 201 响应参数：

|参数名|数据类型|说明|
|-------|--------|------|
|short_link1|字符串|后台重定向短链接|
|short_link_qrcode1|Base64|后台重定向短链接二维码|
|short_link2|字符串|前台重定向短链接|
|short_link_qrcode2|Base64|前台重定向短链接二维码|

* 400 响应参数：

|参数名|数据类型|说明|
|-------|--------|------|
|short_link1|字符串|后台重定向短链接|
|short_link_qrcode1|Base64|后台重定向短链接二维码|
|short_link2|字符串|前台重定向短链接|
|short_link_qrcode2|Base64|前台重定向短链接二维码|
