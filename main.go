package main

import "cafex/configs"

func main() {
	r := configs.SetupRoute()
	r.Run(":8000")
}
