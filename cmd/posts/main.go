package main

import (
	"github.com/3ema208/pythontask/internal/app/posts"
	"github.com/BurntSushi/toml"
)

func main() {
	config := &posts.Config{}
	toml.DecodeFile("config/posts.toml", config)
	apiposts := posts.New(config)
	if err := apiposts.Start(); err != nil {
		panic(err)
	}
}
