package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	cli "github.com/jawher/mow.cli"
)

const cloudinaryURLEnvVar = "CLOUDINARY_URL"

func main() {
	app := cli.App("cld", "Uploads images to Cloudinary")

	imagePath := app.StringArg("INPUT", ".", "The input file or directory to process")
	uploadPreset := app.StringOpt("p preset", "", "Cloudinary upload preset")
	uploadFolder := app.StringOpt("f folder", "", "Cloudinary upload folder")
	fileExtensions := app.StringOpt("e extensions", "jpg,jpeg,png,gif,bmp,tiff,webp",
		"Comma-separated list of file extensions to upload")
	apiKey := app.StringOpt("a api-key", "", "API key to select Cloudinary account")

	app.Action = func() {
		cloudinaryURL := os.Getenv(cloudinaryURLEnvVar)
		if cloudinaryURL == "" {
			log.Fatalf("%s environment variable is not set", cloudinaryURLEnvVar)
		}

		urls := strings.Split(cloudinaryURL, ",")
		switch {
		case apiKey != nil && *apiKey != "":
			url, err := getCloudinaryURL(urls, *apiKey)
			if err != nil {
				log.Fatalf("failed to get Cloudinary URL for API key %s: %s", *apiKey, err)
			}
			cloudinaryURL = url
		case len(urls) > 1:
			cloudinaryURL = urls[0]
		}

		cld, err := cloudinary.NewFromURL(cloudinaryURL)
		if err != nil {
			log.Fatalf("failed to create a new Cloudinary client: %s", err)
		}

		if err := process(cld, fileExtensions, uploadPreset, uploadFolder, imagePath); err != nil {
			log.Fatalf("image upload failed: %s", err)
		}
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
}

// getCloudinaryURL check given URLs and returns Cloudinary URL for the matching API key.
// Returns an error if API key is not found.
func getCloudinaryURL(urls []string, apiKey string) (string, error) {
	if apiKey == "" {
		return "", nil
	}

	cloudinaryRegex := regexp.MustCompile(`cloudinary://([^:]+):.*`)

	for _, url := range urls {
		matches := cloudinaryRegex.FindStringSubmatch(url)
		if len(matches) == 2 {
			if matches[1] == apiKey {
				return url, nil
			}
		}
	}

	return "", errors.New("url not found")
}

func process(
	cld *cloudinary.Cloudinary,
	fileExtensions, uploadPreset, uploadFolder, imagePath *string,
) error {
	extList := strings.Split(*fileExtensions, ",")
	for i, ext := range extList {
		extList[i] = strings.ToLower(strings.TrimSpace(ext))
	}

	fileInfo, err := os.Stat(*imagePath)
	if err != nil {
		return fmt.Errorf("failed to stat input path: %w", err)
	}

	ctx := context.Background()
	if !fileInfo.IsDir() {
		if isImage(fileInfo.Name(), extList) {
			return uploadFile(ctx, cld, *imagePath, *uploadPreset, *uploadFolder)
		}

		return fmt.Errorf("input file is not a supported image: %s", *imagePath)
	}

	files, err := os.ReadDir(*imagePath)
	if err != nil {
		return fmt.Errorf("failed to read image dir: %w", err)
	}

	for _, file := range files {
		if file.IsDir() || !isImage(file.Name(), extList) {
			continue
		}

		if err := uploadFile(ctx, cld, filepath.Join(*imagePath, file.Name()),
			*uploadPreset, *uploadFolder); err != nil {
			return err
		}
	}

	return nil
}

func isImage(fileName string, extList []string) bool {
	for _, ext := range extList {
		if strings.HasSuffix(strings.ToLower(fileName), "."+ext) {
			return true
		}
	}
	return false
}

func uploadFile(
	ctx context.Context,
	cld *cloudinary.Cloudinary,
	filePath, preset, folder string,
) error {
	uploadParams := uploader.UploadParams{}
	if preset != "" {
		uploadParams.UploadPreset = preset
	}
	if folder != "" {
		uploadParams.Folder = folder
	}

	resp, err := cld.Upload.Upload(ctx, filePath, uploadParams)
	if err != nil {
		return fmt.Errorf("failed to upload image %s: %w", filePath, err)
	}

	fmt.Printf("image uploaded: %s\n", resp.SecureURL)
	return nil
}
