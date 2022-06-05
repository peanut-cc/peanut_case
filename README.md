# 帮助女朋友处理订单

程序监控upload 目录下的文件创建,处理excel

```bash
go mode tidy
go run
```

处理中在windows 中读取文件的时候在个别电脑上会经常出现如下错误
```bash
The process cannot access the file because it is being used by another process
```
对于这个错误没有找到很好的解决办法,暂时是在打开文件时增加打开的次数,现在是会进行10次尝试
如果还是失败,则返回失败
