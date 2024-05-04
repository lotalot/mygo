package sql

/*事务
事务是一组原子操作，要么全部成功，要么全部失败。
start transaction;   关闭sql默认的自动提交
commit;  提交
rollback;  回滚
事务的四个特性：
原子性，一致性，持久性，隔离性
事务之间的四个隔离级别：
读未提交              最低   a可以读到b未提交的数据（脏读）
读已提交              	     a可以读到已提交（不可重复读取数据）
可重复读  mysql默认          a每一次读取的数据都是一致的（幻影读，数据不够真实）
串行化(Serializable)  最高   事务排队，解决所有问题，效率最低
*/

/*索引 idIndex
通过B树实现，缩小扫描范围，避免全表扫描
主键和unique自动添加索引，索引需要维护
如果数据量大，经常作为条件，字段很少增删改 来使用索引
索引失效：使用模糊查询，复合索引（两个字段合成）没有使用左侧的列查找，where中索引列参加运算
索引列使用了函数，使用or两边都需要有索引，
*/

/*视图 view
create view a as select * from biao;   创建视图
视图相当于引用，将复杂的sql语句封装起来，方便调用
修改视图对象时本体也会被修改
*/

/*数据导入导出
source D:\data.sql  导入数据
mysqldump ku biao>D:\x.sql 导出数据
*/

/*数据库设计三范式
第一范式 1NF:属性值的原子性，不可再分（既要么全成功要么全失败）
第二范式 2NF:非主键完全依赖于主键
第三范式 3NF:任何非主键都直接依赖于主键，不产生传递依赖
一对多两张表，多的表加外键
多对多三张表，关系表两个外键
*/