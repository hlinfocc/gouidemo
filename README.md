# 介绍

golang fyne 跨平台应用示例，集成web服务，websocket服务

# 构建工具

fyne-cross

# 编译构建

执行`builder.sh`构建，构建完毕后在fyne-cross/bin目录下生成编译好的二进制文件

# fyne-cross 修改

fyne-cross 使用docker构建打包，docker镜像中环境默认从官方下载依赖，可能会不成功，所以需要修改fyne-cross源码，配置GOPROXY代理。

下载源码，

修改internal/command下对应的文件，配置GOPROXY代理

* internal/command/windows.go
* internal/command/linux.go

setupContainerImages方法`for _, arch := range targetArch {`中，
在`image.SetEnv("GOOS", "linux")`或`image.SetEnv("GOOS", "windows")`后面增加

```
    image.SetEnv("GO111MODULE", "on")
	image.SetEnv("GOPROXY", "https://goproxy.cn,direct")
```

修改完后，执行`go build`,将生成`fyne-cross`二进制文件，复制到GOPATH的bin目录下。