package di

import (
	_ "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	_ "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/feature"
	_ "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/feature/retriever"
	_ "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/fsm"
	_ "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/tracking"
	_ "github.com/nonchan7720/user-flex-feature/pkg/interfaces/api/controller"
	_ "github.com/nonchan7720/user-flex-feature/pkg/interfaces/api/gin"
	_ "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc"
	_ "github.com/nonchan7720/user-flex-feature/pkg/services/feature"
)
