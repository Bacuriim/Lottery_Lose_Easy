package main

import (
	"lottery-lose-easy/controllers"

	_ "github.com/lib/pq"
)

func main() {
	controllers.Init()
}
