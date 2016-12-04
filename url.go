package main

import (
	"strconv"

	"github.com/kataras/iris"
)

func postUrl(ctx *iris.Context) {
	url := ctx.Param("url")
	db := openDB(DBpath)
	defer db.Close()
	if _, err := checkURL(db, string(url)); err != nil {
		msg := ctx.FormValueString("msg")
		ip := ctx.RemoteAddr()
		if len(msg) > 0 {
			addUrl(db, url, msg, ip)
			ctx.Redirect("/result/" + string(url))
		} else {
			ctx.Redirect("/url/" + string(url))
		}
	} else {
		ctx.Redirect("/")

	}
}

func getUrl(ctx *iris.Context) {
	var taken bool
	var msg string
	url := ctx.Param("url")
	db := openDB(DBpath)
	defer db.Close()
	if infos, err := checkURL(db, string(url)); err == nil {
		taken = true
		msg = string(infos[2])
		id, _ := strconv.Atoi(infos[0])
		deleteUrl(db, id)
	}
	ctx.Render("url.html", map[string]interface{}{"url": string(url), "taken": taken, "msg": msg, "hostname": ctx.HostString})
}
