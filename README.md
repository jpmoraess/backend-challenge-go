**Database Migration**

**See:** github.com/golang-migrate/migrate/tree/master/cmd/migrate

go get github.com/golang-migrate/migrate/v4

**Create Migration**

**Run:** migrate create -ext- sql -dir db/migration -seq init_schema

---
**SQLC**

**Run:** sqlc generate