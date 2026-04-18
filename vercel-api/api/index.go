package api

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

// app is initialised once at startup and reused across all requests.
var app = buildApp()

func buildApp() *fiber.App {
	a := fiber.New()

	a.Get("/healthcheck", func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"version": "v1",
		})
	})

	a.Get("/api/resume/v1", adaptor.HTTPHandlerFunc(Resume))

	return a
}

// Handler is the Vercel serverless entrypoint.
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in fiber.Ctx
	r.RequestURI = r.URL.String()

	adaptor.FiberApp(app).ServeHTTP(w, r)
}
