CREATE OR REPLACE FUNCTION NthHighestSalary(N INT) RETURNS TABLE (Salary INT) AS $$
BEGIN
  RETURN QUERY (
    -- Write your PostgreSQL query statement below.
    select distinct e1.salary from Employee e1 where N - 1 = (select count(distinct e2.salary) from Employee e2 where e1.salary < e2.salary)
      
  );
END;
$$ LANGUAGE plpgsql;