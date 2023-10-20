package routes

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
    "log"
    "golang-api-settings/internal/infra/database"
    // modules Settings
    controllerSettings   "golang-api-settings/internal/modules/settings/controller"
    serviceSettings      "golang-api-settings/internal/modules/settings/services"
    repositoriesSettings "golang-api-settings/internal/modules/settings/repositories"
    // modules apiSystem
    controllerApi   "golang-api-settings/internal/modules/apiSystem/controller"
    serviceApi      "golang-api-settings/internal/modules/apiSystem/services"
    repositoriesApi "golang-api-settings/internal/modules/apiSystem/repositories"
)

func ConfigureRoutes(router *gin.Engine) {

    gin.SetMode(gin.ReleaseMode)
	db, errDb := database.NewConnect()

	if errDb != nil {
		log.Fatal(errDb)
	}

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

    router.GET("/data", settingsController.GetData)

    router.GET("/api", apiSystemController.GetData)
}
