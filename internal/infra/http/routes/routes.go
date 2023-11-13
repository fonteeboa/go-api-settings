package routes

import (
	"golang-api-settings/internal/infra/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// modules Settings
	controllerSettings "golang-api-settings/internal/modules/settings/controller"
	repositoriesSettings "golang-api-settings/internal/modules/settings/repositories"
	serviceSettings "golang-api-settings/internal/modules/settings/services"

	// modules apiSystem
	controllerApi "golang-api-settings/internal/modules/apiSystem/controller"
	repositoriesApi "golang-api-settings/internal/modules/apiSystem/repositories"
	serviceApi "golang-api-settings/internal/modules/apiSystem/services"
)

func ConfigureRoutes(router *gin.Engine) {

	db, errDb := database.NewConnect()

	if errDb != nil {
		log.Fatal(errDb)
	}

	// Configurar o middleware CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Substitua pelo seu frontend URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	router.Use(cors.New(config))

	// Crie instâncias do repositório e do serviço

	// modules Settings
	settingsRepo := repositoriesSettings.NewSettingsRepository(db)
	settingsService := serviceSettings.NewSettingsService(settingsRepo)
	settingsController := controllerSettings.NewSettingsController(settingsService)

	// modules apiSystem
	apiSystemRepo := repositoriesApi.NewApiSystemRepository(db)
	apiSystemService := serviceApi.NewApiSystemService(apiSystemRepo)
	apiSystemController := controllerApi.NewApiController(apiSystemService)

	// Definindo as rotas
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, time.Now().String())
	})

	// modules Settings
	router.GET("/auth", settingsController.GetData)
	router.POST("/auth", settingsController.InsertData)
	router.PUT("/auth", settingsController.UpdateData)

	// modules apiSystem
	router.GET("/itg", apiSystemController.GetData)
	router.DELETE("/itg", apiSystemController.Delete)
	router.POST("/itg", apiSystemController.InsertData)
	router.PUT("/itg", apiSystemController.UpdateData)

}
