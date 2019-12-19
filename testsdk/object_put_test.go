package testsdk

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
	."github.com/smartystreets/goconvey/convey"
	"os"
	"strings"
	"testing"
)

func TestPutObject(t*testing.T){
	Convey("Put Object", t, func() {
		DeleteTestBucketAndObject()
		defer DeleteTestBucketAndObject()
		sc := s3lib.NewS3(endpoint, accessKey, secretKey)
		// Create a bucket
		err := sc.MakeBucket(bucketName)
		if err != nil {
			HandleError(err)
		}

		// 1. Put a string object
		err = sc.PutObject(bucketName, objectKey, strings.NewReader("NewBucketAndObjectSample"))
		if err != nil {
			HandleError(err)
		}

		err = sc.DeleteObject(bucketName, objectKey)
		if err != nil {
			HandleError(err)
		}

		// 2. Put a file
		f, err := os.Open(localFilePath)
		defer f.Close()
		if err != nil {
			HandleError(err)
		}
		err = sc.PutObject(bucketName, objectKey, f)
		if err != nil {
			HandleError(err)
		}

		var keys []string
		keys, _, _, err = sc.ListObjects(bucketName, objectKey, "", 1000)
		if err != nil {
			HandleError(err)
		}
		var filtered []string
		for _, k := range keys {
			if k == objectKey {
				filtered = append(filtered, k)
			}
		}
		So(len(filtered), ShouldEqual, 1)

		out, err := sc.GetObject(bucketName, objectKey)
		if err != nil {
			HandleError(err)
		}

		out.Close()

		fmt.Printf("PutObjectSample Run Success !\n\n")
	})
}