package main

import (
	"fmt"

	"github.com/kaibling/iggy/service/api"
)

func main() {
	if err := api.Start(); err != nil {
		fmt.Println(err)
	}
}
