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

	iris.Get("/url/:url", getUrl)
	iris.Post("/url/:url", postUrl)

	iris.Get("/result/:url", func(ctx *iris.Context) {
		ctx.Render("result.html", map[string]interface{}{"url": ctx.Param("url"), "hostname": ctx.HostString})
	})

	//Listen
	iris.Listen(":8080")
}
