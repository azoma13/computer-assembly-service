package main

import "github.com/azoma13/computer-assembly-service/order/internal/app"

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}
