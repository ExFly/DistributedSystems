# Distributed-Systems

MIT 课程[《Distributed Systems 》](http://nil.csail.mit.edu/6.824/2018/schedule.html)学习和翻译

- 翻译和完成课程的实验代码，之后在代码里添加了注释说明，去掉了代码实现
- 整理课程，编写简单的分布式入门教程

#### 资料推荐

- [《大规模分布式存储系统》](https://book.douban.com/subject/25723658/)
- [《分布式系统原理介绍》](http://pan.baidu.com/s/1geU1XAz)
- [awesome-distributed-systems](https://github.com/kevinxhuang/awesome-distributed-systems)
- [一名分布式存储工程师的技能树是怎样的？](https://www.zhihu.com/question/43687427/answer/96306564)
- [袖珍分布式系统](http://www.jianshu.com/c/0cf64976a481)
- [Students' Guide to Raft](https://thesquareplanet.com/blog/students-guide-to-raft/)
- [分布式系统学习资料](https://gist.github.com/zjhiphop/c4861a6f586e3fdb2379)
- [如何的才能更好的学习MIT6.824分布式系统课程](https://www.zhihu.com/question/29597104)

# final project idea
Here's a list of ideas to get you started thinking -- but you should feel free to propose your own ideas.

* Re-implement any of the systems described in the papers that 6.824 lectured on.
* Build a distributed, decentralized, fault-tolerant Reddit.
* Make the state synchronization protocol (DDP) in Meteor more efficient (e.g., send fewer bytes between server and client) and more fault-tolerant (e.g., a client should be able to tolerate server failures, as long as enough servers remain live).
* Build a system for making Node.js applications fault-tolerant, perhaps using some form of replicated execution.
* Improve the Roost Javascript Zephyr client by replicating the backend to make it fault-tolerant.
* Add cross-shard atomic transactions to Lab 4, using two-phase commit and/or snapshots.
* Build a system with asynchronous replication (like Dynamo or Ficus or Bayou). Perhaps add stronger consistency (as in COPS or Walter or Lynx).
* Build a file synchronizer (like Unison or Tra).
* Build a coherent caching system for use by web sites (a bit like memcached), perhaps along the lines of TxCache.
* Build a distributed cooperative web cache, perhaps along the lines of Firecoral or Maygh.
* Build a collaborative editor like EtherPad, using eventually-consistent or CRDT primitives.
* Build something useful around a blockchain.
* Add one of the following features to "Distributary", an eventually-consistent research data-flow system (with some similarities to Naiad):
    * Transactions with two-phase commit across different base tables.
    * Client-side data-flow operators and stateful caches.
    * Replication for individual data-flow operators or the whole data-flow.
    * Rollback recovery from failures similar to this proposal.