package main

import "golang-api/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
