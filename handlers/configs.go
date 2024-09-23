package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"test/models"

	"github.com/gofiber/fiber/v2"
)

func CreateConv(c *fiber.Ctx) error {
	config := new(models.Config)

	if err := c.BodyParser(config); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Ошибка парсинга данных")
	}

	filePath := "configs/config-" + config.NameConfig + ".json"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		// Преобразуем структуру в JSON
		jsonData, err := json.MarshalIndent(config, "", "  ") // MarshalIndent для красивого форматирования
		if err != nil {
			fmt.Println("Ошибка сериализации:", err)
			// return
		}

		// Записываем JSON в файл
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Ошибка создания файла:", err)
			// return
		}
		defer file.Close()

		_, err = file.Write(jsonData)
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			// return
		}

		fmt.Println("JSON файл успешно создан!")
	} else {
		return c.Status(500).SendString(fmt.Sprintf("такой файл уже существует: %v", err))
	}

	return c.SendString("Config created successfully")
}
