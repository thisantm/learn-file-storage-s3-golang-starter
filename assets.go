package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}
	return nil
}

func createAssetPath(mediaType string) (string, error) {
	assetPath := make([]byte, 32)
	_, err := rand.Read(assetPath)
	if err != nil {
		return "", err
	}

	assetPathEncoded := base64.RawURLEncoding.EncodeToString(assetPath)
	fileType := mediaTypeToFileType(mediaType)

	return fmt.Sprintf("%s.%s", assetPathEncoded, fileType), nil
}

func mediaTypeToFileType(mediaType string) string {
	mediaTypeSplit := strings.Split(mediaType, "/")
	if len(mediaTypeSplit) != 2 {
		return "bin"
	}
	return mediaTypeSplit[1]
}
