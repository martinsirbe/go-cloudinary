package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	app.Action = func() {
		cld, err := cloudinary.New()
		if err != nil {
			log.Fatalf("failed to create a new Cloudinary client: %s", err)
		}

		process(cld, fileExtensions, uploadPreset, uploadFolder, imagePath)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
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
