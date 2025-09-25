# Questionaire


## 1. How to optimize dbms performance:

    1. Indexing Strategies
    2. Query Optimization
    3. Proper Schema Design
    4. Efficient Use of Joins
    5. Caching
    6. Use Efficient Pagination Queries
    7. Partitioning and Sharding

## 2. Identifying Issues in Long SQL Queries

- Use `EXPLAIN` to analyze query execution plan.

`Query`: 
```sql
SELECT e.name, d.department_name
FROM employees e
JOIN departments d ON e.department_id = d.id
WHERE e.salary > 100000;
```

`Using EXPLAIN`:
```sql
EXPLAIN
SELECT e.name, d.department_name
FROM employees e
JOIN departments d ON e.department_id = d.id
WHERE e.salary > 100000;
```


- Look for:

        Full table scans.

        Missing indexes.

        Large temporary tables.

- Profiling tools helps to identify bottlenecks.

## 3. Pagination with LIMIT and OFFSET

```sql
SELECT * FROM users ORDER BY id LIMIT 10 OFFSET 20;
```

`LIMIT` → number of rows.

`OFFSET` → starting row.

- Alternative: keyset pagination for better performance on large datasets.