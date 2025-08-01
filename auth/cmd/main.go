package main

import (
	"fmt"
	"github.com/rnymphaea/chronoflow/auth/internal/config"
)

func main() {
	cfg, err := config.LoadServerConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cfg)
	}
}
