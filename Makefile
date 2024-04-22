.PHONY: tools/oapi
tools/oapi:
	@go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0

I18N_MODE=
I18N_ROOT_PATH=
.PHONY: run/i18n-extra
run/i18n-extra:
	@go run pkg/infrastructure/i18n/cli/main.go "$(I18N_MODE)" "$(I18N_ROOT_PATH)" pkg/infrastructure/i18n/assets/

.PHONY: i18n-original-extra
i18n-original-extra:
	$(MAKE) run/i18n-extra I18N_MODE=original I18N_ROOT_PATH=.

.PHONY: i18n-validation-extra
i18n-validation-extra:
	$(MAKE) run/i18n-extra I18N_MODE=validation I18N_ROOT_PATH="$(GOPATH)/pkg/mod/github.com/go-ozzo/ozzo-validation/;."

.PHONY: i18n-extra
i18n-extra:
	$(MAKE) i18n-original-extra
	$(MAKE) i18n-validation-extra

BUF_VERSION=v1.26.1
PROTOC_GEN_APIGW_VERSION=v0.1.30

.PHONY: protoc/dev
protoc/dev:
	@go install github.com/bufbuild/buf/cmd/buf@${BUF_VERSION}
	@go install github.com/ductone/protoc-gen-apigw@${PROTOC_GEN_APIGW_VERSION}
	@if !(type "swagger2openapi" > /dev/null 2>&1); then sudo npm i -g swagger2openapi@7.0.8 ; fi

.PHONY: protoc/uff
protoc/uff: protoc/dev
	@buf generate --template docs/schema/user-flex-feature/buf.gen.go-server.yaml --path docs/schema/user-flex-feature
	@swagger2openapi --outfile docs/schema/user-flex-feature/v1/schema.openapi.yaml docs/schema/user-flex-feature/v1/schema.swagger.yaml
	@go run tools/openapi/main.go docs/schema/user-flex-feature/v1/schema.openapi.yaml docs/schema/ofrep/v1/openapi.yaml  > docs/schema/openapi.yaml

.PHONY: protoc/uff
protoc/ofrep: protoc/dev
	@buf generate --template docs/schema/ofrep/buf.gen.go-server.yaml --path docs/schema/ofrep

.PHONY: codegen/api
codegen/api: tools/oapi
	@oapi-codegen --config=pkg/interfaces/api/generated/config.yaml docs/schema/openapi.yaml
