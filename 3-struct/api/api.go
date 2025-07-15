package api

import "3-struct/config"

type JSONBinAPI struct {
	config *config.Config
}

func NewJSONBinAPI(cfg *config.Config) *JSONBinAPI {
	return &JSONBinAPI{
		config: cfg,
	}
}
