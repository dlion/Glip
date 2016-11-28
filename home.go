package main

import (
	"math/rand"

	"github.com/kataras/iris"
)

func home(ctx *iris.Context) {
	const list = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomRoute := make([]byte, 6)
	db := openDB(DBpath)
	defer db.Close()

	for {
		for i := range randomRoute {
			randomRoute[i] = list[rand.Intn(len(list))]
		}
		_, err := checkURL(db, string(randomRoute))
		if err != nil { //If the url is not taken stop the cycle
			break
		}
	}

	ctx.Render("home.html", map[string]interface{}{"URLrandom": string(randomRoute)})
}