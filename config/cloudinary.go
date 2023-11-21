package config

import (
	"os"

	"github.com/cloudinary/cloudinary-go"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cloudinaryAPIKey := os.Getenv("CLOUDINARY_API_KEY")
	cloudinaryAPISecret := os.Getenv("CLOUDINARY_API_SECRET")
	cloudinaryCloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")

	cloud, err := cloudinary.NewFromParams(cloudinaryCloudName, cloudinaryAPIKey, cloudinaryAPISecret)
	if err != nil {
		return nil, err
	}

	return cloud, nil
}
