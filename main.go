package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/nikola43/pdexrefapi/controllers"
	"github.com/nikola43/pdexrefapi/db"
)

var httpServer *fiber.App

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")

	// system config
	numCpu := runtime.NumCPU()
	usedCpu := numCpu
	runtime.GOMAXPROCS(usedCpu)
	fmt.Println("")
	fmt.Println(color.YellowString("  ----------------- System Info -----------------"))
	fmt.Println(color.CyanString("\t    Number CPU cores available: "), color.GreenString(strconv.Itoa(numCpu)))
	fmt.Println(color.MagentaString("\t    Used of CPU cores: "), color.YellowString(strconv.Itoa(usedCpu)))
	fmt.Println(color.MagentaString(""))

	// Initialize the database
	db.InitializeDatabase(MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE, MYSQL_HOST, MYSQL_PORT, false)
	InitializeHttpServer()
}

func InitializeHttpServer() {

	httpServer = fiber.New(fiber.Config{
		BodyLimit: 2000 * 1024 * 1024, // this is the default limit of 2MB
	})

	//httpServer.Use(middlewares.XApiKeyMiddleware)
	// httpServer.Use(cors.New(cors.Config{
	// 	AllowOrigins: "https://google.com",
	// }))

	httpServer.Use(jwtlogger.New())
	httpServer.Use(cors.New(cors.Config{}))

	HandleRoutes(httpServer)

	httpServer.Listen(":2001")
}

func HandleRoutes(router fiber.Router) {
	router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": "ok",
		})
	})

	// router.Get("/users", controllers.GetUser)
	router.Post("/users", controllers.GetOrCreate)
	router.Post("/users/addReferral", controllers.AddReferral)
}
