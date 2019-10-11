# go语言的并发模式
# 1 约束
保证操作安全的方法
- 1.用于共享内存的同步原语 sync.Mutex
- 2.通过通信共享内存来进行同步 channel
- 不会发生改变的数据
- 受保护的数据

约束：确保信息只能从一个并发过程中获取得到。 
约束方法：
- 特定约束 人为约束
- 词法约束 将可能引起并发安全的数据限定在有限作用域，使外部代码不可以操作他。

约束能够通过词法约束，提高了性能、避免了临界区，从而回避了为此的同步成本，从而降低了开发人员的认知负担。

但是，有时建立约束很困难，所以有时我们必须回到我们美妙的go语言的并发原语。
# for-select 循环
for { //无限循环或者range循环
    select{
        case <-time.After(1*time.Second)
        case<-done:
            return
        default:
            //非抢占，当channel还在执行时会进入default
    }
}

# 防止goroutine泄露 负责创建goroutine也要负责他的停止。（后续有context包）
goroutine 不会被运行时垃圾回收
goroutine终止的方式：
- 当他完成工作
- 因为不可恢复的错误停止工作
- 被告知需要终止工作 传入一个用来结束的 done channel，需要结束goroutine时直接close(done)

# or-channel
将多个channel合并成一个channel

# 错误处理

应该将错误视为一等公民，如果goroutine产生错误，name这些错误应该和你想要的结果紧密结合并通过相同的通信线传递。

# pipeline 在系统中形成抽象的另一种工具，特别是需要流处理或批处理时
通常以函数、结构体、方法等形式构造抽象，避免写长函数：
- 部分为了抽象出与大流量无关的细节
- 部分是为了不影响其他区域的情况下处理一个代码区域

pipeline是一系列将数据输入，执行操作并将结果传回的系统。 这些操作是pipeline的一个个stage（阶段）
你可以任意组合pipeline的stage，并单独修改这些stage。

- 一个stage消耗并返回相同类型
- 一个stage必须用语言表达，以便他可以被传递。go语言的函数式一等公民类型满足这一点。
- pipeline和函数式编程密切相关，可以认为是monad的子集
- pipeline始于go语言的channel基元。
- pipeline的生成器是将一组离散值转换为channel上的值流的任何函数

# pipeline一些便利的生成器
# fan-out fan-in 扇出 扇入  
使用条件：它不依赖之前stage计算的值；运行需要很长时间
多个goroutine重用pipeline的单个stage并行化来自上游的stage的pull
## 扇出
启动多个goroutine来处理来自pipeline的输入的过程
## 扇入
将多个结果组合到一个channel的过程

# or-done-channel
用channel中的select包装我们的读操作，并从已完成的channel中进行选择
case v,ok:=<-c:
    if ok == false{
        return
    }
    select{
        case valStream<-v:
        case<-done:
    }

# tee-channel
将一个channel的值同时传递给两个channel

# 桥接channel模式
channel的数据是channel，读取channel中的channel，然后再读取channel中channel中的数据，

# 队列排队
带缓存的channel也是一种队列，不过这里不讨论他
队列通常是优化程序时希望采用的最后一种技术。
预先添加队列会隐藏同步问题（死锁、活锁）
队列不会解决性能问题，只会让程序的行为不同

bufio

# context包
