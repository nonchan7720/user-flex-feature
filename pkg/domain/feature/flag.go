package feature

import "strings"

type Flag struct {
	Variations      *map[string]*interface{} `yaml:"variations,omitempty"`
	Rules           []*Rule                  `yaml:"targeting,omitempty"`
	DefaultRule     *Rule                    `yaml:"defaultRule,omitempty"`
	Experimentation interface{}              `yaml:"experimentation,omitempty"`
	Scheduled       interface{}              `yaml:"scheduledRollout,omitempty"`
	TrackEvents     *bool                    `yaml:"trackEvents,omitempty"`
	Disable         *bool                    `yaml:"disable,omitempty"`
	Version         *string                  `yaml:"version,omitempty"`
	Metadata        *map[string]interface{}  `yaml:"metadata,omitempty"`
}

func (f *Flag) FindRule(name string) *Rule {
	_, rule := f.findRule(name)
	return rule
}

func (f *Flag) findRule(name string) (int, *Rule) {
	for idx, rule := range f.Rules {
		if strings.EqualFold(rule.Name, name) {
			return idx, rule
		}
	}
	return -1, nil
}

func (f *Flag) AppendOrUpdateRule(rule *Rule) {
	idx, rule := f.findRule(rule.Name)
	if rule != nil {
		f.Rules[idx] = rule
	} else {
		f.Rules = append(f.Rules, rule)
	}
}
