package app

import (
	authHandler "github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/handler"
	authService "github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/service"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/user/repository/postgresql"
	userService "github.com/ahmetcancicek/pomodorogo-server/internal/app/user/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
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
	db         *gorm.DB
}

func NewApp() (*pomodoroServerApplication, error) {
	application := new(pomodoroServerApplication)
	return application, nil
}

func (app *pomodoroServerApplication) Init() error {

	router := mux.NewRouter()

	// TODO: We should get these property from config
	app.router = router
	app.httpServer.Addr = ":8500"
	app.httpServer.Handler = app.router

	// Auth Package
	userRepository := postgresql.NewPostgreSQLUserRepository(app.db)
	userLogic := userService.NewUserService(userRepository)
	authLogic := authService.NewAuthService()
	authHandler.NewAuthHandler(router, userLogic, authLogic)

	return nil
}

func (app *pomodoroServerApplication) StartDB() error {

	// TODO: We should get these property from config
	dsn := "host=localhost user=postgres password=password dbname=pomodorogo-server port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.User{})

	if err != nil {
		log.Fatal(err)
	}
	app.db = db

	return nil
}

func (app *pomodoroServerApplication) StartHttpServer() error {
	if err := app.httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return nil
}
