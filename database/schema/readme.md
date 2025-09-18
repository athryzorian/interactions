All "alter" statements need to be maintained in alter.sql separately. Please do not keep them in create.sql.
Execute DDL statements in following orders.
1. All create.sql statements
2. All alter.sql statements
3. Import data from add ../data/<table name> tables. (Exclude id column while importing)