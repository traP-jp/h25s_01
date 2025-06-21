package file

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"path/filepath"

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

func (r *RepositoryImpl) UploadImage(reviewID uuid.UUID, contentType string, reader io.Reader) (uuid.UUID, error) {
	fileID := uuid.New()

	// S3のキーを生成 (reviews/{reviewID}/{fileID})
	key := fmt.Sprintf("reviews/%s/%s", reviewID.String(), fileID.String())

	// io.Readerからバイトデータを読み取り
	data, err := io.ReadAll(reader)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to read data: %w", err)
	}

	// S3にアップロード
	_, err = r.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(r.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to upload image to S3: %w", err)
	}

	return fileID, nil
}

func (r *RepositoryImpl) DeleteImage(fileID uuid.UUID) error {
	// ファイルIDからS3オブジェクトを検索して削除
	// reviews/配下の全てのディレクトリでファイルIDに一致するオブジェクトを探す
	prefix := "reviews/"

	listInput := &s3.ListObjectsV2Input{
		Bucket: aws.String(r.bucket),
		Prefix: aws.String(prefix),
	}

	result, err := r.s3Client.ListObjectsV2(context.TODO(), listInput)
	if err != nil {
		return fmt.Errorf("failed to list objects in S3: %w", err)
	}

	var keyToDelete *string
	targetFileID := fileID.String()

	// オブジェクトリストから該当するファイルを探す
	for _, obj := range result.Contents {
		if obj.Key == nil {
			continue
		}

		// ファイル名がファイルIDと一致するかチェック
		fileName := filepath.Base(*obj.Key)

		if fileName == targetFileID {
			keyToDelete = obj.Key

			break
		}
	}

	if keyToDelete == nil {
		return fmt.Errorf("file not found: %s", fileID.String())
	}

	// S3からオブジェクトを削除
	_, err = r.s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    keyToDelete,
	})
	if err != nil {
		return fmt.Errorf("failed to delete object from S3: %w", err)
	}

	return nil
}
