package minio

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/log"
	"github.com/kubeshop/testkube/pkg/storage"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

var _ storage.Client = (*Client)(nil)

// ErrArtifactsNotFound contains error for not existing artifacts
var ErrArtifactsNotFound = errors.New("Execution doesn't have any artifacts associated with it")

// Client for managing MinIO storage server
type Client struct {
	Endpoint        string
	accessKeyID     string
	secretAccessKey string
	ssl             bool
	location        string
	token           string
	minioclient     *minio.Client
	Log             *zap.SugaredLogger
}

// NewClient returns new MinIO client
func NewClient(endpoint, accessKeyID, secretAccessKey, location, token string, ssl bool) *Client {
	c := &Client{
		location:        location,
		accessKeyID:     accessKeyID,
		secretAccessKey: secretAccessKey,
		token:           token,
		ssl:             ssl,
		Endpoint:        endpoint,
		Log:             log.DefaultLogger,
	}

	return c
}

// Connect connects to MinIO server
func (c *Client) Connect() error {
	mclient, err := minio.New(c.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.accessKeyID, c.secretAccessKey, c.token),
		Secure: c.ssl,
	})
	if err != nil {
		return err
	}
	c.minioclient = mclient
	return err
}

// CreateBucket creates new S3 like bucket
func (c *Client) CreateBucket(bucket string) error {
	ctx := context.Background()
	err := c.minioclient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{Region: c.location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := c.minioclient.BucketExists(ctx, bucket)
		if errBucketExists == nil && exists {
			return fmt.Errorf("bucket %q already exists", bucket)
		} else {
			return err
		}
	}
	return nil
}

// DeleteBucket deletes bucket by name
func (c *Client) DeleteBucket(bucket string, force bool) error {
	return c.minioclient.RemoveBucketWithOptions(context.TODO(), bucket, minio.BucketOptions{ForceDelete: force})
}

// ListBuckets lists available buckets
func (c *Client) ListBuckets() ([]string, error) {
	toReturn := []string{}
	if buckets, err := c.minioclient.ListBuckets(context.TODO()); err != nil {
		return nil, err
	} else {
		for _, bucket := range buckets {
			toReturn = append(toReturn, bucket.Name)
		}
	}
	return toReturn, nil
}

// ListFiles lists available files in given bucket
func (c *Client) ListFiles(bucket string) ([]testkube.Artifact, error) {
	if err := c.Connect(); err != nil {
		return nil, err
	}
	toReturn := []testkube.Artifact{}

	exists, err := c.minioclient.BucketExists(context.TODO(), bucket)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrArtifactsNotFound
	}

	for obj := range c.minioclient.ListObjects(context.TODO(), bucket, minio.ListObjectsOptions{Recursive: true}) {
		if obj.Err != nil {
			return nil, obj.Err
		}
		toReturn = append(toReturn, testkube.Artifact{Name: obj.Key, Size: int32(obj.Size)})
	}

	return toReturn, nil
}

// SaveFile saves file defined by local filePath to S3 bucket
func (c *Client) SaveFile(bucket, filePath string) error {
	if err := c.Connect(); err != nil {
		return err
	}
	object, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("minio saving file (%s) open error: %w", filePath, err)
	}
	defer object.Close()
	objectStat, err := object.Stat()
	if err != nil {
		return fmt.Errorf("minio object stat (file:%s) error: %w", filePath, err)
	}

	fileName := objectStat.Name()

	c.Log.Debugw("saving object in minio", "filePath", filePath, "fileName", fileName, "bucket", bucket, "size", objectStat.Size())
	_, err = c.minioclient.PutObject(context.Background(), bucket, fileName, object, objectStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return fmt.Errorf("minio saving file (%s) put object error: %w", fileName, err)
	}

	return nil
}

// DownloadFile downloads file in bucket
func (c *Client) DownloadFile(bucket, file string) (*minio.Object, error) {
	if err := c.Connect(); err != nil {
		return nil, fmt.Errorf("minio DownloadFile .Connect error: %w", err)
	}

	exists, err := c.minioclient.BucketExists(context.TODO(), bucket)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrArtifactsNotFound
	}

	reader, err := c.minioclient.GetObject(context.Background(), bucket, file, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("minio DownloadFile GetObject error: %w", err)
	}

	_, err = reader.Stat()
	if err != nil {
		return reader, fmt.Errorf("minio Download File Stat error: %w", err)
	}

	return reader, nil
}

// ScrapeArtefacts pushes local files located in directories to given bucket ID
func (c *Client) ScrapeArtefacts(id string, directories ...string) error {
	if err := c.Connect(); err != nil {
		return fmt.Errorf("minio scrape artefacts connection error: %w", err)
	}

	err := c.CreateBucket(id) // create bucket name it by execution ID
	if err != nil {
		return fmt.Errorf("minio failed to create a bucket %s: %w", id, err)
	}

	for _, directory := range directories {

		if _, err := os.Stat(directory); os.IsNotExist(err) {
			c.Log.Debugw("directory %s does not exist, skipping", directory)
			continue
		}

		// if directory exists walk through recursively
		err = filepath.Walk(directory,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return fmt.Errorf("minio path (%s) walk error: %w", path, err)
				}

				if !info.IsDir() {
					err = c.SaveFile(id, path) //The function will detect if there is a subdirectory and store accordingly
					if err != nil {
						return fmt.Errorf("minio save file (%s) error: %w", path, err)
					}
				}
				return nil
			})

		if err != nil {
			return fmt.Errorf("minio walk error: %w", err)
		}
	}
	return nil
}
