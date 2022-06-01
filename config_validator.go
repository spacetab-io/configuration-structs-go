package cfgstructs

import (
	"github.com/asaskevich/govalidator"
	"github.com/spacetab-io/configuration-structs-go/v2/contracts"
)

type ConfigValidator struct {
	errs Errors
	cfgs []contracts.ValidatableInterface
}

func NewConfigValidator() contracts.ConfigValidatorInterface {
	return &ConfigValidator{
		errs: make(Errors, 0),
		cfgs: make([]contracts.ValidatableInterface, 0),
	}
}

func (cv *ConfigValidator) Add(cfg contracts.ValidatableInterface) {
	cv.cfgs = append(cv.cfgs, cfg)
}

func (cv ConfigValidator) GetConfigs() []contracts.ValidatableInterface {
	return cv.cfgs
}

func (cv *ConfigValidator) Validate() error {
	for _, config := range cv.GetConfigs() {
		result, errs := config.Validate()
		if !result && errs != nil {
			//nolint: errorlint
			switch v := errs.(type) {
			case govalidator.Errors:
				cv.errs = append(cv.errs, v...)
			case error:
				cv.errs = append(cv.errs, v)
			}
		}
	}

	if len(cv.errs) == 0 {
		return nil
	}

	return cv.errs
}
