# freewheelin-go

这个 repo 存放一些我学习 Go 语言的笔记。内容比较杂碎，类似记事本一样。

我感觉 Go 语言是新的 C++，它的设计者 Ken Thompson 同时是 C & Unix 的设计者。

Go 的语法和 C 一样相对简单而优雅，同时解决了 C 中的许多问题：

- 内存管理 --- 垃圾回收
- 组合接口
- 包管理 （ CMake，永远的痛）
- 并发
- 测试（unit testing, benchmark, mocking...）
- 有一套非常完整的标准库，其中包含：http 服务器，grpc，加密库，图像处理....

## 实用工具

- go run \* 运行某个文件，类似脚本语言。
- go build 构建项目
- go fmt 格式化项目
- go vet 检查错误
- go doc 自动生成文档
- go test
  - go test -bench 用 testing.B 做基准测试
  - +build `<tag>` 给测试加上 Tags
  - Table testing
  - Httptest
  - -race 并发竞争错误测试

## 语句

- Go 语言的 switch 自带 break，如果需要 fallthrough 需要手打 fallthrough；而且 case 中可以不是常量表达式。
- 我们可以给 for 语句加标签，continue 可以跳出多层循环，直接跳到标签 for 上去。

## 类型

- Go 语言中的 string 以 utf-8 编码，不能通过索引遍历，需要通过 for range 语句遍历。因为索引是按字节读取的，而 utf-8 中的一个字符未必能由一个字节表示出来。
- Go 语言通常不用数组，用切片。声明切片可以用 make 声明，其中第一个参数为类型，第二个参数为长度，第三个参数为容量，后两个参数是可选的。
- Map 和 切片 都是指针，传递的时候传指针值。
- 可变参数用 ...type 切片作为参数传入的时候要加上 ...
- 结构可以嵌入，将一个类型名写在结构里面，这个结构就可以直接访问那个类型里的东西。
- interface 是 Go 唯一的抽象类型，interface{} 可以当作任意类型。
- 依赖注入 --- 处理 App 各个组件之间的依赖关系。

## 函数

- 可以给函数返回值名称。
- Go 的函数是头等公民，可以构造闭包作为返回值 --- 这和 Python 是一样的
- defer 可以将一个函数在 return 之后执行，通常是 Cleanup 的函数。

## 指针

- 多写纯函数，少做 mutable 的事情。

## 错误

- 用 fmt.Errorf 中的 %w 可以包错误
- 利用一段 defer 的函数可以做到 panic recover

```go
func div60(i int) {
    defer func() {
        if v := recover(); v != nil {
            fmt.Println(v)
        }
    }()
    fmt.Println(60 / i)
}
```

## 模块

- 用 go.mod 处理包依赖关系。可以给 import 的 package 加别名，只需要写在包名的前面。
- 叫做 internal 的 package 只能被 siblings package 访问
- 首字母大写的函数是导出函数，是 API，其他的只能在包内访问
- go mod vendor 可以把外部依赖库也给加上
- 版本管理使用语义版本号，需要能够向后兼容。

## 并发

- 原生并发是 Go 语言的重要特性
- 使用 go func() {} 执行一个并发的 goroutine，需要注意此函数一般为包装函数，参数通常从上下文中捕获（但如果是一个经常变化的变量，需要作为参数)。这个函数不要写在 API 接口里，否则会给调用者带来很大的麻烦。另外，要记得确认这个函数会正常退出。
- channel 是重要的并发同步机制，通常使用无缓存形式。需要手动关闭。
- select 类似 switch，但是 case 是在满足读取条件（事件驱动 event-driven )的诸多条件里随机（防止饥荒)选取一个。
- done channel pattern 用一个 done channel，当完成任务的时候 close(done) 这样，select 里面的 done case 就会可用，返回。
- waitgroups 等待组
- once 只运行一次
- mutex 读写锁

## 值得一提的标准库
