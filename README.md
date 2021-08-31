# Trouble about Ent with MySQL 5.7

## Enviroment

Go
```sh
go version
go version go1.17 darwin/amd64
```

- Ent: v0.9.1
- MySQL: 5.7

## Schema

```go
package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknown"),
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
```

## Reproduction Method

```sh
git clone git@github.com:iwata/ent-issue.git
cd ent-issue
docker compose up -d
go mod download
go run cmd/main.go
```

`cmd/main.go` executes a simple task, so migrates schema, creates a `user`, and selects a `user`.
But it occures an error.

## Current Behavior

```sh
go run cmd/main.go
2021/08/31 20:08:49 failed to get a user: sql: Scan error on column index 2, name "created_at": unsupported Scan, storing driver.Value type []uint8 into type *time.Time
exit status 1
```

And it creates `users`, but not expected about `created_at`.
```sh
mysql -h 127.0.0.1 -P 3306 -u root -ppass test -e 'desc users'
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | bigint(20)   | NO   | PRI | NULL    | auto_increment |
| name       | varchar(255) | NO   |     | unknown |                |
| created_at | timestamp    | YES  |     | NULL    |                |
+------------+--------------+------+-----+---------+----------------+
```

## Expected Behavior

```sh
go run cmd/main.go
select a user successfully

```

And it should create `users` looks like below.
```sh
mysql -h 127.0.0.1 -P 3306 -u root -ppass test -e 'desc users'
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | bigint(20)   | NO   | PRI | NULL    | auto_increment |
| name       | varchar(255) | NO   |     | unknown |                |
| created_at | timestamp    | NO   |     | NULL    |                |
+------------+--------------+------+-----+---------+----------------+
```
`created_at` shouldn't be nullable.