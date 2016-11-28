package main

import (
	"math/rand"
	"time"

	"github.com/kataras/go-template/django"
	"github.com/kataras/iris"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	//Conf
	iris.UseTemplate(django.New()).Directory("./templates", ".html")
	iris.Config.Gzip = true

	//Middleware to init the db
	iris.UseFunc(initDB)

	/**
	 * Routes
	 */

	//Home
	iris.Get("/", home)

	//getURL
	iris.Get("/:url", getUrl)
	//iris.POST("/:url", postUrl)

	//Listen
	iris.Listen(":8080")
}

func getUrl(ctx *iris.Context) {
	url := ctx.Param("url")
	ctx.Render("url.html", map[string]interface{}{"url": string(url), "taken": false, "msg": false})
}
