package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Define a struct to read the JSON incoming from the frontend
type RunRequest struct {
	Args string `json:"args"`
}

func main() {
	app := fiber.New()
	app.Static("/", "./public")

	app.Post("/api/run/:project", func(c *fiber.Ctx) error {
		project := c.Params("project")

		// Parse the incoming JSON to get the arguments
		var req RunRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString("Invalid request body format")
		}

		var cmd *exec.Cmd

		switch project {
		case "song_set_maker":
			// 1. Base command and script path
			commandArgs := []string{"./projects/Song_Set_Maker/main.py"}

			// 2. If the user provided inputs, split them by comma and add them as arguments
			if strings.TrimSpace(req.Args) != "" {
				// Split by comma and clean up spaces
				inputs := strings.Split(req.Args, ",")
				for _, input := range inputs {
					cleanedInput := strings.TrimSpace(input)
					if cleanedInput != "" {
						commandArgs = append(commandArgs, cleanedInput)
					}
				}
			}

			// 3. Construct the final exec.Cmd
			// (python3 is the executable, commandArgs contains the script and the arguments)
			cmd = exec.Command("python3", commandArgs...)

		default:
			return c.Status(404).SendString("Project not found")
		}

		out, err := cmd.CombinedOutput()
		if err != nil {
			return c.Status(500).SendString(string(out) + "\nExecution Error: " + err.Error())
		}

		return c.SendString(string(out))
	})

	log.Fatal(app.Listen(":3000"))
}
