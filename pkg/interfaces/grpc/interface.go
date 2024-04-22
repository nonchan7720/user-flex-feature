package grpc

import (
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/ofrep"
	user_flex_feature "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature"
)

type ServiceServer interface {
	user_flex_feature.UserFlexFeatureServiceServer
	ofrep.OFREPServiceServer
}
