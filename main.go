package main

import (
	"github.com/brunoaalexandree/gook/config"
	"github.com/brunoaalexandree/gook/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.ErrF("Config initialization error: %v", err)
		return
	}

	router.Initialize()
}
