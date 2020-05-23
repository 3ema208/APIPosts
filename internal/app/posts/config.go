package posts

import "github.com/3ema208/pythontask/internal/app/store"

// Config ..
type Config struct {
	BindAddr string `toml:"addr_bind"`
	Store    *store.Config
}
