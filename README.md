# thor
调度系统

##TODO
[ ]`manager`管理资源，负责调度`worker`，
[ ]`worker`执行`job`，并检查自身资源的有效性
[ ]资源存储在`boltdb`中
[ ]基于[`kit`](github.com/go-kit/kit)，便于支持`grpc`/`http`等多种协议