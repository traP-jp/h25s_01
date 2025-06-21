package file

import (
	"context"
	"fmt"
	"io"

	"backend/internal/domain/repository"
	"backend/pkg/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type RepositoryImpl struct {
	s3Client *s3.Client
	bucket   string
}

func NewFileRepository() (repository.FileRepository, error) {
	awsConf := config.AWS()

	var cfg aws.Config
	var err error

	// 基本的な設定オプションを準備
	configOptions := []func(*awsConfig.LoadOptions) error{
		awsConfig.WithRegion(awsConf.Region),
	}

	// アクセスキーとシークレットキーが設定されている場合は静的クレデンシャルを使用
	if awsConf.AccessKey != "" && awsConf.SecretKey != "" {
		configOptions = append(configOptions, awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			awsConf.AccessKey,
			awsConf.SecretKey,
			"",
		)))
	}

	cfg, err = awsConfig.LoadDefaultConfig(context.TODO(), configOptions...)

	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	// S3クライアントのオプションを設定
	var s3Options []func(*s3.Options)

	// カスタムエンドポイントが設定されている場合
	if awsConf.Endpoint != "" {
		s3Options = append(s3Options, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(awsConf.Endpoint)
		})
	}

	// Path-style URLを使用する場合
	if awsConf.PathStyle {
		s3Options = append(s3Options, func(o *s3.Options) {
			o.UsePathStyle = true
		})
	}

	client := s3.NewFromConfig(cfg, s3Options...)

	return &RepositoryImpl{
		s3Client: client,
		bucket:   awsConf.BucketName,
	}, nil
}

func (r *RepositoryImpl) UploadImage(ctx context.Context, contentType string, reader io.Reader) (uuid.UUID, error) {
	fileID := uuid.New()

	key := fileID.String()

	// S3にアップロード
	_, err := r.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r.bucket),
		Key:         aws.String(key),
		Body:        reader,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to upload image to S3: %w", err)
	}

	return fileID, nil
}

func (r *RepositoryImpl) DeleteImage(ctx context.Context, fileID uuid.UUID) error { // S3からオブジェクトを削除
	_, err := r.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(fileID.String()),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object from S3: %w", err)
	}

	return nil
}

func (r *RepositoryImpl) GetImage(ctx context.Context, fileID uuid.UUID) (io.ReadCloser, string, error) {
	key := fileID.String()

	// S3からオブジェクトを取得
	result, err := r.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, "", fmt.Errorf("failed to get object from S3: %w", err)
	}

	contentType := "application/octet-stream"
	if result.ContentType != nil {
		contentType = *result.ContentType
	}

	return result.Body, contentType, nil
}
