# 链路跟踪

## 1. 安装 jaeger
```bash
docker run -d -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
```

## 2. 启动demo
1. 启动server
2. 启动client
3. 查看 http://localhost:16686