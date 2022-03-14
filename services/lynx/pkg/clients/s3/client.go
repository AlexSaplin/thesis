package s3

import (
	"io"

	"lynx/pkg/entities"
)

type S3Client interface {
	UploadModelData(model entities.Model, data io.Reader) (string, error)
	UploadFunctionData(model entities.Function, data io.Reader) (string, error)
}
