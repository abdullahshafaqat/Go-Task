package services

import (
	"io"
	"mime/multipart"
)

func UploadedFile(fileHeader *multipart.FileHeader) ([]byte, error) {
	Data, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer Data.Close()

	fileData, err := io.ReadAll(Data)
	if err != nil {
		return nil, err
	}

	return fileData, nil
}
