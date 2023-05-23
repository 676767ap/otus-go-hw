package app

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/676767ap/project/infrastructure/database"
	"github.com/676767ap/project/internal/config"
	"github.com/676767ap/project/internal/services"
	"github.com/676767ap/project/internal/usecase"
	"github.com/676767ap/project/util/log"
	"gorm.io/gorm"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	gormsessions "github.com/gin-contrib/sessions/gorm"
)

type App struct {
	cfg          *config.Config
	router       *gin.Engine
	rep          usecase.Repos
	db           *gorm.DB
	queue        services.Exchange
	exchangeName string
}

func NewApp(cfg *config.Config) (App, error) {

	if !cfg.DevMode {
		gin.SetMode(gin.ReleaseMode)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()
	queue, err := services.NewQueue(ctx, cfg)
	if err != nil {
		return App{}, err
	}
	defer queue.Close(ctx)
	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		return App{}, err
	}
	store := gormsessions.NewStore(db, true, []byte(cfg.Session.SessionKey))
	router := gin.Default()
	router.Use(sessions.Sessions(cfg.Session.CookieName, store))
	a := App{
		cfg:          cfg,
		router:       router,
		rep:          usecase.NewRepos(db),
		db:           db,
		queue:        queue,
		exchangeName: cfg.RabbitMQ.Exchange.Name,
	}
	api := router.Group("/api")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	/** Objects **/
	api.POST("/add-banner-to-slot/{banner_id}/{slot_id}", a.addBannerToSlot)
	api.POST("/remove-banner-from-slot/{banner_id}/{slot_id}", a.removeBannerFromSlot)
	api.POST("/click-on-banner/{banner_id}/{slot_id}/{soc_group_id}", a.clickOnBanner)
	api.POST("/choose-banner-for-slot/{banner_id}/{slot_id}", a.chooseBannerForSlot)

	return a, nil
}

func (a *App) Run() error {
	err := a.router.Run(a.cfg.GetAddressWithPort())
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	srv := &http.Server{
		Addr:    a.cfg.GetAddressWithPort(),
		Handler: a.router,
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server Shutdown:", err)
	}
	log.Info("Server exiting")

	if err := database.CloseDatabase(a.db); err != nil {
		log.Error("DB Shutdown:", err)
	}
}

type ErrorStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ServeError(c *gin.Context, code int, message string) {
	c.JSON(code, ErrorStruct{code, message})
}
