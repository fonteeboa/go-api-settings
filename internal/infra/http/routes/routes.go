package routes

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
    "golang-api-settings/internal/domain/settings/Services"
    repositorysSettings "golang-api-settings/internal/domain/settings/repositorys"
    "golang-api-settings/internal/domain/apiSystem/Services"
    repositorysApi "golang-api-settings/internal/domain/apiSystem/repositorys"
    "golang-api-settings/internal/infra/database"
    "log"
)

func ConfigureRoutes(router *gin.Engine) {

    gin.SetMode(gin.ReleaseMode)
	db, errDb := database.NewConnect()

	if errDb != nil {
		log.Fatal(errDb)
	}

    // Crie instâncias do repositório e do serviço
    settingsRepo := repositorysSettings.NewSettingsRepository(db)
    settingsService := settings.NewSettingsService(settingsRepo)
    // Domain apiSystem
    apiSystemRepo := repositorysApi.NewApiSystemRepository(db)
	apiSystemService := apiSystem.NewApiSytemService(apiSystemRepo)


    router.GET("/health", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, time.Now().String())
    })

    router.GET("/data", settingsService.GetDataService)

    router.GET("/api", apiSystemService.GetDataService)    
}
