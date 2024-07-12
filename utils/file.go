package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var DefaultPathAssetImage = "./public/covers/"

func HandleSingleFile(ctx *fiber.Ctx) error {
	// Handle FIle
	file, errFile := ctx.FormFile("cover")

	if errFile != nil {
		log.Println("error file = ", errFile)
	}

	var filename string

	if file != nil {

		// Validation File
		errCheckContentType := CheckContentType(file, "image/jpeg", "image/png")
		if errCheckContentType != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": errCheckContentType.Error(),
			})
		}

		// Validation Size
		if file.Size > 2000000 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "Max Size File 2 MB",
			})
		}

		uuid, _ := uuid.NewV6()
		fileExtention := filepath.Ext(file.Filename)
		filename = fmt.Sprintf("%s%s", uuid, fileExtention)

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename))

		if errSaveFile != nil {
			log.Println("error save file = ", errSaveFile)
		}
	} else {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Upload File Failed",
		})
	}

	ctx.Locals("filename", filename)

	return ctx.Next()
}

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println("error file = ", err)
	}

	files := form.File["images"]

	var filenames []string

	for _, file := range files {
		var filename string

		if file != nil {
			uuid, _ := uuid.NewV6()
			fileExtention := filepath.Ext(file.Filename)
			filename = fmt.Sprintf("%s%s", uuid, fileExtention)

			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename))

			if errSaveFile != nil {
				log.Println("error save file = ", errSaveFile)
			}
		} else {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":   true,
				"message": "Upload File Failed",
			})
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}
		ctx.Locals("filenames", filenames)
	}
	log.Println(filenames)

	return ctx.Next()
}

func HandleRemoveFile(filename string, path ...string) error {

	if len(path) > 0 {
		DefaultPathAssetImage = path[0]
	}
	err := os.Remove(DefaultPathAssetImage + filename)
	if err != nil {
		log.Println("error remove file = ", err)
		return err
	}
	return nil
}

func CheckContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			if contentType == file.Header.Get("Content-Type") {
				return nil
			}
		}
		return fmt.Errorf("invalid content type, must be one of: %s", contentTypes)
	} else {
		return errors.New("not found content type")
	}
}
