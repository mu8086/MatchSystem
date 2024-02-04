package main

import (
	"github.com/MatchSystem/matchingsystem"
	"github.com/MatchSystem/router"
)

func main() {
	matchingsystem.InitEngine()
	router.InitRouter()
}
