package feature

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/validator"
)

type Rule struct {
	Name               string              `yaml:"name"`
	Query              string              `yaml:"query"`
	VariationResult    string              `yaml:"variation"`
	Percentages        map[string]float64  `yaml:"percentage,omitempty"`
	ProgressiveRollout *ProgressiveRollout `yaml:"progressiveRollout,omitempty"`
	Disable            *bool               `yaml:"disable,omitempty"`
}

func (r *Rule) Validate(variations []string) error {
	return validator.ValidateStruct(r,
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Query, validation.Required),
		validation.Field(&r.VariationResult, validation.Required, validation.In(variations)),
	)
}
