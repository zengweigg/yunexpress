package yunexpress

import (
	"github.com/go-resty/resty/v2"
	"github.com/zengweigg/yunexpress/config"
)

type service struct {
	config     *config.Config // Config
	logger     Logger         // Logger
	httpClient *resty.Client  // HTTP client
}

type services struct {
	Base baseService
	// Gts  gtsService
	// Icms icmsService
	// Icsm icsmService
}
