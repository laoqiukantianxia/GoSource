# 编译
1. go 命令
* run：编译一个或多个以`.go`结尾文件，链接库文件，并运行文件
* build：编译程序，生成可执行程序
2. 程序介绍：
* go代码通过包(package)组织，一个包由单个目录下的一个或多个.go文件组成，目录定义包的作用
* 每个源文件以`package`声明开始，表示该文件属于哪个包
*  `main`是特殊包，定义的第一个独立可执行的程序，而不是一个库。`mian`包中的`main`函数是整个程序的入口。

