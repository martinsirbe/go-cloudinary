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

func main() {
	app := cli.App("cld", "Uploads images to Cloudinary")

	imagePath := app.StringArg("INPUT", ".", "The input file or directory to process")
	uploadPreset := app.StringOpt("p preset", "", "Cloudinary upload preset")
	uploadFolder := app.StringOpt("f folder", "", "Cloudinary upload folder")
	fileExtensions := app.StringOpt("e extensions", "jpg,jpeg,png,gif,bmp,tiff,webp",
		"Comma-separated list of file extensions to upload")
	apiKey := app.StringOpt("a api-key", "", "API key to select Cloudinary account")

	app.Action = func() {
		cloudinaryURL := os.Getenv("CLOUDINARY_URL")
		if cloudinaryURL == "" {
			log.Fatalf("CLOUDINARY_URL environment variable is not set")
		}

		urls := strings.Split(cloudinaryURL, ",")
		if len(urls) > 1 {
			switch {
			case apiKey != nil && *apiKey != "":
				url, err := getCloudinaryURL(urls, *apiKey)
				if err != nil {
					log.Fatalf("failed to get Cloudinary URL for API key %s: %s", *apiKey, err)
				}
				cloudinaryURL = url
			default:
				cloudinaryURL = urls[0]
			}
		}

		cld, err := cloudinary.NewFromURL(cloudinaryURL)
		if err != nil {
			log.Fatalf("failed to create a new Cloudinary client: %s", err)
		}

		process(cld, fileExtensions, uploadPreset, uploadFolder, imagePath)
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
) {
	extList := strings.Split(*fileExtensions, ",")
	for i, ext := range extList {
		extList[i] = strings.ToLower(strings.TrimSpace(ext))
	}

	fileInfo, err := os.Stat(*imagePath)
	if err != nil {
		log.Fatalf("failed to stat input path: %v", err)
	}

	ctx := context.Background()
	if !fileInfo.IsDir() {
		if isImage(fileInfo.Name(), extList) {
			uploadFile(ctx, cld, *imagePath, *uploadPreset, *uploadFolder)
			return
		}

		log.Fatalf("input file is not a supported image: %s\n", *imagePath)
	}

	files, err := os.ReadDir(*imagePath)
	if err != nil {
		log.Fatalf("failed to read image dir: %v", err)
	}

	for _, file := range files {
		if file.IsDir() || !isImage(file.Name(), extList) {
			continue
		}

		uploadFile(ctx, cld, filepath.Join(*imagePath, file.Name()), *uploadPreset, *uploadFolder)
	}
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
) {
	uploadParams := uploader.UploadParams{}
	if preset != "" {
		uploadParams.UploadPreset = preset
	}
	if folder != "" {
		uploadParams.Folder = folder
	}

	resp, err := cld.Upload.Upload(ctx, filePath, uploadParams)
	if err != nil {
		fmt.Printf("failed to upload image %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("image uploaded: %s\n", resp.SecureURL)
}
