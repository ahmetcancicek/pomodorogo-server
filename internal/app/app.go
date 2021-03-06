package app

import (
	accountRepository "github.com/ahmetcancicek/pomodorogo-server/internal/app/account/repository/postgresql"
	accountService "github.com/ahmetcancicek/pomodorogo-server/internal/app/account/service"
	authHandler "github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/handler"
	authService "github.com/ahmetcancicek/pomodorogo-server/internal/app/auth/service"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/model"
	pomodoroHandler "github.com/ahmetcancicek/pomodorogo-server/internal/app/pomodoro/handler"
	pomodoroRepository "github.com/ahmetcancicek/pomodorogo-server/internal/app/pomodoro/repository/postgresql"
	pomodoroService "github.com/ahmetcancicek/pomodorogo-server/internal/app/pomodoro/service"
	tagHandler "github.com/ahmetcancicek/pomodorogo-server/internal/app/tag/handler"
	tagRepository "github.com/ahmetcancicek/pomodorogo-server/internal/app/tag/repository/postgresql"
	tagService "github.com/ahmetcancicek/pomodorogo-server/internal/app/tag/service"
	"github.com/ahmetcancicek/pomodorogo-server/internal/app/utils"
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

	// Router
	router := mux.NewRouter()
	// Logger
	logger := utils.NewLogger()
	// Configuration
	configs := utils.NewConfigurations(logger)

	// TODO: Refactor Subrouter

	// Auth Package
	accountRepo := accountRepository.NewPostgreSQLAccountRepository(logger, app.db)
	accountServ := accountService.NewAccountService(accountRepo)
	authServ := authService.NewAuthService(logger, configs)
	authHand := authHandler.NewAuthHandler(router.NewRoute().Subrouter(), logger, accountServ, authServ)

	// Tag Package
	tagRepo := tagRepository.NewPostgreSQLTagRepository(logger, app.db)
	tagServ := tagService.NewTagService(tagRepo)
	tagHandler.NewTagHandler(router.NewRoute().Subrouter(), logger, tagServ, authHand.MiddlewareValidateAccessToken)

	// Pomodoro Package
	pomodoroRepo := pomodoroRepository.NewPostgreSQLPomodoroRepository(logger, app.db)
	pomodoroServ := pomodoroService.NewStatisticService(pomodoroRepo, tagRepo)
	pomodoroHandler.NewStatisticHandler(router.NewRoute().Subrouter(), logger, pomodoroServ, authHand.MiddlewareValidateAccessToken)

	//
	app.router = router
	app.httpServer.Addr = configs.ServerAddress
	app.httpServer.Handler = app.router

	return nil
}

func (app *pomodoroServerApplication) StartDB() error {

	// TODO:
	logger := utils.NewLogger()
	configs := utils.NewConfigurations(logger)

	dsn := "host=" + configs.DBHost + " port=" + configs.DBPort + " user=" + configs.DBUser + " dbname=" + configs.DBName + " password=" + configs.DBPass
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	app.db = db

	err = app.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (app *pomodoroServerApplication) Migrate() error {
	err := app.db.AutoMigrate(&model.User{}, &model.Tag{}, &model.Setting{}, &model.Pomodoro{})
	if err != nil {
		return err
	}
	return nil
}

func (app *pomodoroServerApplication) StartHttpServer() error {
	if err := app.httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return nil
}
