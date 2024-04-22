package retriever

import (
	"github.com/goccy/go-yaml"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
)

func ConvertToFlagStruct(buf []byte) (feature.Flags, error) {
	var flags feature.Flags
	if err := yaml.Unmarshal(buf, &flags); err != nil {
		return nil, err
	}
	return flags, nil
}
