package feature

import "time"

type ProgressiveRollout struct {
	Initial *ProgressiveRolloutStep `yaml:"initial,omitempty"`
	End     *ProgressiveRolloutStep `yaml:"end,omitempty"`
}

type ProgressiveRolloutStep struct {
	Variation  *string    `yaml:"variation,omitempty"`
	Percentage *float64   `yaml:"percentage,omitempty"`
	Date       *time.Time `yaml:"date,omitempty"`
}
