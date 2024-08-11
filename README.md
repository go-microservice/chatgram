# chatgram

## Layer

The application with the following layers:

- 1. HTTP Layer: Responsible for handling incoming HTTP requests and returning responses.
- 2. Service Layer: Responsible for handling business logic and communicating with the data layer.
- 3. Data Layer: Responsible for handling communication with other micro-service.

## 接口文档

http://localhost:8081/docs

## 常用命令

```bash
# 生成 proto 和 grpc 文件
make grpc

# 生成 http 结构
make http

# 生成文档
make openapi

# 运行服务
make run
```
