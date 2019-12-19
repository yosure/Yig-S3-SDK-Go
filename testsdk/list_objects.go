package testsdk

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
	."github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestListObjects(t *testing.T) {
	Convey("List Objects", t, func() {
		var keys = []string{
			objectKey + "1",
			objectKey + "2",
			objectKey + "3",
			objectKey + "4",
			objectKey + "/1-1",
			objectKey + "/1-2",
			objectKey + "/2-1",
			objectKey + "/2-2",
		}
		sc := s3lib.NewS3(endpoint, accessKey, secretKey)
		for _, k := range keys {
			sc.DeleteObject(bucketName, k)
		}
		DeleteTestBucketAndObject()


		// Create a bucket
		err := sc.MakeBucket(bucketName)
		if err != nil {
			HandleError(err)
		}

		for _, k := range keys {
			err := sc.PutObject(bucketName, k, strings.NewReader(k))
			if err != nil {
				HandleError(err)
			}
		}

		keys, _, _, err = sc.ListObjects(bucketName, objectKey+"/", "/", 1000)
		if err != nil {
			HandleError(err)
		}

		for _, k := range keys {
			err := sc.DeleteObject(bucketName, k)
			if err != nil {
				HandleError(err)
			}
		}
		DeleteTestBucketAndObject()
		fmt.Printf("ListObjectsSample Run Success !\n\n")
	})
}