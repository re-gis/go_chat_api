package utils

import (
	"chatapp/config"
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadToCloudinary(file multipart.File, filepath string) (string, error) {
	ctx := context.Background()
	cloud, err := config.SetupCloudinary()
	if err != nil {
		return "", err
	}

	uploadParams := uploader.UploadParams{
		PublicID: filepath,
	}

	result, err := cloud.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", err
	}

	imageUrl := result.SecureURL
	return imageUrl, nil
}
