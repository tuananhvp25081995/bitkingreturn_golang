package main

import (
	"github.com/tuananhvp25081995/bitkingreturn_golang/db"
	"github.com/tuananhvp25081995/bitkingreturn_golang/routers"
)

func main() {

	db, err := db.New("mongodb://localhost:27017", "bitkingreturns")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := routers.Create(db)

	r.Logger.Fatal(r.Start("127.0.0.1:3000"))
}
