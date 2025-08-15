-- Write your PostgreSQL query statement below
SELECT MAX(SecondHighestSalary) AS SecondHighestSalary
FROM (
    SELECT salary AS SecondHighestSalary
    FROM (
        SELECT salary,
               DENSE_RANK() OVER (ORDER BY salary DESC) AS rnk
        FROM Employee
    ) ranked
    WHERE rnk = 2
) t;
