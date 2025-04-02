package service

import (
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

type UIConfigService struct {
	cfg config.Configuration
}

func NewUIConfigService(cfg config.Configuration) *UIConfigService {
	return &UIConfigService{cfg}
}

func (cs *UIConfigService) GenerateUIConfigs() entity.UIConfig {
	return entity.UIConfig{
		ImportLocalPath:  cs.cfg.App.ImportLocalPath,
		ExmportLocalPath: cs.cfg.App.ExportLocalPath,
	}
}
