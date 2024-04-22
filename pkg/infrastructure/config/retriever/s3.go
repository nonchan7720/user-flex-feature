package retriever

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/validator"
)

type S3 struct {
	Bucket string `yaml:"uploadBucketName"`
}

func (c *S3) Validation() error {
	return validator.ValidateStruct(c,
		validation.Field(&c.Bucket, validation.Required),
	)
}
