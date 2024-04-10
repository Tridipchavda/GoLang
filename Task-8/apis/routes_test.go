package apis

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRoutes(t *testing.T) {
	app := fiber.New()

	r, _ := http.NewRequest(http.MethodGet, "/", nil)

	w, err := app.Test(r)
	if err != nil {
		t.Errorf("Error By Error Kya kar raha hai")
	}
	if w.StatusCode != 200 {
		t.Errorf("Expected 200, got %d", w.StatusCode)
	}
}
