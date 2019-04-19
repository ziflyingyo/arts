## ARTS

### 算法题(Algorithm)

[Two Sum](https://github.com/ziflyingyo/arts/tree/master/leetcode/1)


### 阅读点评(Review)

#### 文章链接
[CAP theorem](https://en.wikipedia.org/wiki/CAP_theorem)


#### 阅读笔记

CAP 三个性质都是一个度的问题。

当代 CAP 实践应将目标定为针对具体的应用，在合理范围内最大化数据一致性和可用性的“合力”。脱离具体的应用，讨论CAP意义并不是很大。

数据库集群：用至少两台或者多台数据库服务器，构成一个虚拟单一数据库逻辑映像，表面上看跟一个数据库是的，而实际上多台数据库实例。


#### 思考总结

面对一个具体的应用，可以先围绕CAP三方面提一些问题：
C: 数据不一致的后果会有哪些，有哪些方案可以弥补数据一致性？
A：停服5min的代价是什么，停服1h的代价是什么？
P：数据量预估能有多大，如果很大的话，数据能否允许定期归档后移除，单个节点是否能满足业务需求。 目前是否有必要考虑分区容错性问题？

然后根据问题的答案在CAP方面做取舍选择。

一般来说，CAP不是C，A，P三选二的问题，而是在P情况，如何选择A或C。

hbase是什么类型的服务？
CP系统。
Hbase具有以下特点：

每个值只出现在一个REGION
同一时间一个Region只分配给一个Region服务器
行内的mutation操作都是原子的(原子性操作是指：如果把一个事务可看作是一个程序,它要么完整的被执行,要么完全不执行)。
put操作要么成功，要么完全失败。

参考资料：https://www.cnblogs.com/captainlucky/p/4720986.html

Cassandra是什么类型的服务？
AP系统。
写操作一致性决定了在向客户端确认写操作成功之前，多少个节点必须被成功写入（Commitlog and Memtable）。

假设：R=Nodes Read, W=Node Written, N=Replication Factor，Q=QUORUM=N/2+1。

ANY：                   至少成功写入一个节点，即使是一个Hinted Handoff
ONE：                   至少成功写入一个复制节点（replica node）
QUORUM：            至少成功写入Q个复制节点（Q=N/2+1）
LOCAL_QUORUM： 至少在coordinator node所在的当前DC成功写入Q个复制节点
EACH_QUORUM：  在每个DC成功写入Q个复制节点
ALL：                    成功写入集群中的每个复制节点


Cassandra允许用户来决定每个请求的CAP特性，在一致性、性能和错误容忍度之间做出选择。

### 技术技巧(Tip)

#### redis内存碎片问题

在Redis4.0以前，回收redis内存碎片需要重启redis服务器。

Redis4.0以后，使用jemalloc作为内存分配器。可以使用新增指令手动回收内存碎片。也可以配置自动内存碎片清理。

手动回收命令：memory purge

参考链接：https://my.oschina.net/watliu/blog/1620666



### 分享(Share)

#### 文章链接

[Fallacies of distributed computing](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing)

#### 阅读笔记
1. 网络是可靠的。
2. 延迟是零。
3. 带宽是无限的。
4. 网络是安全的。
5. 拓扑结构不会改变。
6. 只有一个管理员。
7. 运输成本为零。
8. 网络是同构的。

#### 思考总结

在进行故障定位时，可以思考下分布式计算的八大谬论，没准可以找到定位问题的新思路。


