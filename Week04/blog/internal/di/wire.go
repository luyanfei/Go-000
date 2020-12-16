// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"blog/internal/dao"
	"blog/internal/server/grpc"
	"blog/internal/server/http"
	"blog/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
