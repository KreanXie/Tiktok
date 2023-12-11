package main

import (
	db "tiktok/middleware/database"
	"tiktok/middleware/router"
)

func main() {
	db.InitDB()
	// rd.InitRedis()
	// kk.InitKafka()
	router.InitRouter()
}
