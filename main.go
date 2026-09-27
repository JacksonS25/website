package main

import (
	"log"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Serve your HTML/CSS/JS from a folder named "public"
	app.Static("/", "./public")

	// API endpoint triggered by your website buttons
	app.Post("/api/run/:project", func(c *fiber.Ctx) error {
		project := c.Params("project")
		var cmd *exec.Cmd

		// Route to the correct language/script based on the project name
		switch project {
		case "python-scraper":
			cmd = exec.Command("python3", "./projects/scraper.py")
		case "go-math-tool":
			cmd = exec.Command("go", "run", "./projects/math.go")
		default:
			return c.Status(404).SendString("Project not found")
		}

		// Execute the script and capture all terminal output
		out, err := cmd.CombinedOutput()
		if err != nil {
			return c.Status(500).SendString(string(out) + "\nExecution Error: " + err.Error())
		}

		return c.SendString(string(out))
	})

	log.Fatal(app.Listen(":3000"))
}
