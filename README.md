Image_Beego_Qiniu
------------

基于Beego的图片上传,采用七牛云存储,这只是个Demo...，演示地址：[http://fp.simman.cc](http://fp.simman.cc)


### 如何安装

1、下载源码

```
go get -u github.com/simman/Image_Beego_Qiniu
```

2、Beego、Bee工具

```
go get -u github.com/astaxie/beego
go get github.com/beego/bee
```

3、Xorm、go-sqlite3

```
go get -u github.com/go-xorm/xorm
go get -u github.com/mattn/go-sqlite3
```

4、无闻的com包

```
go get -u github.com/Unknwon/com
```

5、七牛官方的Go-SDK

```
go get -u github.com/qiniu/api
```

### 运行

```
bee run
```