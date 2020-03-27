# RPC

##  RPC 需要满足什么条件
> func (t *T) MethodName(argType T1, replyType *T2) error

- 方法类型（T）是导出的（首字母大写）
- 方法名（MethodName）是导出的
- 方法有2个参数(argType T1, replyType *T2)，均为导出/内置类型
- 方法的第2个参数一个指针(replyType *T2)
- 方法的返回值类型是 error


## 实例
- lv1: 最简单的rpc调用
- lv2: 封装server和client到pb中
- lv3: 跨语言rpc调用
- lv4: http注册，通过URL访问rpc server中方法
- lv5: 基于HTTP，异步调用, 证书认证

### lv1
> 直接通过 rpc.Register 注册rpc服务，client 通过 rpc.Dial 获取client调用server方法

**缺点**
> 需要知道server层的方法名才可以使用，方法名容易写错

### lv2
> 将server的注册与client的获取与调用封装到`pb`里，解决方法调用容易出错问题

**缺点**
> 无法跨语言调用，用go写的rpc方法，不能用java去调用


### lv3
> 将 `rpc.ServeConn`函数用替代 `rpc.ServeCodec`函数,
>其他语言可通过 `json` 格式请求接口方式，请求到 rpc server 中的方法
> 解决了跨语言调用的问题

**缺点**
> 无法通过URL直接方法，使用上不是很方便

### lv4
> 将 `rpc.ServeCodec` 替代为 `rpc.ServeRequest`, 让其可直接解析URL，达到直接调用rpc server中的方法

**调用**
```bash
curl localhost:8000/hello -X POST --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
```
**缺点**
> 没有异步调用和证书认证等方式，不是很安全

### lv5
> 添加证书和异步调用方式

**生成证书**
```bash
# 生成私钥
openssl genrsa -out server.key 2048
# 生成证书
openssl req -new -x509 -key server.key -out server.crt -days 3650
# 只读权限
chmod 400 server.key
```
**注意**
> 生成证书时，Common Name 写 localhost