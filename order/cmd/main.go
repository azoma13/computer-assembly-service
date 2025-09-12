package main

import (
	"github.com/azoma13/computer-assembly-service/order/internal/app"
)

const configPath = "order/config/config.yaml"

func main() {
	app.Run(configPath)
}
