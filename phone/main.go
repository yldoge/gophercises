package main

import (
	"bytes"
	"fmt"

	_ "github.com/lib/pq"
	phoneDB "yldoge.com/phone/db"
)

const (
	host     = "xxxxdb.cangatcicpl4.us-west-1.rds.amazonaws.com"
	port     = 5432
	user     = "yldog"
	password = "xxxxxxxxxxxx"
	dbname   = "gophercises_phone"
)

func main() {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		host,
		port,
		user,
		password,
	)
	// connect to default postgres database to do database drop/create
	must(phoneDB.Reset("postgres", connStr, dbname))
	connStr = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)
	// actually connect to gophercises_phone database
	must(phoneDB.Migrate("postgres", connStr))

	db, err := phoneDB.Open("postgres", connStr)
	must(err)
	defer db.Close()

	must(db.Seed())

	phones, err := db.AllPhones()
	must(err)
	for _, p := range phones {
		fmt.Printf("Working on %+v...\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating or removing...", number)
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				// delete this number
				must(db.DeletePhone(p.ID))
			} else {
				// update this number
				p.Number = number
				must(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No actions required")
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func normalize(phone string) string {
	var buf bytes.Buffer
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

// func normalize(phone string) string {
// 	re := regexp.MustCompile("\\D")
// 	return re.ReplaceAllString(phone, "")
// }
