# goproject
goproject

# 使用go mod和go proxy 
在当前操作系统中添加如下两个环境变量
GO111MODULE = on
GOPROXY=https://goproxy.io

# 新建项目goproject
go mod init goproject
每次go.mod得更新，go build命令创建一个名为的go.sum文件，go.sum其中包含特定模块版本内容的预期加密校验和。

# go mod的一些注意事项
1. go mod init name中的name可以随便取
2. 使用了 Go modules 后所有的 import path 都得以 module path 开头，当前工作目录的话就以步骤 1 中的 module path 开头；
3. 如果你指的“其他目录”是别的模块的，同理步骤 2，如果指的是普通文件夹，那得用 `go mod replace` 替换为你的目标路径模块；
   1. 比如依赖同级目录的xx文件夹，则在go.mod文件中添加如下: replace xx => ../xx
4. 现在大多数 IDE 如果它有现成的 Go 插件的话那么可以直接使用，插件会自动处理，你只需将 `GO111MODULE=on` 即可。


# 日志文件
log.go

# 测试文件

