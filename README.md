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