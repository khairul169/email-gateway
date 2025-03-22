package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wneessen/go-mail"
)

type SMTPConfig struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
}

var apiKeys map[string]SMTPConfig

func loadConfig() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&apiKeys); err != nil {
		return err
	}
	return nil
}

func sendEmailHandler(c *fiber.Ctx) error {
	apiKey := c.Get("X-API-Key")

	config, exists := apiKeys[apiKey]
	if !exists {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid API key"})
	}

	email := c.FormValue("email")
	title := c.FormValue("title")
	content := c.FormValue("content")

	if email == "" || title == "" || content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	message := mail.NewMsg()
	if err := message.From(config.Username); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to set From address")
	}
	if err := message.To(email); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "failed to set To address")
	}
	message.Subject(title)
	message.SetBodyString(mail.TypeTextHTML, content)

	if file, err := c.FormFile("attachment"); err == nil {
		src, err := file.Open()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to open attachment file")
		}

		defer src.Close()

		tempFile, err := os.CreateTemp("/tmp", "attachment-*")
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to create temporary attachment file")
		}

		defer os.Remove(tempFile.Name())
		io.Copy(tempFile, src)
		message.AttachFile(tempFile.Name(), mail.WithFileName(file.Filename))
	}

	// Deliver the mails via SMTP
	client, err := mail.NewClient(config.Host,
		mail.WithPort(config.Port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithUsername(config.Username), mail.WithPassword(config.Password),
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("failed to create new mail delivery client: %s", err))
	}
	if err := client.DialAndSend(message); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("failed to deliver mail: %s", err))
	}

	return c.JSON(fiber.Map{"success": "Email sent successfully"})
}

func main() {
	if err := loadConfig(); err != nil {
		fmt.Println("Failed to load config:", err)
		os.Exit(1)
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Post("/send-email", sendEmailHandler)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen(host + ":" + port))
}
