# 命令生成mock文件

```
mockgen -destination mock/mock_spider.go -package mock go-demo/mock Spider
```
mockgen -destination 生成文件位置 -package 包名 要mock的接口的位置  接口名


## 使用go generate

在接口文件处添加注释
```go
//go:generate mockgen -destination mock_spider.go -package mock go-demo/mock Spider
```
这样直接在mock文件夹下面执行 `go generate` 即可

## 参数详解

-source： 指定接口文件

-destination: 生成的文件名

-package:生成文件的包名

-imports: 依赖的需要import的包

-aux_files:接口文件不止一个文件时附加文件

-build_flags: 传递给build工具的参数

## 更多

[点击查看更多](https://www.jianshu.com/p/598a11bbdafb)