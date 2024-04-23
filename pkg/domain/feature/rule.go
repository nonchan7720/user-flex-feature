package feature

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/validator"
)

type Rule struct {
	Name               string              `yaml:"name,omitempty"`
	Query              string              `yaml:"query,omitempty"`
	VariationResult    string              `yaml:"variation,omitempty"`
	Percentages        map[string]float64  `yaml:"percentage,omitempty"`
	ProgressiveRollout *ProgressiveRollout `yaml:"progressiveRollout,omitempty"`
	Disable            *bool               `yaml:"disable,omitempty"`
}

func (r *Rule) Validate() error {
	return validator.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Query, validation.Required),
	)
}
