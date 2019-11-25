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




