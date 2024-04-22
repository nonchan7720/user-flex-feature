package cmd

import validation "github.com/go-ozzo/ozzo-validation/v4"

type commonArgs struct {
	configFilePath string
}

func (args *commonArgs) Validate() error {
	return validation.ValidateStruct(args,
		validation.Field(&args.configFilePath, validation.Required),
	)
}
