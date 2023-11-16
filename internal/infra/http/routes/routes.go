package routes

import (
	"golang-api-settings/internal/infra/database"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// modules Authorization
	controllerAuthorization "golang-api-settings/internal/modules/authorization/controller"
	repositoriesAuthorization "golang-api-settings/internal/modules/authorization/repositories"
	serviceAuthorization "golang-api-settings/internal/modules/authorization/services"

	// modules integration
	controllerApi "golang-api-settings/internal/modules/integration/controller"
	repositoriesApi "golang-api-settings/internal/modules/integration/repositories"
	serviceApi "golang-api-settings/internal/modules/integration/services"
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

	// modules Authorization
	authorizationRepo := repositoriesAuthorization.NewAuthorizationRepository(db)
	authorizationService := serviceAuthorization.NewAuthorizationService(authorizationRepo)
	authorizationController := controllerAuthorization.NewAuthorizationController(authorizationService)

	// modules integration
	integrationRepo := repositoriesApi.NewIntegrationRepository(db)
	integrationService := serviceApi.NewIntegrationService(integrationRepo)
	integrationController := controllerApi.NewApiController(integrationService)

	// Definindo as rotas
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, time.Now().String())
	})

	// modules Authorization
	router.GET("/auth", authorizationController.GetData)
	router.POST("/auth", authorizationController.InsertData)
	router.PUT("/auth", authorizationController.UpdateData)

	// modules integration
	router.GET("/itg", integrationController.GetData)
	router.DELETE("/itg", integrationController.Delete)
	router.POST("/itg", integrationController.InsertData)
	router.PUT("/itg", integrationController.UpdateData)

}
