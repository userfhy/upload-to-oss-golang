## 关于
> - 基于 Aliyun 官方 sdk 封装，用于上传文件到 oss。
> - 支持从配置文件，或环境变量中读取 `oss` 参数信息。
> - 支持接受命令行文件夹参数。
> - 支持文件/文件夹（及子目录）上传。


## 运行环境

> - Go 1.5及以上。



## 安装方法

> - 执行命令`go get github.com/aliyun/aliyun-oss-go-sdk/oss`获取远程代码包。



## 快速使用

上传 `folderName` 文件夹，为相对路径

```bash
go run uploadToOSS.go folderName
```

静态编译
```bash
CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' upload_oss.go
```


## 上传示例输出

```bash
$ chmod +x uploadToOSS && ./uploadToOSS dist
dist
OSS Go SDK Version:  1.9.1
Need to upload file lists：

dist/index.html
dist/static/css/app.6e652bea9a275362707593477cf5aa43.css
dist/static/fonts/ionicons.05acfdb.woff
dist/static/fonts/ionicons.24712f6.ttf
dist/static/fonts/ionicons.2c2ae06.eot
dist/static/img/background.dafce9d.jpg
dist/static/img/ionicons.621bd38.svg
dist/static/js/app.8cf9c010d98ee3001563.js
dist/static/js/app.8cf9c010d98ee3001563.js.map
dist/static/js/manifest.092b78c69cf6f007609f.js
dist/static/js/manifest.092b78c69cf6f007609f.js.map
dist/static/js/vendor.7244d9bcf45df5b1cec0.js
dist/static/js/vendor.7244d9bcf45df5b1cec0.js.map

Uploading...

Transfer Started, ConsumedBytes: 0, TotalBytes 471.

Transfer Data, ConsumedBytes: 471, TotalBytes 471, 100%.
Transfer Completed, ConsumedBytes: 471, TotalBytes 471.
Transfer Started, ConsumedBytes: 0, TotalBytes 223959.
```



## 提示

> - 使用前请将 `oss.config.example` 文件重命名为 `oss.config`
> - 配置文件中 `readVariable=1` 时读取 `gitlab runner variables`
> - 如上传到根目录请设置 `remoteFolder=/`
