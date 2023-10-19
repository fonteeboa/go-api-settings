package routes

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
    "log"
    "golang-api-settings/internal/infra/database"
    // Domain Settings
    controllerSettings   "golang-api-settings/internal/domain/settings/controller"
    serviceSettings      "golang-api-settings/internal/domain/settings/services"
    repositoriesSettings "golang-api-settings/internal/domain/settings/repositories"
    // Domain apiSystem
    controllerApi   "golang-api-settings/internal/domain/apiSystem/controller"
    serviceApi      "golang-api-settings/internal/domain/apiSystem/services"
    repositoriesApi "golang-api-settings/internal/domain/apiSystem/repositories"
)

func ConfigureRoutes(router *gin.Engine) {

    gin.SetMode(gin.ReleaseMode)
	db, errDb := database.NewConnect()

	if errDb != nil {
		log.Fatal(errDb)
	}

    // Crie instâncias do repositório e do serviço

    // Domain Settings
    settingsRepo := repositoriesSettings.NewSettingsRepository(db)
    settingsService := serviceSettings.NewSettingsService(settingsRepo)
    settingsController := controllerSettings.NewSettingsController(settingsService)
    
    // Domain apiSystem
    apiSystemRepo := repositoriesApi.NewApiSystemRepository(db)
    apiSystemService := serviceApi.NewApiSystemService(apiSystemRepo)
    apiSystemController := controllerApi.NewApiController(apiSystemService)

    // Definindo as rotas
    router.GET("/health", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, time.Now().String())
    })

    router.GET("/data", settingsController.GetData)

    router.GET("/api", apiSystemController.GetData)    
}
