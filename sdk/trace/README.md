# 链路跟踪

## demo
> 简单示例

### 1. 安装 jaeger
doc: [点击查看](https://www.jaegertracing.io/docs/1.17/getting-started)
```bash
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.17
```

### 2. 启动demo
1. 启动server
2. 启动client
3. 查看 http://localhost:16686

## app
> 实际使用，gin、gorm、grpc 集成trace

1. 启动 app/main.go
2. 浏览器访问：localhost:8081/hello
3. 浏览器访问：localhost:16686 查看链路