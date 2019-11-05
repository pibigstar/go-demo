# Go性能分析

# 1. 准备工作
## 1.1 下载go-wrk
>  这个是用来进行http接口压测的,
>  官网地址：https://github.com/tsliwowicz/go-wrk
>
[七牛云下载](http://store.pibigstar.com/go-wrk.exe)

**使用**
```bash
go-wrk -d 500 http://localhost:8080/hello
```

## 1.2 安装生成火焰图工具
### 1.2.1 下载go-torch
```bash
go get github.com/uber/go-torch
```
### 1.2.2 安装perl
官网下载地址：https://www.activestate.com/products/perl/downloads/

[七牛云下载](http://store.pibigstar.com/ActivePerl-5.28.1.0000-MSWin32-x64-92425271.msi)

`注意`：安装时记得把添加到环境变量PATH选项勾上
### 1.2.3 下载FlameGraph
```bash
git clone https://github.com/brendangregg/FlameGraph.git
```
`注意`：将此文件夹加入到`PATH`环境变量中

## 1.3 下载graphviz
> 生成SVG 图，能够更好地看到函数调用 CPU 占用情况
### 1.3.1 Windows安装

> 官网下载地址：https://graphviz.gitlab.io/_pages/Download/Download_windows.html

[七牛云下载](http://store.pibigstar.com/graphviz-2.38.msi)

安装好之后，添加其 bin文件到`PATH`环境变量中

### 1.3.2 Linux安装
```bash
# Ubuntu
sudo apt-get install graphviz

# Centos
yum install graphviz
```
### 1.3.3 测试
> 检查是否安装成功
```bash
dot -version
```

# 2. 性能分析
## 2.1 开启性能分析
开启性能分析只需要导入下面这个包即可
```go
import _ "net/http/pprof"
```


## 2.2 开始压测
**代码见文章末尾附录**
1. 启动 `main.go`
2. 开始压测
```bash
go-wrk -d 500 http://localhost:8080/hello
```

## 2.3 web查看
浏览器打开：[http://localhost:8080/debug/pprof/](http://localhost:8080/debug/pprof/)

![](https://img-blog.csdnimg.cn/20191104171603929.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2p1bm1veGk=,size_16,color_FFFFFF,t_70)
- `allocs`: 内存分配情况
- `block`: 导致阻塞同步的堆栈跟踪
- `cmdline`: 当前程序激活的命令行
- `goroutine`: 当前运行的goroutine
- `heap`: 存活对象的内存分配情况
- `mutex`: 互斥锁的竞争持有者的堆栈跟踪
- `profile`: 默认进行 30s 的 CPU Profiling
- `threadcreate`:  操作系统线程跟踪
- `trace`: 当前程序执行情况
- `full goroutine stack dump`: 所有goroutine栈输出
## 2.4 采样分析
下面命令会打开一个交互页面
```bash
go tool pprof --seconds 5 http://localhost:8080/debug/pprof/profile
```
**常用命令**
- `top`  默认显示 10 个占用 CPU 时间最多的函数
- `tree` 树形结构查看goroutine情况。
- `web` 生成SVG函数调用图（需安装`graphviz`）
- `exit` 退出分析

打开`pprof`可视化页面，与上面 可交互页面中`web`命令效果一样
```bash
go tool pprof -http=:8080 http://localhost:8080/debug/pprof/profile?seconds=60
```
![](https://img-blog.csdnimg.cn/20191104181816875.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2p1bm1veGk=,size_16,color_FFFFFF,t_70)

## 2.5 生成火焰图
> 进行 5 秒钟的 CPU 性能采样并生成火焰图
```bash
go-torch --seconds 5 http://localhost:8080/debug/pprof/profile
```
命令执行完会在该目录下生成一个`torch.svg`文件，可用浏览器打开查看火焰图
![](https://img-blog.csdnimg.cn/2019110417565173.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2p1bm1veGk=,size_16,color_FFFFFF,t_70)
# 3. Go压力测试分析
1.生成 `pprof.cpu`文件
```bash
go test -bench . -benchmem -cpuprofile pprof.cpu
```
2.分析`pprof.cpu`文件
```go
go tool pprof pprof.cpu
```