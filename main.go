package main

import (
	db "tiktok/middleware/database"
	rd "tiktok/middleware/redis"
	"tiktok/middleware/router"
)

func main() {
	db.InitDB()
	rd.InitRedis()
	router.InitRouter()
}
