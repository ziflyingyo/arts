## ARTS(week02)

### 算法题(Algorithm)

[Reverse Integer](https://github.com/ziflyingyo/arts/tree/master/leetcode/7)


### 阅读点评(Review)

#### 文章链接

[ARCHITECTURE](http://docs.ceph.com/docs/master/architecture/)

#### 阅读笔记

1. ceph分布数据过程

step1：数据x的hash值和PG（placement group）数目取余，得到PG编号；
step2：通过CRUSH算法将PG映射到一组OSD中；
step3：根据CRUSH算法把数据x存放到PG对应的OSD中。

2. ceph存储集群可无限伸缩的原因

Ceph 客户端、监视器和 OSD 守护进程可以相互直接交互。

3. 数据条带化

ceph支持条带化（把连续的信息分片存储于多个设备）以增加吞吐量和性能。

4. Ceph 客户端类型

- 块设备： Ceph 块设备（也叫 RBD ）
    提供了大小可调、精炼、支持快照和克隆的块设备。为提供高性能， Ceph 把块设备条带化到整个集群。 Ceph 同时支持内核对象（ KO ） 和 QEMU 管理程序直接使用``librbd`` ——避免了内核对象在虚拟系统上的开销。
- 对象存储： Ceph 对象存储（也叫 RGW ）
    提供了RESTful API ，它与 Amazon S3 和 OpenStack Swift 兼容。
- 文件系统： Ceph 文件系统（ CephFS ）
    提供了兼容 POSIX 的文件系统，可以直接 mount 或挂载为用户空间文件系统（ FUSE ）


### 技术技巧(Tip)

#### MySQL kill线程后出现killed死锁问题

场景现象：
大表（行数过亿），
一个巨大的delete语句 执行一小时后kill ，
show processlist出现killed进程 。


注意：此时不要盲目重启！ 重启MySQL后进程消失但锁依然存在！因为回滚还要继续，这是MySQL对数据的保护机制。

通过下列语句查询锁情况：
```sql
SELECT
 r.trx_id waiting_trx_id,
 r.trx_mysql_thread_id waiting_thread,
 r.trx_query waiting_query,
 b.trx_id blocking_trx_id,
 b.trx_mysql_thread_id blocking_thread,
 b.trx_query blocking_query
FROM
 information_schema.innodb_lock_waits w
INNER JOIN information_schema.innodb_trx b ON b.trx_id = w.blocking_trx_id
INNER JOIN information_schema.innodb_trx r ON r.trx_id = w.requesting_trx_id;
```

结论：
时间过长的update、delete等语句在kill之后会进行回滚操作，会锁表，盲目的杀死正在长时间运行的进程后并不能马上对表进行新的操作，只能等待之前的操作回滚结束，本想用更快的方式操作表结果得不偿失，所以还是建议选择好对表修改操作方式然后一次运行，不再修改。

### 分享(Share)

#### 文章链接
[GIN](https://github.com/gin-gonic/gin)

#### 阅读笔记

func postParams(c *gin.Context){}

看gin的文档，接收从客户端发来的各种参数，有两大类方式：

1. 接收单个参数各种方法

c.Param()
c.Query
c.DefaultQuery
c.PostForm
c.DefaultPostForm
c.QueryMap
c.PostFormMap
c.FormFile
c.MultipartForm

2. 各种绑定方法

c.Bind
c.BindJSON
c.BindXML
c.BindQuery
c.BindYAML
c.ShouldBind
c.ShouldBindJSON
c.ShouldBindXML
c.ShouldBindQuery
c.ShouldBindYAML


常用的获取参数类型demo

1. form：name=lanyangyang
name := c.PostForm("name")

2. JSON body
err = c.BindJSON(&structname)

3. Parameters in path
```go
// This handler will match /user/john but will not match /user/ or /user
router.GET("/user/:name", func(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
})
```
4. Querystring parameters
```go
// Query string parameters are parsed using the existing underlying request object.
// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
router.GET("/welcome", func(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
})
```

#### 思考总结

对于不同的http请求方式，可以参考github上gin的文档找对应获取参数的方法。