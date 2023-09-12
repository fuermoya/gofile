### 一个非常简单的资源查看器
电脑不方便的时候可以拿手机查阅（目前只支持windows）
支持下载和查看pdf、txt、视频、图片、
记住不要暴露到外网使用！！！

#### 外挂字幕
外挂字幕需要和视频文件同名，目前DPlayer只支持 .vtt 文件（".ass", ".srt", ".ssa", ".stl", ".ttml"转为vtt），某些浏览器不支持外挂字幕，建议使用内嵌字幕

```
- golang版本 v1.20.5
- go mod tidy
- go build -o go-file.exe main.go
```
### 运行
```
直接双击运行 go-file.exe 

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

### 使用了以下开源项目
+ [gin](https://github.com/gin-gonic/gin)
+ [DPlayer](https://github.com/DIYgod/DPlayer)
+ [viewerjs](https://github.com/fengyuanchen/viewerjs)
+ [PDFObject](https://github.com/pipwerks/PDFObject)
+ [pdfh5](https://github.com/gjTool/pdfh5)
+ [layui](https://github.com/layui/layui)
+ [subtitles](https://github.com/asticode/go-astisub)









