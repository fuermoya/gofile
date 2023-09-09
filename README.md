### 一个非常简单的资源查看器
主要是因为电脑不方便的时候拿手机看（目前只支持windows）
支持下载和查看视频、图片、
记住不要暴露到外网使用！！！
```
- golang版本 v1.20.5
- go mod tidy
- go build -o go-file.exe main.go
```
### 运行
```
直接双击运行 gofile.exe 

PowerShell后台运行和关闭
 - Start-Process -WindowStyle hidden -FilePath go-file.exe
 - taskkill /f /t /im go-file.exe
 
创建Windows服务(管理员运行cmd) 静态资源已内嵌
 - cd /d exe文件的目录
 - go-file.exe install
 - go-file.exe start
 - go-file.exe stop
 - go-file.exe uninstall
```

感谢以下开源项目
```
https://github.com/DIYgod/DPlayer
https://github.com/fengyuanchen/viewerjs
https://github.com/pipwerks/PDFObject
https://github.com/layui/layui
```




