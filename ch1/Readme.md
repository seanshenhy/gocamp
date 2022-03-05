# 使用说明
1 sql数据放在data文件夹下，需要本地启动数据库，创建sql库表。
2 克隆代码
```
git clone --depth=1 https://github.com/seanshenhy/gocamp.git
```
3 跳到目录ch1下
```
cd projectName
```
4 修改数据库访问地址
见model目录下的db.go代码文件

5 启动服务
```
go run main.go
```
6 输入URL地址
```
http://127.0.0.1:8080/userinfo
```