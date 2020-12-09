`Go` 语言学习笔记。

## Install
对一门语言的快速上手最简单粗暴的方式就是不管三七二十一，先输出个 `hello world` 再说。

### Mac
[点击下载](https://golang.org/dl/go1.15.6.darwin-amd64.pkg) Mac安装包，双击 pkg 文件进行安装即可。

验证是否安装成功：
```
$ go version
go version go1.15.3 darwin/amd64
```

### Linux
[点击下载](https://golang.org/dl/go1.15.6.linux-amd64.tar.gz) Linux 安装包，解压进行安装：
```
$ tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz
```

添加环境变量：

```
export PATH=$PATH:/usr/local/go/bin
```

验证是否安装成功：
```
$ go version 
```

### Windows
[点击下载](https://golang.org/dl/go1.15.6.windows-amd64.msi) Windows安装包，双击 msi 文件进行安装。

验证是否安装成功：
```
$ go version
```

## hello world

新建 `hello_world.go` 文件，添加以下内容：
```
package main
import "fmt"
    func main() {
        fmt.Println("Hello, World!")
    }
```
在当前目录执行：`go run test.go` 。

## 常用命令

|命令|作用|
|---|---|
|go run|编译并执行|
|go build|编译代码包或者源码文件|
|go get|下载第三方代码包并编译安装|
|go install|编译安装main 包和非main 包|
|go env|打印GO语言的环境信息|
|go vet|代码静态检查工具|
