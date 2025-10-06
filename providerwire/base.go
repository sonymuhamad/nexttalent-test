package providerwire

import (
	"github.com/sonymuhamad/nexttalent-test/timeapi"
	"net/http"

	"github.com/sonymuhamad/nexttalent-test/config"
	"github.com/sonymuhamad/nexttalent-test/httpserver"
	"github.com/sonymuhamad/nexttalent-test/repository"
	"github.com/sonymuhamad/nexttalent-test/usecase"
)

func ProvideApiServer(cfg config.EnvConfig) *http.Server {
	db := provideGormSQLDatabase(cfg)
	personRepository := repository.NewPerson(db)
	personUsecase := usecase.NewPerson(personRepository)
	timeApiClient := timeapi.NewTimeApiClient(cfg)
	handler := httpserver.NewHandler(personUsecase, timeApiClient)
	httpServer := provideHttpServer(cfg, handler)

	return httpServer
}
