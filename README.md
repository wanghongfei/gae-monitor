# Monitor

曝光监测服务。

接收曝光请求 -> 记日志 -> 投递kafka消息 -> 返回1像素透明gif图。



## 为什么不用Nginx

现在曝光监控的普遍做法是将1*1像素图片放到nginx下，然后通过nginx的access.log监测曝光。这么做的成本是：

- 编译nginx(也可直接下载预编译版)
- 配置nginx

如果不想调整nginx日志格式，也不想直接在lua中操作kafka:

- 另启动一个进程监控nginx访问日志
- 将访问日志格式修改成后端业务需要的格式
- 投放消息到kafka中

或者，使用一种专门的日志传输服务将访问日志传到目标位置，但需要自研基础设施。

如果想调整nginx日志格式，直接在lua中完成kafka消息发送：

- 编写并调试lua脚本
- 配置nginx

   

可以看到，比较麻烦。使用go专门写一个监测服务有好处有：

- 能够添加任意业务逻辑，且可维护性强(相比lua脚本)
- 部署方便，直接执行`./monitor`即可(免去nginx配置)
- 不会比nginx差很多的性能(协程、停顿<1ms的垃圾回收,、不像Java一样需要等JIT编译)

## 构建运行

将项目clone到`$GOPATH`下，然后下载依赖：

```
godep restore
```

编译：

```
go build gaemonitor
```

确保`config.json`与编译后的可执行文件在同一目录，然后就可以直接运行了：

```
./geamonitor
```

