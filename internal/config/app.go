package config

import (
	"github.com/go-chi/chi/v5"
	"github.com/irinaponzi/package-tracker/internal/batch"
	_package "github.com/irinaponzi/package-tracker/internal/package"
	"github.com/irinaponzi/package-tracker/internal/tracking"
	"net/http"

	//_ "github.com/irinaponzi/package-tracker/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// NewAppDefault creates a default app configuration
func NewAppDefault() *AppConfig {
	return &AppConfig{
		Mux:           chi.NewRouter(),
		ServerAddress: ":8080",
	}
}

// AppConfig is a struct that represents the configuration for the app
type AppConfig struct {
	ServerAddress string
	Mux           *chi.Mux
	Handlers      *HandlerContainer
}

// HandlerContainer is a struct that contains all the handlers
type HandlerContainer struct {
	TrackingHandler *tracking.TrackingHandler
	PackageHandler  *_package.PackageHandler
	BatchHandler    *batch.BatchHandler
}

// SetDependencies is a function that sets the app dependencies
func (app *AppConfig) SetDependencies() {
	// - connect to sql database
	database, err := GetConnection()
	if err != nil {
		panic(err)
	}
	// - repository
	rpTracking := tracking.NewTrackingRepository(database)
	rpPackage := _package.NewPackageRepository(database)
	rpBatch := batch.NewBatchRepository(database)
	// - service
	svTracking := tracking.NewTrackingService(rpTracking)
	svPackage := _package.NewPackageService(rpPackage)
	svBatch := batch.NewBatchService(rpBatch)
	// - handler
	hdTracking := tracking.NewTrackingHandler(svTracking)
	hdPackage := _package.NewPackageHandler(svPackage)
	hdBatch := batch.NewBatchHandler(svBatch)

	app.Handlers = &HandlerContainer{
		TrackingHandler: hdTracking,
		PackageHandler:  hdPackage,
		BatchHandler:    hdBatch,
	}
}

// SetMappings is a function that sets the app mappings
func (app *AppConfig) SetMappings() {
	// - health check
	app.Mux.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`pong`))
	})

	// - api routes
	app.Mux.Route("/api/v1", func(rt chi.Router) {
		// - tracking
		rt.Route("/tracking", func(r chi.Router) {
			r.Post("/", app.Handlers.TrackingHandler.Create())
		})
		// - packages
		rt.Route("/packages", func(r chi.Router) {
			r.Get("/", app.Handlers.PackageHandler.GetAll())
		})
		// - batches
		rt.Route("/batches", func(r chi.Router) {
			r.Get("/", app.Handlers.BatchHandler.GetAll())
		})
	})

	// - swagger
	// todo: it is not configured yet
	app.Mux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
}

// Run is a method that runs the application
func (app *AppConfig) Run() error {
	err := http.ListenAndServe(app.ServerAddress, app.Mux)
	return err
}
