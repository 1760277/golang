package main

import "app/configs"

func main() {
	r := configs.SetupRoute()
	r.Run(":8000")
}
