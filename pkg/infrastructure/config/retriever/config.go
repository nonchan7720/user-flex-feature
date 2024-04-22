package retriever

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	. "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/inmemory"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/validator"
)

type RetrieverType string

const (
	FileType     RetrieverType = "file"
	InMemoryType RetrieverType = "inMemory"
)

type Retriever struct {
	Type     RetrieverType `yaml:"type"`
	File     *File         `yaml:"file"`
	InMemory *InMemory     `yaml:"inMemory"`
}

func (c *Retriever) when(typ RetrieverType, rules ...validation.Rule) validation.WhenRule {
	return validation.When(c.Type == typ, append([]validation.Rule{validation.Required}, rules...)...)
}

func (c *Retriever) Validate() error {
	return validator.ValidateStruct(c,
		validation.Field(&c.Type, validation.In(FileType)),
		validation.Field(c.File, c.when(FileType)),
		validation.Field(c.InMemory, c.when(InMemoryType)),
	)
}
