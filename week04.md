## ARTS(week04)

### 算法题(Algorithm)

[AddTwoNumbers](https://github.com/ziflyingyo/arts/tree/master/leetcode/2)


### 阅读点评(Review)

#### 文章链接

[10 Tips for Monitoring Best Practices (Alerting and Notifications)](https://www.monitis.com/blog/10-tips-for-monitoring-best-practices-alerting-and-notifications/)

#### 阅读笔记

1. 计划配置对你有用的告警
分层次告警，比如：
阈值1：告警对象为新员工和凌晨还在值班的员工。
阈值2：告警对象为leader，这个时候可能新员工和值班员工已经忙不过来了。
阈值3：告警对象为产品负责人或直接负责人。

2. 根据重要性设置告警级别
分级别告警，比如：
级别1：不处理好，会导致丢失工作或者客户。
级别2：不处理好，领导会很紧张。
级别3：犹如来了个电话，还是要及时处理。
级别4：没人注意到的事情，可以延缓处理。

3. 不允许单点故障

4. 了解你的受众

5. 测试你的监控工具和告警系统

6. 不要设置告警邮件过滤，如果有必要，参照1

7. 如果系统太安静了，很可能有异常。此时要多检查各项服务

8. 提供告警的解决处理过程

9. 寻求帮助。找技服，找专家，找客户沟通。

10. 将所有事情文档化。
实时按需更新相关文档。将变更同步到团队每个成员。



#### 思考总结

1. 配置告警策略时，应以实用性为标准进行决策。
	
2. 监控告警服务本身也需要监控。

3. 告警的同时提供告警的处理参考办法有利于提高处理效率。

### 技术技巧(Tip)

#### MySQL排序分组问题

针对一张表，按字段1分组，分组内部取字段2最大的那一行。

使用MySQL5.7前，可以使用order by子句完成，如：

```sql
select * from (select * from tablename order by column2 desc)  as a group by column1;
```

MySQL5.7后，MySQL团队为了提高数据库sql执行性能，去掉了group by的隐式排序功能，并做了相关调整，因此上述方法在MySQL5.7中不适合使用。

新的办法：
按column1分组取到column2的最大值，构成一个子句；用原始表inner join子句后再group bycolumn1。SQL如下：

```sql
select a.* from tablename a inner join (select column1,max(column2) as max_colum2 from tablename group by column1) as b on a.column1=b.column1 and a.column2=b.max_column2 group by a.column1;
```

#### 参考链接
[MySql5.7 分组查询排序有Bug](https://bbs.aliyun.com/read/312722.html?spm=a2c4e.11155515.0.0.2ae030b1vRTOhy)
[MySQL order by before group by](https://stackoverflow.com/questions/14770671/mysql-order-by-before-group-by)

### 分享(Share)

#### 文章链接
[Back to Basics: Current Monitoring and Alerting Best Practices](https://victorops.com/blog/back-to-basics-current-monitoring-and-alerting-best-practices)

#### 阅读笔记

理解软件有很多种形式，有时候最好的起点不是开始，而是结束。为了做好监控告警工具，首先需要确定您最希望了解关于系统的哪些内容。比如：

1. 识别当前和潜在的问题
2. 收集可操作的指标，包括内部的和外部的（比如客户影响因子）
3. 可操作的告警

#### 思考总结

通过学习此文章，引发了如下思考：
1. 潜在问题和现有问题是否有必要tag区分；
2. 如何提供自动补救工具，以减少人为干预频次；
3. 智能运维如何做好容量规划；
4. 全局视图如何更好的提供更全面的信息。


