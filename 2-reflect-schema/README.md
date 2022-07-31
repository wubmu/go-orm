
## 对象表结构映射
1). 为适配不同的数据库，映射数据类型和特定的 SQL 语句，创建 Dialect 层屏蔽数据库差异

2). 设计 Schema，利用反射(reflect)完成结构体和数据库表结构的映射，包括表名、字段名、字段类型、字段 tag 等。

3). 构造创建(create)、删除(drop)、存在性(table exists) 的 SQL 语句完成数据库表的基本操作。

## Schema
- Field 包含 3 个成员变量，字段名 Name、类型 Type、和约束条件 Tag
- Schema 主要包含被映射的对象 Model、表名 Name 和字段 Fields。
- FieldNames 包含所有的字段名(列名)，fieldMap 记录字段名和 Field 的映射关系，方便之后直接使用，无需遍历 Fields。\

