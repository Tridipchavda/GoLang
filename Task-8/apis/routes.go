package apis

import (
	"log"
	"net/http"

	"github.com/Tridipchavda/FiberWithMongoDB/controllers"
	"github.com/Tridipchavda/FiberWithMongoDB/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// App struct holding all Controllers
type App struct {
	cs *controllers.Controller
}

// Function to return App struct with Controllers and DB in controllers
func NewApp(db *mongo.Client) *App {
	return &App{cs: controllers.NewController(db)}
}

// Initialize the Server and Hadling the Routes
func (a *App) InitializeRoutes(app *fiber.App) {

	// Get The Content From Database
	app.Get("/", func(c *fiber.Ctx) error {
		// Get all Data from controller
		content, err := a.cs.GetAllContentData()
		// Send Data to User/Client
		c.Status(http.StatusOK).JSON(content)
		return err
	})

	// Get the Data by Id From Database
	app.Get("/:id", func(c *fiber.Ctx) error {
		// Fetch Id
		id := c.Params("id")

		// Get One Document from controllers
		res, err := a.cs.GetOneContentData(id)

		// Handling Error from Controller
		if err != nil {
			c.SendString(err.Error())
		}
		// Send Data to user
		c.Status(http.StatusOK).JSON(res)
		return nil
	})

	// Create the Content in Mongo Database
	app.Post("/", func(c *fiber.Ctx) error {
		// Parsing the Body to fetch the JSON content
		var content models.Content
		err := c.BodyParser(&content)

		// Handling Error from Body parser
		if err != nil {
			c.SendString(err.Error())
			return err
		}
		log.Println(content)

		// Insert the Content by controller
		err = a.cs.InsertContent(&content)

		// Handling Error from controller
		if err != nil {
			return err
		}

		// Status Ok with INSERT message
		c.Status(http.StatusOK).SendString("Insert Successfully")
		return nil
	})

	// Delete the Content in Mongo Database by id
	app.Delete("/:id", func(c *fiber.Ctx) error {
		// Fetch id from Params
		id := c.Params("id")

		// Delete the content with Controllers
		err := a.cs.DeleteContent(id)

		// Handling the error by controller
		if err != nil {
			c.Status(http.StatusOK).SendString(err.Error())
			return err
		}

		// Status ok and Message for DELETE to user
		c.Status(http.StatusOK).SendString("Deleted Successfully")
		return err
	})

	app.Put("/:id", func(c *fiber.Ctx) error {
		// Fetch id from Params
		id := c.Params("id")
		var content models.Content

		// Parse the Body to Fetrch JSON
		c.BodyParser(&content)

		// Call the Controller to update content
		err := a.cs.UpdateContent(content, id)
		if err != nil {
			c.SendString(err.Error())
			return err
		}

		// Send Message to Client
		c.SendString("Updated Successfully")
		return nil
	})

	// Listening app on port 3333
	app.Listen(":3333")
}
