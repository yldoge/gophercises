# Learn DB Connection in Go

## -- with Postgres

### Put some phone numbers into the database table

like:

```
1234567890
123 456 7891
(123) 456 7892
(123) 456-7893
123-456-7894
123-456-7890
1234567892
(123)456-7892
```

### Normalize them into one form:

to:

```
1234567890
1234567891
1234567892
1234567893
1234567894
```

#### p.s. change your own database connection info in main.go, the original one won't work

```go
const (
	host     = "xxxxdb.cangatcicpl4.us-west-1.rds.amazonaws.com"
	port     = 5432
	user     = "yldog"
	password = "xxxxxxxxxxxx"
	dbname   = "gophercises_phone"
)
```
