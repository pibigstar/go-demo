# SQL面试题

## 简单

### 组合两个表(175)
> 无论person是否有值，都显示
```sql
SELECT t1.FirstName, t1.LastName, t2.City, t2.State
FROM person t1
	LEFT JOIN Address t2 ON t1.PersonId = t2.PersonId
```

### 第二高的薪水(176)
> Salary 小于 最大值中的最大值就是第二高的
```sql
SELECT MAX(Salary) AS SecondHighestSalary
FROM Employee
WHERE Salary < (
	SELECT MAX(Salary)
	FROM Employee
)
```

### 第N高的薪水(177)
> 根据薪水从高到低排序，取offset
```sql
SELECT (
		SELECT DISTINCT Salary
		FROM Employee
		ORDER BY Salary DESC
		LIMIT N, 1
	) AS NthHighestSalary
```

### 超过经理收入的员工(181)
> 自连表进行查询，t1为员工，t2为经理
```sql
SELECT t1.Name
FROM Employee t1, Employee t2
WHERE t1.ManagerId = t2.Id
	AND t1.Salary > t2.Salary
```

### 查找重复的电子邮箱(182)
```sql
SELECT Email FROM Person GROUP BY Email Having count(*) > 1
```

### 从不订购的客户(183)
> 使用not in
```sql
SELECT t1.Name
FROM Customers t1
WHERE t1.Id NOT IN (
	SELECT CustomerId
	FROM Orders
)
```

### 上升的温度(197)
> 查找昨天的温度比今天高的Id
>
> DATEDIFF(t1.date, t2.date) 两个日期相差天数
```sql
SELECT t1.id
FROM weather t1
	JOIN weather t2
	ON DATEDIFF(t1.date, t2.date) = 1
	AND t1.Temperature > t2.Temperature
```

### 删除重复邮箱(196)
> 自连接
```sql
DELETE P1
FROM Person P1, Person P2
WHERE P1.Email = P2.Email
	AND P1.Id > P2.Id
```
> 查出重复邮箱中ID最大的，然后 使用 in 删除这些最大的Id
```sql
DELETE FROM person WHERE id in 
(SELECT id FROM 
( SELECT DISTINCT t1.id AS id FROM person t1,person t2
WHERE t1.email=t2.email AND t1.id>t2.id) AS temp);
```

### 员工奖金(577)
```sql
SELECT t1.name, t2.bonus
FROM Employee t1
	LEFT JOIN Bonus t2 ON t1.empId = t2.empId
WHERE t2.bonus < 1000
	OR t2.bonus IS NULL
```

### 寻找用户推荐人(584)
```sql
SELECT name FROM customer WHERE referee_id <> 2 OR referee_id IS NULL
```

### 订单最多的客户(586)
```sql
SELECT customer_number
FROM orders
GROUP BY customer_number
ORDER BY COUNT(*) DESC
LIMIT 1
```

### 大的国家(595)
```sql
SELECT name,population,area FROM World WHERE area >= 3000000 OR population >= 25000000;
```

### 超过5名学生的课(596)
```sql
SELECT class FROM courses GROUP BY class Having COUNT(DISTINCT(student)) >= 5;
```
### 总体通过率(597)
```sql
SELECT ROUND( IFNULL((
		SELECT COUNT(*)
		FROM (
			SELECT DISTINCT requester_id, accepter_id
			FROM request_accepted
		) A
	) / (
		SELECT COUNT(*)
		FROM (
			SELECT DISTINCT sender_id, send_to_id
			FROM friend_request
		) B
	),0), 2) AS accept_rate;
```

### 连续空余座位(603)
> 只要挨着我的是空余的，那么就是连续座位
>
> 利用 abs 相减 等于 1 即为相邻
```sql
SELECT DISTINCT t1.seat_id
FROM cinema t1, cinema t2
WHERE (abs(t1.seat_id - t2.seat_id) = 1
	AND t1.free = 1
	AND t2.free = 1)
ORDER BY t1.seat_id
```

### 销售员(607)
> 找出没有向 RED 公司销售过的员工
```sql
SELECT name FROM salesperson WHERE sales_id NOT IN (
    SELECT distinct sales_id FROM orders t1 INNEr JOIN company t2 ON t1.com_id = t2.com_id  AND t2.name = 'RED'
)
```

### 判断三角形(610)
> 如果任意一个两边之和小于等于第三边的则不能构成
```sql
SELECT *,IF((x+y-z<=0 or x+z-y<=0 or y+z-x<=0), 'No','Yes') AS triangle FROM triangle;
```

### 直线上的最近距离(613)
```sql
SELECT MIN(abs(t1.x - t2.x)) AS shortest FROM point t1,point t2 WHERE t1.x <> t2.x;
```

### 只出现过一次的最大数字(619)
> GROUP 找出出现过多次的，再使用 NOT IN 过滤掉这些
```sql
SELECT MAX(num) AS num
FROM my_numbers
WHERE num NOT IN (
	SELECT num
	FROM my_numbers
	GROUP BY num
	HAVING COUNT(*) > 1
)
```

### 有趣的电影(620)
> 找出id为奇数，且不无聊的电影
```sql
SELECT * FROM cinema WHERE description <> 'boring' AND id%2<>0 ORDER BY rating DESC;
```

### 交换工资(627)
> 将 f 和 m 的值互换
```sql
UPDATE salary set sex = IF(sex='f', 'm', 'f');
```
### 合作过至少三次的演员和导演(1050)
```sql
SELECT actor_id,director_id FROM ActorDirector GROUP BY actor_id,director_id Having count(*) >=3;
```
### 产品销售分析I(1068)
```sql
SELECT t2.product_name,t1.year,t1.price FROM Sales t1 Left JOIN Product t2 ON t1.product_id = t2.product_id;
```
### 产品销售分析II(1069)
```sql
SELECT product_id,SUM(quantity) AS total_quantity FROM Sales GROUP BY product_id;
```

### 项目员工I(1075)
> AVG求平均值，ROUND确定精度
```sql
SELECT t1.project_id, ROUND(AVG(t2.experience_years), 2) AS average_years
FROM Project t1
	LEFT JOIN Employee t2 ON t1.employee_id = t2.employee_id
GROUP BY t1.project_id;
```

### 项目员工I(1076)
> 查询项目中员工最多的项目，值得注意的是有可能有多个项目员工并列最多
```sql
SELECT project_id
FROM Project
GROUP BY project_id
HAVING COUNT(*) = (
	SELECT COUNT(*) AS num
	FROM Project
	GROUP BY project_id
	ORDER BY COUNT(*) DESC
	LIMIT 1
)
```
### 产品分析I(1082)
> 查询总销售额最高的销售者，可能有销售额相等者
```sql
SELECT seller_id
FROM Sales
GROUP BY seller_id
HAVING SUM(price) = (
	SELECT SUM(price) AS num
	FROM Sales
	GROUP BY seller_id
	ORDER BY SUM(price) DESC
	LIMIT 1
)
```

### 产品分析II(1083)
> 找出购买了 S8 但没有购买 Iphon的用户
```sql
SELECT t1.buyer_id
FROM Sales t1
	LEFT JOIN Product t2 ON t1.product_id = t2.product_id
GROUP BY t1.buyer_id
HAVING SUM(t2.product_name = 'S8') > 0
AND SUM(t2.product_name = 'iPhone') = 0;
```
> 也可以使用 In 和 NOT In来求解
```sql
SELECT DISTINCT buyer_id
FROM Sales
WHERE buyer_id IN (
		SELECT t1.buyer_id
		FROM Sales t1
			LEFT JOIN Product t2 ON t1.product_id = t2.product_id
		WHERE t2.product_name = 'S8'
	)
	AND buyer_id NOT IN (
		SELECT t3.buyer_id
		FROM Sales t3
			LEFT JOIN Product t4 ON t3.product_id = t4.product_id
		WHERE t4.product_name = 'iPhone'
	)
```

### 产品分析III(1084)
> 找出仅在 '2019-01-01' 到 '2019-03-31' 之间销售的产品
>
> 使用`NOT BETWEEN ... AND ...`
```sql
SELECT DISTINCT t1.product_id, t2.product_name
FROM Sales t1
	LEFT JOIN Product t2 ON t1.product_id = t2.product_id
WHERE t1.sale_date BETWEEN '2019-01-01' AND '2019-03-31'
	AND t1.product_id NOT IN (
		SELECT product_id
		FROM Sales
		WHERE sale_date NOT BETWEEN '2019-01-01' AND '2019-03-31'
	)
```

### 文章浏览(1148)
```sql
SELECT distinct author_id AS id FROM Views WHERE author_id = viewer_id ORDER BY id;
```

### 重新格式化部门表(1179)
```sql
SELECT id, 
SUM(IF(month = 'Jan', revenue, null)) AS Jan_Revenue,
SUM(IF(month = 'Feb', revenue, null)) AS Feb_Revenue,
SUM(IF(month = 'Mar', revenue, null)) AS Mar_Revenue,
SUM(IF(month = 'Apr', revenue, null)) AS Apr_Revenue,
SUM(IF(month = 'May', revenue, null)) AS May_Revenue,
SUM(IF(month = 'Jun', revenue, null)) AS Jun_Revenue,
SUM(IF(month = 'Jul', revenue, null)) AS Jul_Revenue,
SUM(IF(month = 'Aug', revenue, null)) AS Aug_Revenue,
SUM(IF(month = 'Sep', revenue, null)) AS Sep_Revenue,
SUM(IF(month = 'Oct', revenue, null)) AS Oct_Revenue,
SUM(IF(month = 'Nov', revenue, null)) AS Nov_Revenue,
SUM(IF(month = 'Dec', revenue, null)) AS Dec_Revenue
FROM Department GROUP BY id
```

### 每个帖子的评论数(1241)
```sql
SELECT
	post_id,
	COUNT( DISTINCT S2.sub_id ) AS number_of_comments 
FROM
	( SELECT DISTINCT sub_id AS post_id FROM Submissions WHERE parent_id IS NULL ) S1
	LEFT JOIN Submissions S2 ON S1.post_id = S2.parent_id 
GROUP BY S1.post_id
```

## 中等

### 分数排名(178)
> 通过自连接的方式,查找出大于等于它的成绩的数量即为其排名，整理排序即得结果
```sql
SELECT a.Score AS Score, COUNT(DISTINCT b.Score) AS Rank
FROM Scores a, Scores b
WHERE b.Score >= a.Score
GROUP BY a.id
ORDER BY a.Score DESC
```
### 连续出现的数字(180)
> 查找所有至少连续出现三次的数字
```sql
SELECT DISTINCT t1.Num AS ConsecutiveNums
FROM Logs t1, Logs t2, Logs t3
WHERE (t1.Num = t2.Num
	AND t2.Num = t3.Num
	AND t1.Id + 1 = t2.Id
	AND t2.Id + 1 = t3.Id)
```

### 部门中最高的工资(184)
> 通过 `IN` max(Salary)
```sql
SELECT t2.name AS 'Department', t1.name AS 'Employee', t1.Salary
FROM Employee t1
	JOIN Department t2 ON t1.DepartmentId = t2.Id
WHERE (t1.DepartmentId, t1.Salary) IN (
	SELECT DepartmentId, MAX(Salary)
	FROM Employee
	GROUP BY DepartmentId
);
```
### 找出至少有五名下属的经理(570)
```sql
SELECT Name
FROM Employee
WHERE Id IN (
	SELECT ManagerId
	FROM Employee
	GROUP BY ManagerId
	HAVING COUNT(*) >= 5
)
```

### 找出记录最多者(574)
> 使用 ORDER 排序，然后用Limit 1获取最多者
```sql
SELECT Name
FROM Candidate
WHERE id = (
	SELECT CandidateId
	FROM (
		SELECT CandidateId
		FROM Vote
		GROUP BY CandidateId
		ORDER BY COUNT(*) DESC
		LIMIT 1
	) t2
)
```

### 找出回答率最高的问题(578)
> 通过 action和question_id进行分组统计，使用limit 1获取最多的那个
```sql
SELECT question_id AS survey_log
FROM survey_log
GROUP BY action, question_id
HAVING action = 'answer'
ORDER BY COUNT(*) DESC
LIMIT 1
```
### 统计各专业学生人数(580)
```sql
SELECT t1.dept_name, COUNT(t2.student_id) AS student_number
FROM department t1
	LEFT JOIN student t2 ON t1.dept_id = t2.dept_id
GROUP BY t1.dept_id
ORDER BY student_number DESC, t1.dept_name
```

### 2016年的投资(585)
```sql
SELECT
    SUM(insurance.TIV_2016) AS TIV_2016
FROM
    insurance
WHERE
    insurance.TIV_2015 IN
(SELECT TIV_2015 FROM insurance GROUP BY TIV_2015 HAVING COUNT(*) > 1)
AND CONCAT(LAT, LON) IN
(SELECT CONCAT(LAT, LON) FROM insurance  GROUP BY LAT , LON HAVING COUNT(*) = 1)
```

### 好友申请 II ：谁有最多的好友(602)
```sql
SELECT id,num FROM (
SELECT id, COUNT(*) as num FROM (
    SELECT requester_id as id FROM request_accepted UNION ALL 
    SELECT accepter_id FROM request_accepted
) as t1 GROUP BY id ) as t2 ORDER BY num DESC limit 1
```

### 树节点(608)
```sql
select id,
case when t.p_id is null then 'Root' 
     when t.id in (select p_id from tree ) then 'Inner'
     else 'Leaf' 
     end as Type
from tree t 
```

### 平面上的最近距离(612)
```sql
SELECT MIN(ROUND(POW(POW(p1.x-p2.x,2)+POW(p1.y-p2.y,2),1/2),2)) as shortest 
FROM point_2d p1 LEFT JOIN point_2d p2 ON p1.x!=p2.x OR p1.y!=p2.y
```

### 二级关注者(614)
```sql
select f1.follower, count(distinct f2.follower) as num
from (select distinct follower from follow) f1 join follow f2 on f1.follower = f2.followee
group by f1.follower
```

### 换座位(626)
```sql
select (case
when id%2=0 then id-1
when id=(select max(id) from seat) then id
else id+1 end) as id ,student
from seat order by id;
```

### 买下所有产品的客户(1045)
```sql
select customer_id
from 
(select customer_id,count(distinct product_key) as num 
 from Customer
 group by customer_id
) t
join (
    select count(product_key) as num
    from Product
) m 
on t.num = m.num;
```

### 产品销售分析 III(1070)
```sql
select  
    product_id,
    year first_year,
    quantity,
    price
from Sales
where (product_id,year)
in (select product_id,min(year) from Sales group by product_id)
```

### 项目员工 III(1077)
```sql
select p.project_id, p.employee_id
from Project p
join Employee e
on p.employee_id = e.employee_id
where (p.project_id, e.experience_years) in (
select p.project_id,max(e.experience_years)
from project p join employee e on p.employee_id = e.employee_id
group by p.project_id )
```

### 小众书籍(1098)
```sql
select book_id,name
from Books
where book_id not in (select book_id from Orders
where dispatch_date between '2019-06-23'-interval 1 year and '2019-06-23'
group by book_id having sum(quantity)>=10) and available_from<'2019-06-23'-interval 1 month;
```

### 每日新用户统计(1107)
```sql
SELECT login_date, COUNT(user_id) AS user_count
FROM (SELECT user_id, MIN(activity_date) AS login_date
      FROM Traffic
      WHERE activity = 'login'
      GROUP BY user_id) tmp
WHERE DATEDIFF('2019-06-30', login_date) <= 90
GROUP BY login_date
ORDER BY login_date
```

### 每位学生的最高成绩(1112)
```sql
select 
    t.student_id,
    if(count(e.grade) > 1 ,min(e.course_id),course_id) as course_id,
    t.max1 as grade
from Enrollments e 
right join (select student_id,max(grade) as max1  from Enrollments group by student_id )t
    on t.student_id=e.student_id and t.max1 = e.grade
group by e.student_id 
order by t.student_id;
```

### 每月交易 I(1193)
```sql
select
date_format(trans_date,'%Y-%m') as month,
country,
count(*) as trans_count,
sum(if(state='approved',1,0)) as approved_count,
sum(amount) as trans_total_amount,
sum(if(state='approved',amount,0)) as approved_total_amount
from Transactions t
group by
date_format(trans_date,'%Y-%m'),country
```

### 每月交易II(1205)

```sql
select
month,
country,
sum(case when state = 'approved' then 1 else 0 end) as approved_count,
sum(case when state = 'approved' then amount else 0 end) as approved_amount,
sum(case when state = 'charged' then 1 else 0 end) as chargeback_count,
sum(case when state = 'charged' then amount else 0 end) as chargeback_amount
from
(
select c.trans_id as id, t.country,'charged' as state,amount,c.trans_date,date_format(c.trans_date,'%Y-%m') as month
from chargebacks c
left join transactions t on c.trans_id = t.id
union all
select *,date_format(trans_date,'%Y-%m') as month
from transactions
where state = 'approved' #去零
) as temp
group by month,country
order by month
```

## 困难

### 游戏玩法分析
1. 玩家第一天登录游戏日期(511)
```sql
SELECT player_id, MIN(event_date)  AS first_login FROM Activity GROUP BY player_id
```
2. 玩家第一天登录游戏设备名称(512)
```sql
SELECT player_id, device_id
FROM Activity
WHERE (player_id, event_date) IN (
	SELECT player_id, MIN(event_date)
	FROM Activity
	GROUP BY player_id
)
```
3. 统计各个日期前玩家的登录次数(534)
```sql
SELECT t2.player_id, t2.event_date, SUM(t1.games_played) AS games_played_so_far
FROM Activity t1
	JOIN Activity t2
	ON t1.event_date <= t2.event_date
		AND t1.player_id = t2.player_id
GROUP BY t2.player_id, t2.event_date
```
4. 求首次登陆之后，第二天也登陆的玩家占总玩家的比率(550)
```sql
SELECT ROUND(SUM(IF(DATEDIFF(t1.event_date, t2.first_date) = 1, 1, 0)) / 
        COUNT(DISTINCT t1.player_id), 2) AS fraction
FROM Activity t1, (
		SELECT player_id, MIN(event_date) AS first_date
		FROM activity
		GROUP BY player_id
	) t2
WHERE t1.player_id = t2.player_id
```

### 求各个公司薪水中位数(569)
```sql
SELECT e1.Id, e1.Company, e1.Salary
FROM Employee e1, Employee e2
WHERE e1.Company = e2.Company
GROUP BY e1.Company, e1.Salary
HAVING SUM(CASE 
	WHEN e1.Salary >= e2.Salary THEN 1
	ELSE 0
END) >= COUNT(*) / 2
AND SUM(CASE 
	WHEN e1.Salary <= e2.Salary THEN 1
	ELSE 0
END) >= COUNT(*) / 2
ORDER BY e1.Company;
```

### 行程和用户取消率(262)
> ROUND(x, d)  x: 要处理的数 d: 保留几位小数
>
> IF(表达式, x, y): 表达式为true时取 x, false时取 y, 类似于Java的三元表达式
```sql
SELECT t1.Request_at, 
        ROUND(SUM(IF(t1.Status = 'cancelled', 1, 0)) / COUNT(t1.Status), 2)
FROM Trips t1
	JOIN Users t2
	ON t1.Client_Id = t2.Users_Id
		AND t2.Banned = 'No'
	JOIN Users t3
	ON t1.Driver_Id = t3.Users_Id
		AND t3.Banned = 'No'
WHERE t1.Request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY t1.Request_at
```

### 给定数字的频率查询中位数(571)

```sql
select
avg(t.number) as median
from
(
select
n1.number,
n1.frequency,
(select sum(frequency) from numbers n2 where n2.number<=n1.number) as asc_frequency,
(select sum(frequency) from numbers n3 where n3.number>=n1.number) as desc_frequency
from numbers n1
) t
where t.asc_frequency>= (select sum(frequency) from numbers)/2
and t.desc_frequency>= (select sum(frequency) from numbers)/2
```

### 查询员工的累计薪水(579)
```sql
SELECT
    E1.id,
    E1.month,
    (IFNULL(E1.salary, 0) + IFNULL(E2.salary, 0) + IFNULL(E3.salary, 0)) AS Salary
FROM
    (SELECT
        id, MAX(month) AS month
    FROM
        Employee
    GROUP BY id
    HAVING COUNT(*) > 1) AS maxmonth
        LEFT JOIN
    Employee E1 ON (maxmonth.id = E1.id
        AND maxmonth.month > E1.month)
        LEFT JOIN
    Employee E2 ON (E2.id = E1.id
        AND E2.month = E1.month - 1)
        LEFT JOIN
    Employee E3 ON (E3.id = E1.id
        AND E3.month = E1.month - 2)
ORDER BY id ASC , month DESC
```

### 体育馆的人流量(601) 
```sql
SELECT distinct a.*
FROM stadium as a,stadium as b,stadium as c
where ((a.id = b.id-1 and b.id+1 = c.id) or
       (a.id-1 = b.id and a.id+1 = c.id) or
       (a.id-1 = c.id and c.id-1 = b.id))
  and (a.people>=100 and b.people>=100 and c.people>=100)
order by a.id;
```

### 平均工资：部门与公司比较(615)
```sql
select department_salary.pay_month, department_id,
case
  when department_avg>company_avg then 'higher'
  when department_avg<company_avg then 'lower'
  else 'same'
end as comparison
from
(
  select department_id, avg(amount) as department_avg, date_format(pay_date, '%Y-%m') as pay_month
  from salary join employee on salary.employee_id = employee.employee_id
  group by department_id, pay_month
) as department_salary
join
(
  select avg(amount) as company_avg,  date_format(pay_date, '%Y-%m') as pay_month from salary group by date_format(pay_date, '%Y-%m')
) as company_salary
on department_salary.pay_month = company_salary.pay_month
```
### 游戏玩法分析 V(1097)
```sql
SELECT t0.install_dt, installs, ROUND(IFNULL(retention, 0)/installs, 2) Day1_retention FROM
(SELECT install_dt, COUNT(*) installs FROM 
(SELECT player_id, MIN(event_date) install_dt FROM Activity GROUP BY player_id) t1
GROUP BY install_dt) t0
LEFT JOIN
(SELECT install_dt, COUNT(*) retention FROM
(SELECT player_id, DATE_SUB(event_date, INTERVAL 1 DAY) install_dt FROM Activity WHERE (player_id, DATE_SUB(event_date, INTERVAL 1 DAY)) IN 
(SELECT player_id, MIN(event_date) install_dt FROM Activity GROUP BY player_id)
 ) t3 GROUP BY install_dt) t4 ON t0.install_dt = t4.install_dt 
```

### 锦标赛优胜者(1194)
```sql
select group_id,player_id from 
(select group_id,player_id,sum((
    case when player_id = first_player then first_score
         when player_id = second_player then second_score
         end
)) as totalScores
from Players p,Matches m
where p.player_id = m.first_player
or p.player_id = m.second_player
group by group_id,player_id
order by group_id,totalScores desc,player_id) as temp
group by group_id
order by group_id,totalScores desc,player_id
```

### 报告系统状态的连续日期(1225)
```sql
select status as period_state, min(date) as start_date, max(date) as end_date
from
(select * from (select fail_date as date, 'failed' as status, if(datediff(@pre, @pre := fail_date) = -1, @a, @a := @a + 1) as groupmark from Failed, (select @a := 0, @pre := "2018-07-07") as temp
union
select success_date as date, 'succeeded' as status, if(datediff(@cur, @cur := success_date) = -1, @b, @b := @b + 1) as groupmark from Succeeded, (select @b := 100, @cur := "2018-07-07") as temp
) as temp where date between "2019-01-01" and "2019-12-31"
order by date) as s
group by groupmark
order by min(date);
```

