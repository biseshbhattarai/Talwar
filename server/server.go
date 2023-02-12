package main

import (
	"fmt"
	"log"

	"github.com/biseshbhattarai/talwar/controllers"
	"github.com/biseshbhattarai/talwar/cronjob"
	database "github.com/biseshbhattarai/talwar/db"
	"github.com/biseshbhattarai/talwar/migrations"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	loadEnv()
	loadDatabase()
	fmt.Println("Starting Talwar")
	new_cronjob := cron.New()
	new_cronjob.AddFunc("0 */1 * * * *", func() {
		fmt.Println("Checking for scheduled scans")
		c, _ := cronjob.RunScheduler()
		c.Start()
	})
	new_cronjob.Start()
	SetupRouter()

}

func loadDatabase() {
	database.InitDB()
	migrations.Migrate()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func SetupRouter() {

	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length"},

		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))
	publicRoutes := router.Group("/api")
	publicRoutes.POST("/targets", controllers.RegisterTarget)
	publicRoutes.GET("/targets", controllers.ListTargets)
	publicRoutes.GET("/targets/:target_id", controllers.ListDomainsByTargetID)
	publicRoutes.POST("/targets/:target_id/startScan", controllers.StartScan)
	publicRoutes.GET("/targets/:target_id/subdomains", controllers.ListSubDomainsByTargetID)
	publicRoutes.POST("/targets/:target_id/organization", controllers.RegisterOrganizationRepos)
	publicRoutes.GET("/targets/:target_id/organization", controllers.ListOrganizationRepos)
	publicRoutes.POST("/targets/:target_id/organization/startSearchRepo", controllers.SearchOrganizationReposFromGithub)
	publicRoutes.GET("/targets/organization/:organization_id", controllers.ListRepositoryByOrganizationID)
	publicRoutes.GET("/targets/:target_id/organization/repos", controllers.ListOrganizationReposByTargetID)
	publicRoutes.POST("/targets/:target_id/scanSchedule", controllers.RegisterScanSchedule)
	publicRoutes.GET("/targets/scanSchedule", controllers.ListScanSchedules)
	publicRoutes.GET("/targets/:target_id/scanHistory", controllers.ListScanHistory)
	publicRoutes.POST("/targets/:target_id/nucleiScan", controllers.StartNucleiScan)
	publicRoutes.POST("/targets/:target_id/portScan", controllers.StartPortScan)
	publicRoutes.GET("/targets/:target_id/ports", controllers.ListPortScan)
	publicRoutes.GET("/targets/:target_id/vulnerabilities", controllers.ListNucleiScan)
	publicRoutes.POST("/targets/:target_id/startGithubScan", controllers.StartTruffleScan)
	publicRoutes.GET("/targets/:target_id/scanResults", controllers.ListTruffleScan)
	publicRoutes.POST("/targets/emails", controllers.StoreEmail)
	publicRoutes.POST("/targets/sendEmail", controllers.SendEmails)
	router.Run(":8000")
	fmt.Println("Server running on port 8000")

}
