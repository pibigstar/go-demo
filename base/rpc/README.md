# rpc

- lv1: 最简单的rpc调用
- lv2: 封装server和client到pb中
- lv3: 跨语言rpc调用
- lv4: http上的rpc调用

## lv3
> 将 `rpc.ServeConn`函数用替代 `rpc.ServeCodec`函数,
>其他语言可通过 `json` 格式请求接口方式，请求到 rpc server 中的方法

## lv4 
> 将 `rpc.ServeCodec` 替代为 `rpc.ServeRequest`

**调用**
>  curl localhost:8000/hello -X POST --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
