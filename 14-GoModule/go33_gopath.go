package main

/*
GOPATH 的弊端：
无版本控制概念：比如通过 go get github.com/aceld/zinx 拉代码的时候 无法指定版本，默认直接拉最新的版本
无法同步一致第三方版本号：无法保证其他人期待的第三方依赖库和我们的是一致的
无法指定当前项目引用的第三方版本号：就是说引入项目的 v1、v2 等版本无法指定

GoModule 是1.14版本推荐使用的：
使用 Go Modules 创建项目，建议为了与 GOPATH 分开，不要将项目创建在 GOPATH/src 下
命令              使用
go mod init      生成 go.mod 文件
go mod vendor    导出项目所有的依赖到 vendor 目录
go mod download  下载 go.mod 文件中指明的所有依赖
go mod tidy      整理现有的依赖
go mod graph     查看现有的依赖结构
go mod edit      编辑 go.mod 文件
go mod verify    校验一个模块是否被篡改过
go mod why       查看为什么需要依赖某模块
...

go mod 环境变量
GO111MODULE：是使用 go modules 的开关。
	auto：只要项目包含了 go.mod 文件的话启用 go modules
	on：启用 go modules，推荐设置
	off：禁用 go modules，不推荐
通过 go env -w GO111MODULE=on 来设置
GOPROXY：表示导包的时候自动从哪个路径去下包，默认是国外的 https://proxy.golang.org,direct  如果想弄快点，就设置成阿里云的 https://mirrors.aliyun.com/goproxy/
GOSUMDB：校验从网络上拉取第三方库的时候是一个完整的库，不是经过篡改过的。通过配置 GOPROXY，GOSUMDB 就默认是 GOPROXY 的路径了
GONOPROXY/GONOSUMDB/GOPRIVATE：哪些不需要代理的 / 哪些不需要校验的 / 哪些是私有的即公司内部的，通常配置 GOPRIVATE 就可以了，另外两个会自动执行 GOPRIVATE


通过 Go Modules 开启一个项目：
1. 开启 GO111MODULES=on
2. 初始化项目：任意文件夹创建一个项目（不要求在$GOPATH/src） mkdir ProjectName/ModulesName。
3. 创建 go.mod 文件，同时起当前项目的模块名称，比如 go mod init github.com/xileijia/test 表示其他人在导我的包的时候就是写的这个。（或者直接用 goland 打开，就会自动生成go.mod文件）
4. 创建之后就会生成 go.mod 文件，有两行，分别是 模块名称、go 的版本号
5. 在该项目中编写源代码。如果源代码中依赖某个库，比如依赖 github.com/aceld/zinx/znet，需要 down 下来（手动 down：go get github.com/aceld/zinx/znet；自动 down：直接 go mod tidy 就会自动 down）
6. down 完之后 go.mod 文件中会添加一行新代码：require github.com/aceld/zinx v0.0.1-20200315073925-f09df55dc746 // indirect
	表示当前模块依赖 require 后面的模块，版本号是v0.0.1..., //indirect表示间接依赖，因为项目直接依赖的是znet包，znet包是zinx包的子包，所以对于 zinx 是间接依赖
7. 之后会生成一个 go.sum 文件，会罗列当前项目直接或间接依赖所有模块版本，保证今后项目依赖的版本不会被篡改。有两种形式(h1:hash 表示整体项目的zip文件打开之后的全部文件的校验和生成的hash。如果不存在，表示依赖的库可能用不上)
	(xxx/go.mod h1:hash 表示对 go.mod 文件做的hash)


如果我们的项目引入了 github.com/aceld/zinx v0.0.1-20200315073925-f09df55dc746 也就是0315这个版本，现在我们想用之前的版本，使用 replace 命令：
go mod edit -replace=现在的版本=需要替换的版本   使用这个命令之后版本就会替换
go.mod 文件就会被修改，在 reuqire 之后添加一行 replace，表示 require 的依赖版本被替换成了 现在替换之后的版本
*/
