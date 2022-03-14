package api

import (
	"context"
	"lynx/pkg/clients/arietes"
	"lynx/pkg/clients/ibis"
	"lynx/pkg/clients/picus"
	"lynx/pkg/clients/rhino"
	"net/http"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	//"github.com/rakyll/statik/fs"

	"lynx/pkg/api/handlers"
	"lynx/pkg/clients/ardea"
	"lynx/pkg/clients/auth"
	"lynx/pkg/clients/gorilla"
	"lynx/pkg/clients/hippo"
	"lynx/pkg/clients/ovis"
	"lynx/pkg/clients/prom"
	"lynx/pkg/clients/s3"
	"lynx/pkg/clients/slav"
	"lynx/pkg/config"

	_ "lynx/pkg/api/prometheus"
	// _ "lynx/pkg/docs"
)

type LynxAPI struct {
	echo *echo.Echo

	config config.LynxConfig
}

func NewLynxAPI(initCtx context.Context, config config.LynxConfig) (*LynxAPI, error) {
	api := &LynxAPI{
		config: config,
	}
	if err := api.init(initCtx); err != nil {
		return nil, err
	}
	return api, nil
}

func (a *LynxAPI) init(initCtx context.Context) (err error) {
	ardeaClient, err := ardea.NewGRPCClient(initCtx, a.config.Ardea)
	if err != nil {
		return err
	}
	hippoClient, err := hippo.NewGRPCClient(initCtx, a.config.Hippo)
	if err != nil {
		return
	}
	gorillaClient, err := gorilla.NewGRPCClient(initCtx, a.config.Gorilla)
	if err != nil {
		return
	}
	s3Client, err := s3.NewMinioS3Client(a.config.S3)
	if err != nil {
		return
	}
	ovisClient, err := ovis.NewRabbitMQClient(a.config.Ovis)
	if err != nil {
		return
	}
	slavClient, err := slav.NewGRPCClient(initCtx, a.config.Slav)
	if err != nil {
		return
	}
	ibisClient, err := ibis.NewGRPCClient(initCtx, a.config.Ibis)
	if err != nil {
		return
	}
	rhinoClient, err := rhino.NewGRPCClient(initCtx, a.config.Rhino)
	if err != nil {
		return
	}
	arietesClient, err := arietes.NewKafkaClient(a.config.Arietes)
	if err != nil {
		return
	}
	picusClient, err := picus.NewGRPCClient(initCtx, a.config.Picus)
	if err != nil {
		return
	}
	promClient, err := prom.NewPromClient(a.config.Prometheus)
	if err != nil {
		return err
	}

	authClient := auth.NewDjangoAuthClient(a.config.Auth)

	api := handlers.NewHandlers(
		a.config.API.Users, ardeaClient, gorillaClient, s3Client, hippoClient, ovisClient, authClient, slavClient,
		ibisClient, rhinoClient, arietesClient, picusClient, promClient,
	)

	e := echo.New()

	// Enable metrics middleware
	p := prometheus.NewPrometheus("lynx", nil)
	p.Use(e)

	// Logger
	e.Use(middleware.Logger())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"https://app.deepmux.com",
			"http://app.deepmux.com",
			"http://polygon.endevir.ru:3000",
			"http://localhost:3000",
		},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	v1 := e.Group("v1")
	v1A := e.Group("v1-a")

	initV1(api, v1)
	initV1a(api, v1A)

	initDocs(e)

	a.echo = e

	return
}

func initV1a(api *handlers.Handlers, g *echo.Group) {
	g.Use(api.ErrorsMiddleware)
	g.GET("/function/:name/stream_logs", api.GetStreamLogsHandler)
}

func initV1(api *handlers.Handlers, g *echo.Group) {
	g.Use(api.ErrorsMiddleware)
	g.Use(api.AuthMiddleware)

	g.PUT("/model/:name", api.PutModelHandler)
	g.POST("/model/:name", api.PostModelHandler)
	g.GET("/model/:name", api.GetModelHandler)
	g.DELETE("/model/:name", api.DeleteModelHandler)
	g.GET("/model", api.ListModelsHandler)

	g.POST("/model/:name/run", api.PostRunModelHandler)

	g.GET("/container", api.ListContainersHandler)
	g.GET("/container/:name", api.GetContainerHandler)
	g.POST("/container/:name", api.PostContainerHandler)
	g.PATCH("/container/:name", api.PatchContainerRequest)
	g.DELETE("/container/:name", api.DeleteContainerHandler)

	g.PUT("/function/:name", api.PutFunctionHandler)
	g.POST("/function/:name", api.PostFunctionHandler)
	g.GET("/function/:name", api.GetFunctionHandler)
	g.DELETE("/function/:name", api.DeleteFunctionHandler)
	g.GET("/function", api.ListFunctionsHandler)
	g.GET("/function_envs", api.GetFunctionEnvHandler)

	g.POST("/function/:name/run", api.PostRunFunctionHandler)
	g.GET("/function/:name/metrics", api.GetMetricsHandler)

	g.GET("/function/:name/full_logs", api.GetLogsHandler)
}

func initDocs(_ *echo.Echo) { /*
		statikFS, err := fs.New()
		if err != nil {
			panic(err)
		}
		staticServer := http.FileServer(statikFS)
		eh := echo.WrapHandler(http.StripPrefix("/docs/v1", staticServer))
		e.GET("/docs/v1*", eh)*/
}

func (a *LynxAPI) Run() {
	a.echo.HideBanner = true
	a.echo.Logger.Fatal(a.echo.Start(a.config.API.Bind))
}
