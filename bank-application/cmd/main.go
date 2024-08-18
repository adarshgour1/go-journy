package main

import (
	"github.com/adarshgour1/go-journy/bank-application/pkg/utils"
)

func main() {

	log := utils.NewLogger("bank-management.log")
	log.Print("starting bank management application")
}
