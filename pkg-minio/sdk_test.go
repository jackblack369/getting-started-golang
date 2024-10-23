package pkg_minio

import (
	"context"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"log"
	"path/filepath"
	"testing"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func TestInitClient(t *testing.T) {
	endpoint := "172.20.7.232:19000"
	accessKeyID := "gon6jjZ9qk5jORg0jQdy"
	secretAccessKey := "9KdhttWQiDISuGVYKLHUUP8XA0sRBy9SOHOzaFbN"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up
}

func TestHomeDir(t *testing.T) {
	homeDIR, err := homedir.Dir()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Home directory is %s\n", homeDIR)
}

func TestFileUploader(t *testing.T) {
	ctx := context.Background()
	endpoint := "172.20.7.232:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called testbucket.
	bucketName := "testbucket"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Upload the test file
	// Change the value of filePath if the file is in another location
	objectName := "test-file.txt"
	homeDIR, err := homedir.Dir()

	filePath := filepath.Join(homeDIR, "/Documents/tmp/test-file.txt")
	contentType := "application/octet-stream"

	// Upload the test file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}

func TestFileList(t *testing.T) {
	ctx := context.Background()
	endpoint := "172.20.7.232:19000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called testbucket.
	bucketName := "gateway-jfs"

	// Check to see if we already own this bucket (which happens if you run this twice)
	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
	if errBucketExists == nil && exists {
		log.Printf("We already own %s\n", bucketName)
	} else {
		log.Fatalln(err)
	}

	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		// Prefix: "myprefix",
		Recursive: true,
	})
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object)
	}

}
