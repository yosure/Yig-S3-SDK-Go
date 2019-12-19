package testsdk

import (
	"fmt"
	"os"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
	"io"
)

func GetObjectSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	sc := s3lib.NewS3(endpoint, accessKey, secretKey)
	// Create a bucket
	err := sc.MakeBucket(bucketName)
	if err != nil {
		HandleError(err)
	}

	// Put a file
	f, err := os.Open(localFilePath)
	defer f.Close()
	if err != nil {
		HandleError(err)
	}
	err = sc.PutObject(bucketName, objectKey, f)
	if err != nil {
		HandleError(err)
	}

	// Get the reader
	out, err := sc.GetObject(bucketName, objectKey)
	if err != nil {
		HandleError(err)
	}

	// Download to a file
	f2, err := os.OpenFile("sample/Download.jpeg", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	defer f2.Close()
	if err != nil {
		HandleError(err)
	}
	io.Copy(f2, out)
	out.Close()

	fmt.Printf("GetObjectSample Run Success !\n\n")
}
