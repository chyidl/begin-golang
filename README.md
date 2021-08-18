Golang
======

* 软件开发挑战
```
1. 多核硬件架构
2. 超大规模的分布式集群
3. Web模式导致的前所未有的开发规模和更新速度
```

*  Go语言简述
```
# Go 创始人:
    Rob Pike: - Unix 早期开发者 UTF-8创始人
    Ken Thompson: Unix 创始人，C语言创始人
    Robert Griesemer: Google V8 JS Engine, JVM Hot Spot 开发者

1. 简单
2. 高效:
3. 生产力
    Go语言支持复合
```

* 开发环境构建
```
# GOROOT: Go语言安装根目录的路径
# GOBIN: Go程序生成的可执行文件 executable file
# GOPATH: 若干工作区目录的路径
    1. 1.8版本以前必须设置这个环境变量
    2. 1.8版本以后，如果没有设置会使用默认值
        Unix默认 $HOME/go
# src: 
# pkg: 归档文件.a扩展名
# bin: 可执行文件

➜ go version
go version go1.16.6 linux/arm64

src/ch1/main via 🐹 v1.16.6
➜ go run hello_world.go
Hello World
```

* 语言基本知识
```
# 应用程序入口:
    1. 必须是main包: package main
    2. 必须是main方法: func main()
    3. 文件名不一定是main.go

# 退出返回值
    Go 中main函数不支持任何返回值
    os.Exit 返回状态

# 获取命令行参数
    main 函数不支持传入参数
    在程序中直接通过os.Args获取命令行参数

# 编写测试程序
    1. 源码文件以 _test结尾: xxx_test.go
    2. 测试方法名 Test开头: func TestXXX(t *testting.T) {...}
    $ go test -json xxx_test.go

# 变量赋值
    格式代码统一规范
    赋值可以进行自动类型推断 var a int = 1
    在一个赋值语句中可以对多个变量进行同时赋值

# 常量定义
    快速设置连续值

# 基本数据类型
    bool
    string
    int int8 int16 int32 int64
    uint uin8 uint16 uint32 uint64 uint64ptr
    byte // alias for uint8
    rune // alias for int32, represents a Unicode code point
    float32 float64
    complex64 complex128

# 类型转换
    Go 语言不允许隐式类型转换
    别名和原有类型也不能进行隐式类型转换 (1.16 貌似支持别名类型转换)

# 类型的预定义值
    math.MaxInt64
    math.MaxFloat64
    math.MaxUint32

# 指针类型
    1. 不支持指针运算
    2. string是值类型，默认的初始化值为空字符串，不是nil

# 算术运算符
    +
    -
    *
    /
    %   求余 : B % A
    ++  自增
    --  自减

    Go语言没有前置的++,--

# 比较运算符
    ==
        数组比较
        相同纬数且含有相同个数元素的数组才可以比较
        每个元素都相同才相同
    !=
    >
    <
    >=
    <=

# 逻辑运算符
    && : 两边都为True
    || :
    !  :

# 位运算符
    &
    |
    ^
    <<
    >>

    &^ 按位置零
    1 &^ 0 -- 1
    1 &^ 1 -- 0
    0 &^ 1 -- 0
    0 &^ 0 -- 0

# 循环
    Go语言仅支持循环关键字 for

# if 条件语句
    1. condition 表达式结果必须为布尔值
    2. if 条件语句 支持变量赋值

# switch 条件
    条件表达式不限制为常量或者整数
    单个case中，可以出现多个结果选项，使用逗号分隔
    与C语言等规则相反，Go语言不需要使用break来明确退出一个case
    可以不设定switch之后的条件表达式，在此情况下，整个switch结构与多个if...else...的逻辑作用等同

# 数组声明
    var a [3]int // 声明并初始化为默认零值
    b := [3]int{1,2,3} // 声明同时初始化

# 数组截取
    a[开始索引(包含), 结束索引(不包含)]

# 数组内部结构
    结构体:
        ptr: 指向连续的地址空间
        len: 元素个数
        cap: 内部数组的容量

# 切片共享存储结构
# 数组 vs. 切片
    1. 容量是否可伸缩 // 数组不可伸缩
    2. 是否可以进行比较 //

# Map 声明
    > 不支持并发读写
    m := map[string]int{"one":1, "two":2}
    m := map[string]int{}
    // 自增
    m := make(map[string]int, 10) // initial capacity 初始化容量

# Map元素访问
    在访问Key不存在时，仍然会返回零值，不能通过返回nil来判断元素是否存在

# Map 遍历
    for range

# Map 与工厂模式
    Map 的值可以是一个方法
    与Go的Dock Type接口方式一起，可以方便实现单一方法对象的工厂模式

# Map 实现Set
    Go 的内置集合中没有Set实现，可以map[type]bool
    1. 元素的唯一性
    2. 基本操作
        1. 添加元素
        2. 判断元素是否存在
        3. 删除元素
        4. 元素个数

# 字符串
    1. string 是数据类型，不是引用或者指针类型
    2. string是只读的byte slice, len函数可以包含的byte数
    3. string的byte数组可以存放任何数据
    不可变类型

# Unicode UTF8
    Unicode 是一种字符集 code point
    UTF8是unicode的存储实现

# 常用的字符串函数
    strings 包:
    strconv 包:

# 函数是一等公民
    1. 可以有多个返回值
    2. 所有参数都是值传递 slice, map, channel
    3. 函数可以作为变量的值
    4. 函数可以作为参数和返回值

# 可变传递参数
# defer 函数: 延迟执行函数


# Go面向对象编程

Is Go an object-oriented language?
> Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concept of "interface" in Go provides a different approach that we believe is easy to use and in some ways more general.
Also, the lack of a type hierarchy makes "objects" in Go feel much more lightweight than in languages such as C++ or Java

# 封装数据和行为
# 结构体定义
type Employee struct {
    Id string
    Name string
    Age int
}

# 行为(方法)定义

# 接口
    # Duck Type 式接口实现
        接口定义
        接口实现 -- 方法的签名一致
    1. 接口为非入侵性，实现不依赖与接口定义
    2. 所有接口定义可以包含在接口使用者内
    3. 空接口对象可以代表任何类型
    obj interface{}

# 接口变量
    var prog Coder = &GoProgrammer{}
        类型: type GoProgrammer struct {}
        数据: &GoProgrammer

# 自定义类型
# 扩展
# LSP: 所有父类使用的场景，我们可以使用子类代替
# 多态:
# 空接口与断言
    1. 空接口可以表示任何类型
    2. 通过断言来将空接口转换为特定类型
        v,ok := p.(int) // ok=true 为转换成功
# Go 接口最佳实践
    1. 倾向于使用小的接口定义，很多接口只包含一个方法
    type Reader interface {
        Read(p []byte) (n int, err error)
    }
    type Writer interface {
        Write(p []byte) (n int, err error)
    }

    2. 较大的接口定义，可以由多个小接口定义组合而成
    type ReadWriter interface {
        Reader
        Writer
    }

    3. 只依赖必要功能的最小接口
    func StoreData(reader Reader) error {

    }

# 错误机制
    1. 没有异常机制
    2. error类型实现了error接口
    3. 可以通过errors.New 来快速创建错误实例

    type error interface {
        Error() string
    }

    errors.New("n must be in the range[0,1]")

# Panic
    1. panic 用于不可恢复的错误
    2. panic 退出前会执行defer指定的内容

# os.Exit
    1. os.Exit 退出时不会调用defer指定的函数
    2. os.Exit 退出时不输出当前调用栈信息
# recover “错误恢复”
    def func() {
        if err := recover(); err != nil {
            log.Error("recovered panic", err)
        }
    }()

    1. 形成僵尸服务进程，导致health check失效
    2. Let it Crash 往往是恢复不确定性错误最好的方法

# package
    1. 基本服用模块单元
        以首字母大写来表明被包外代码访问
    2. 代码的package可以和所在的目录不一致
    3. 同一目录里的go代码的package要保持一致

# remote package
    1. 通过 go get 获取远程依赖
        go get -u 强制从网络更新远程依赖
    2. 注意代码在GitHub上的组织形式，使用go test
        直接从代码路径开始，不要有src

# init 方法 -- 初始化方法
    1. 在main被执行前，所有依赖的package的init方法都会被执行
    2. 不同包的init函数按照包导入的依赖关系决定执行顺序
    3. 每个包可以有多个init函数
    4. 包的每个源文件也可以有多个init函数

# 依赖管理
    -- 未解决问题
        1. 同一环境下，不同项目使用同一包的不同版本
        2. 无法管理对包的特定版本的依赖
    vendor 路径
        Go 1.5 release版本发布，vendor目录被添加到除了GOPATH和GOROOT之外的依赖目录查找解决方案
        -- 查找依赖包路径的解决方案
            1. 当前包下的vendor目录
            2. 向上级目录查找，知道找到src下的vendor目录
            3. 在GoPATH下面查找依赖包
            4. 在GOROOT目录下查找
```

* Go 并发
```
# Thread vs. Groutine
线程 - 协程
    1. 创建默认的stack的大小
        JDK 5以后Java Thread Stack 默认为1M
        Groutine 的stack初始化大小为2K
    2. 和 Kernel Space Entity 的对应关系
        Java Thread 是 1:1
        Groutine 是 M:N

user space thread sheduler:
kernel space sheduler:

M: System Thread
P: Processor
G: Goroutine : 协程对象

# 共享内存并发机制
    Lock:
        sync.Mutex
            互斥锁
            Lock() -- 加锁
            Unlock() -- 去锁
        sync.RWLock
            读锁
            写锁分开

    sync.WaitGroup -- 等待线程完成
        Add()
        Done()
        Wait

# CSP communicating sequential processes 并发机制
    -- CSP vs. Actor
        1. 和Actor直接通讯不同，CSP模式则是通过Channel进行通讯，更松耦合一些
        2. Go 中的channel是有容量限制并且独立处理Groutine, erlang 的Actor模式中的mailbox容量是无限，接收进程总是被动处理消息

# Channel:
    no buffer channel: 无缓存
    buffer channel: 缓存

    -- channel 关闭
        1. 向关闭的channel发送数据，会导致panic
        2. v,ok <- ch; ok 为bool值，true表示正常接收，false表示通道关闭
        3. 所有的channel接收者都会在channel关闭时，立刻从阻塞等待中返回并且 ok=fasle, 这种广播机制常被利用，进行向多个订阅者同时发送信号

# 多路选择和超时控制
    select:
        多渠道选择
        select {
            case ret := <-retCh1:
                t.Logf("result %s", ret)
            case ret := <-retCh2:
                t.Logf("result %s", ret)
            default:
                t.Error("No one returned")
        }
        超时控制
        select {
            case ret := <-retCh:
                r.Logf("result %s", ret)
            case <-time.After(time.Second * 1):
                t.Error("time out")
        }

# 任务的取消
# Context 与 任务取消
    -- 关联任务的取消
    1. 根Context 通过context.Background()创建
    2. 子Context context.WithCancel(parentContext) 创建
        ctx, cancel := context.WithCancel(context.Background())
    4. 当前Context 被取消时，基于子context都会被取消
    5. 接收取消通知 <- ctx.Done()
```

* 并发任务
```
# 仅执行一次
    单例模式 (懒汉式，线程安全)
        double check 机制 二次验证

# 有一个成功即成功

# 所有任务都完成

# 对象池
    -- 使用buffered channel 实现对象池
    slow response 非常可怕
    使用不同的池缓存不同的对象

# sync.Pool 对象缓存
    Processor:
        私有对象 -- 协程安全
        共享池 -- 协程不安全 需要锁
    -- 对象获取
        1. 尝试从私有对象获取
        2. 私有对象不存在，尝试从当前processor的共享池获取
        3. 如果当前Processor共享池也是空的，那么就尝试去其他Processor的共享池获取
        4. 如果所有子池都是空的，最后就用用户指定的New函数产生一个新的对象返回
    -- 对象放回
        1. 如果私有对象不存在则保存为私有对象
        2. 如果私有对象存在，放入当前Processor子池的共享池中
    -- sync.Pool 对象的声明周期
        1. GC 会清除sync.pool 缓存的对象
        2. 对象的缓存有效期为下一次GC之前
    -- sync.Pool 总结
        适合于通过复用，降低复杂对象的创建和GC代价
        协程安全，会有锁的开销
        生命周期受GC影响，不适合于做连接池，需要自己管理声明周期的资源的池化
```

* Go 测试
```
# 单元测试
    Fail, Error: 该测试失败，该测试继续，其他测试继续
    FailNow, Fatal: 该测试失败，该测试中止，其他测试继续执行
    -- 代码覆盖率
        go test -v - cover
    -- 断言
        https://github.com/stretcher/testify

# Benchmark
    go test -bench=. -benchmem
        -bench=<相关benchmark测试>
        -benchmem

    begin-golang/src/ch23/benchmark on  master [!+?] via 🐹 v1.16.6
    ➜ go test -bench=.
    goos: darwin
    goarch: amd64
    pkg: begin-golang/src/ch23/benchmark
    cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
    BenchmarkConcatStringByAdd-4           	 6788464	       169.4 ns/op
    BenchmarkConcatStringByBytesBuffer-4   	12485692	        94.15 ns/op
    PASS
    ok  	begin-golang/src/ch23/benchmark	2.861s

    begin-golang/src/ch23/benchmark on  master [!+?] via 🐹 v1.16.6 took 3s
    ➜ go test -bench=. -benchmem
    goos: darwin
    goarch: amd64
    pkg: begin-golang/src/ch23/benchmark
    cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
    BenchmarkConcatStringByAdd-4           	 7047180	       164.6 ns/op	      16 B/op	       4 allocs/op
    BenchmarkConcatStringByBytesBuffer-4   	12206203	        91.85 ns/op	      64 B/op	       1 allocs/op
    PASS
    ok  	begin-golang/src/ch23/benchmark	3.191s


# BDD - Behavior Driven Development
    让业务领域的专家参与开发
    用业务领域的语言来描述
    tests are also good documents
    验收框架

    goconvey : github.com/smartystreets/goconvey/convey
    # 启动WEB UI
        $ GOPATH/bin/goconvey
```

* 高级编程
```
# 反射编程
    reflect.TypeOf: 返回类型(reflect.Type)
    reflect.ValueOf: 返回值(reflect.Value)
    可以从reflect.Value获得类型
    通过kind判断类型

    1. 按照名字访问结构的成员
        reflect.ValueOf(*e).FieldByName("Name")
    2. 按照名字访问结构的方法
        reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})

    Struct Tag:
    type BasicInfo struct {
        Name string `json:"name"`
        Age int `json:"age"`
    }

# DeepEqual
# "反射注意事项"
    1. 提高程序的灵活性
    2. 降低程序的可读性
    3. 降低程序的性能

# 不安全编程 unsafe
    i := 10
    f := *(*float64)(unsafe.Pointer(&i))
```

* 架构模式
```
23种设计模式

# Pipe-Filter Pattern
    Pump -> Filter -> Filter -> Sink
    1. 非常适合与数据处理以及数据分析系统
    2. Filter封装数据处理的功能
    3. 松耦合: Filter 只跟数据耦合
    4. Pipe 用于连接Filter传递数据或者在异步处理过程中缓冲数据流进程内同步调用时，
        pipe演变为数据在方法调用间传递
# 组合模式
# Micro Kernel -- 微内核
    Core system + Plug-in Component

    特点
    1. 易于扩展
    2. 错误隔离
    3. 保持架构一致性

    要点
    1. 内核包含公共流程或通用逻辑
    2. 将可变或可扩展部分规划为扩展点
    3. 抽象扩展点行为，定义接口
    4. 利用插件进行扩展

# JSON 解析
    内置JSON解析
        性能比较低
        利用反射实现，通过FeildTag来标识对应的json值
    EasyJson: 采用代码生成而非反射
        go get -u easyjson

# HTTP 服务
    -- 路由规则
        1. URL分为两种: 末尾是/: 表示一个子树，后面可以根其他子路径，末尾不是/表示叶子，固定的路径
        2. 采用最长匹配原则，如果有多个匹配，一定采用匹配路径最长的那个进行处理
        3. 如果没有找到任何匹配项， 会返回404

# 构建RestFul 服务
    -- 更好的路由 httprouter

```

* Go 性能分析
```
# 性能分析工具
    1. graphviz
        brew install graphviz
    2. go-torch
        go get go-torch

    -- 通过文件方式输出Profile
        1. 灵活性高， 适用于特定代码段的分析
        2. 通过手动调用 runtime/pprof 的API
        3. go tool pprof

    Go 支持多种profile
    // 查看CPU profile
    begin-golang/src/ch30/tools/file on  master [!+?] via 🐹 v1.16.6
    ➜ go tool pprof prof cpu.prof

    -- 通过HTTP方式输出Profile
        1. 简单 适合持续性运行的应用
        2. 在应用程序中导入 import _ "net/http/pprof" 并启动http server 即可
        3. http://<host>:<port>/debug/pprof
        4. go tool pprof _http://<host>:<port>/debug/pprof/profile?seconds=10
        5. go-torch -seconds 10 http://<host>:<port>/debug/pprof/profile

# 性能调优过程
    S -> 设定优化目标 -> 分析系统瓶颈点 -> 优化瓶颈点 -> E
    Wall Time: 运行绝对时间
    CPU Time:
    Block Time:
    Memory allocation:
    GC times/time spent:

# sync.RWMutex


# sync.Map
    1. 适合读多写少，且key相对稳定的环境
    2. 采用空间换时间的方案，并且采用指针的方式简介实现值的映射，所以存储空间比较built-in map 大

    分成两块区域:
        Read(R) -- 无锁访问
        Dirty(RW)
# Concurrent Map
    适用读写都很繁琐的情况
    分区域,降低锁冲突的几率

    1. 减少锁的影响范围
    2. 减少发生锁冲突的概率
        sync.Map
        ConcurrentMap
    3. 避免锁的使用

# GC
    1. 避免内存分配和复制
        复杂对象尽量传递引用
        数组的传递
        结构体传递
    -- 打开GC日志
        在程序执行之前加上环境变量 GODEBUG=gctrace=1
        $ GODEBUG=gctrace=1 go test -bench=.
    2. 初始化合适的大小
        自动扩容是有代价
    3. 复用内存

# go tool trace:
    -- 测试程序输出trace 信息
    go test -trace trace.out

    -- 可视化trace信息
    go tool trace trace.out

```

* 分布式系统架构
```
# 面向错误的设计
    > Once you accept that failures will happen, you have the ability to design your system's reaction to the faulures.
    隔离:
        micro kernel
        分布式部署
    重用 vs 隔离
        逻辑结构的重用
        部署结构的隔离
    冗余
        standby Service: 备用
        Online Redundancy:
    单点失效:
    限流:
        token -> leak Bucket
    慢响应:
        A quick rejection is better than a slow response
    不要无休止的等待:
        给阻塞操作加上期限
    错误传递
        断路器:

# 面向恢复的设计
    健康检查:
        注意僵尸进程
            池化资源耗尽
            死锁

    Let is Crash:
        defer func() {
            if err := recover(); err != nil {
                log.Error)"recovered panic", err_
            }
        }()
    构建可恢复的系统:
        1. 拒绝单体系统
        2. 面向错误和恢复的设计
            1.在依赖服务不可用时， 可以继续存活
            2. 快速启动
            3. 无状态
    与客户端协商

# 混沌工程 Chaos Engineering
    If something hurts, do it more often!

    Chaos under control
        Terminate host:
        Inject latency:
        Inject failure:

    Chaos Engineering 原则:
        1.  Build a Hypothesis around Steady State Behavior 定义稳定行为标准
        2. Vary Real-world Events:
        3. Run Experiments in Production
        4. Automate Experiements to Run Continuously
        5. Minimize Blast Radius :
```

书籍推荐
-------
```
Structure and Interpretation of Computer Porgram
RESTful Web Service
```
