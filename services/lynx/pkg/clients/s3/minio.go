package s3

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io"
	"strings"

	"github.com/minio/minio-go/v6"

	"lynx/pkg/config"
	"lynx/pkg/entities"
)


type MinioS3Client struct {
	modelsBucket string

	client *minio.Client
	config config.S3ClientConfig
}

func NewMinioS3Client(cfg config.S3ClientConfig) (*MinioS3Client, error) {
	client, err := minio.NewWithRegion(cfg.Target, cfg.AccessKey, cfg.SecretKey, cfg.Secure, cfg.Region)
	if err != nil {
		return nil, err
	}
	return &MinioS3Client{
		client: client,
		config: cfg,
	}, nil
}

func (c *MinioS3Client) UploadModelData(model entities.Model, data io.Reader) (path string, err error) {
	_, err = c.client.PutObject(c.config.ModelsBucket, model.ID.String(), data, -1, minio.PutObjectOptions{})
	if err != nil {
		return
	}
	path = c.makeModelS3URL(model)
	return
}

func (c *MinioS3Client) UploadFunctionData(fn entities.Function, data io.Reader) (path string, err error) {
	objectName := fmt.Sprintf("%s-%s", fn.ID.String(), uuid.NewV4().String())
	_, err = c.client.PutObject(c.config.FunctionsBucket, objectName, data, -1, minio.PutObjectOptions{})
	if err != nil {
		return
	}
	path = c.makeFunctionS3URL(objectName)
	return
}

func (c *MinioS3Client) makeModelS3URL(model entities.Model) string {
	return strings.Join([]string{c.config.ModelsBucket, model.ID.String()}, "/")
}

func (c *MinioS3Client) makeFunctionS3URL(fn string) string {
	return strings.Join([]string{c.config.FunctionsBucket, fn}, "/")
}