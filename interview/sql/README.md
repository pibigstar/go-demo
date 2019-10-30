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




