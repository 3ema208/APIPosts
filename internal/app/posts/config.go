package posts

import "github.com/3ema208/APIPosts/internal/app/store"

// Config ..
type Config struct {
	BindAddr string `toml:"addr_bind"`
	Store    *store.Config
}
