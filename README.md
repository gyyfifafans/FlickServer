# FlickServer

##### 在 config/app.ini 配置mysql服务器信息
##### 进入mysql，创建flickdb数据库
```
create database flickdb;
```
##### 编译FlickServer，通过命令行创建数据库表
```
./FlickServer -syncdb
```
##### 启动服务器
```
./FlickServer
```
##### 使用Postman对本地服务器进行测试