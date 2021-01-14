package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"

	authHandler "github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/handler"
	authRepo "github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/repository"
	authService "github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/service"
)

type App interface {
	// Start all dependencies services
	Init() error
	// Start http server
	StartHttpServer() error
}

type pomodoroServerApplication struct {
	router     *mux.Router
	httpServer http.Server
	dbConn     *sql.DB
}

func NewApp() (*pomodoroServerApplication, error) {
	application := new(pomodoroServerApplication)
	return application, nil
}

func (app *pomodoroServerApplication) Init() error {

	timeoutContext := time.Duration(10)
	router := mux.NewRouter()

	app.router = router
	app.httpServer.Addr = ":8500"
	app.httpServer.Handler = app.router

	// Auth Package
	authRepo := authRepo.NewPostgreSQLUserRepository(app.dbConn)
	authService := authService.NewUserService(authRepo, timeoutContext)
	authHandler.NewAuthHandler(router, authService)

	return nil
}

func (app *pomodoroServerApplication) StartDB() error {
	connStr := "postgres://postgres:password@localhost/pomodorogo-server?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Defer

	return nil
}

func (app *pomodoroServerApplication) StartHttpServer() error {
	if err := app.httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return nil
}
