package main

/*登入mysql:mysql -uroot -p+密码
show databases;          显示可用的数据库，都使用；结尾
create database biao;    创建一个表
select * from biao;      查询表的数据
desc +表名               查询表的结构
use +库名                使用数据库
in 等于多个or()
like模糊查询  '%+字符串+%'   查询包含某个字符串的数据 _ 表示单个字符 % 表示多个
添加/表示转意
explain +语句 查询类型

select 列名 as '别名'
from 表名 where 条件
group by 分组列名和分组函数 + having
order by 按照什么列排序 asc升序 desc（降序）, 若一样再按照此排序
limit 随机查询行数
空集null加任何数都为null需要使用ifnull()函数处理

执行优先级：在where中使用别名会报错因为where先执行
from >> (xxx join) >> where >> group by >> 聚合 >>having >> select>> order by >> limit
*/

/*单行处理函数：
1. lower(a) 转换小写
2. upper(a) 转换大写
3. substr 取子串（substr( 被截取的字符串, 起始下标,截取的长度)）
4. concat(a,b,c,d,e,f....) 函数进行字符串的拼接
5. length(a) 求长度
6. trim 去空格
7. round 四舍五入
9. ifnull(数据，该数据为null时被当做哪个值) 可以将 null 转换成一个具体值
10.case..when..then..when..then..else..end
case匹配一个列名 满足when条件时执行then效果
11.str_to_date 将字符串转换成日期
12.date_format 格式化日期
13.format 设置千分位
14.distinct 去除重复记录（后面加多个时，联合一起后去除重复的），在select后直接添加
*/

/*多行处理函数（分组函数）:自动忽略null
group_concat(a) 字符串拼接函数，将多个值连接成一个字符串
group by 分组
count(列名) 统计某列的行数
sum(列名) 求和
avg(列名) 平均值
max(列名) 最大值
min(列名) 最小值

分组查询group by，先进行分组再进行查询，可以使用where对数据进行先过滤，再使用分组
在group by后可以添加having对数据进行再过滤，在无法使用where时使用having
*/

/*连接查询：join(xxx,yyy) 左连接、右连接、内连接、全连接
select t1.name ,t2.blood from myku.t1 , myku.t2 where t1.id = t2.id;
连接时会将每行互相组合给表起别名
在后续版本：把  ‘ ，’改为join  where改为on
select name ,blood from myku.tone join myku.two on m.t.id = m.t.id ;

内连接，能够完全匹配条件，两表没有主次关系；外连接（左右）
a right join b b为主表，a为从表；  left join b b为主表，a为从表

select...
from a
join b on...          a和b连，条件
right join c on...    a和c连，条件
*/

/*子查询 在select中嵌套select语句

select id,name,sleeptime from myku.t1     where中的子查询
where sleeptime > (select min(sleeptime) from myku.t1 );

select t.*, t3.call        from中的子查询，相当于在from中新建了一张表
from  (select t1.sex , avg(sleeptime) as ta from myku.t1 group by sex) t
join myku.t3 on t.ta between t3.dt and t3.ut; 通过起别名ta防止被视为函数

select中的子查询，只能访问一条
*/

/*
union合并查询结果集,减少匹配的数量级做加法
select name, xp from myku.t1 where xp = '女'
union
select name, xp from myku.t1 where xp = '男';

limit 起始下标(从0开始) , 长度 :用于分页
约束：非空约束,唯一性约束,主键约束,外键约束,检查约束
source +sql文件地址执行sql文件

外键约束FK,为字段添加约束条件,条件为父表(有唯一性，可以空),被约束方为子表,子表只引用父表中的数据

存储引擎：存储数据的方式（mysql中）
InnoDB引擎默认，支持事务处理（安全），MyISAM不支持事务处理（可以压缩）。
*/
