package config

import (
	"xd/lib/configparser"
)

type RPCConfig struct {
	Enabled bool
	Bind    string
	// TODO: authentication
}

func (cfg *RPCConfig) FromSection(s *configparser.Section) {
	if s == nil {
		return
	}
	cfg.Bind = s.Get("bind", "127.0.0.1:1188")
	cfg.Enabled = s.Get("enabled", "1") == "1"
}

func (cfg *RPCConfig) Options() map[string]string {
	enabled := "1"
	if !cfg.Enabled {
		enabled = "0"
	}
	opts := map[string]string{
		"enabled": enabled,
	}
	if cfg.Bind != "" {
		opts["bind"] = cfg.Bind
	}
	return opts
}
