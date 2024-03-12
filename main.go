package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	var err error

	// init config
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	PrintJSON(cfg)

	// init router
	router := fiber.New()

	// init minio
	minio, err := NewMinio(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// init controller
	controller := NewController(minio)

	/* Setup HTTPServer*/
	router, err = setupHTTPServer(router)
	if err != nil {
		log.Fatal(err)
	}

	/* Setup Router */
	router, err = setupAPIRouter(router, controller)
	if err != nil {
		log.Fatal(err)
	}

	// server listen
	err = router.Listen(cfg.HTTP_PORT)
	if err != nil {
		log.Fatal(err)
	}
}

func setupHTTPServer(router *fiber.App) (*fiber.App, error) {

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*", // Be cautious with '*', adjust based on your needs
	// 	AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
	// 	AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
	// 	AllowCredentials: false,
	// 	ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type",
	// 	MaxAge:           5200, // Adjust based on your needs
	// }))

	router.Use(cors.New())

	router.Static("/", "./public")

	return router, nil

}

func setupAPIRouter(router *fiber.App, ctlr *Controller) (*fiber.App, error) {

	router.Get("/ping", ctlr.Ping)

	router.Post("/post", ctlr.Post)

	router.Post("/upload", ctlr.Upload)

	return router, nil

}
