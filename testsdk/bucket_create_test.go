package testsdk

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
	"github.com/journeymidnight/aws-sdk-go/service/s3"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	Convey("Create bucket", t, func() {
		ReadConfig()
		DeleteTestBucketAndObject()
		defer DeleteTestBucketAndObject()
		sc := s3lib.NewS3(endpoint, accessKey, secretKey)

		// Create a bucket
		err := sc.MakeBucket(bucketName)
		if err != nil {
			HandleError(err)
		}

		out, err := sc.ListBuckets()
		if err != nil {
			HandleError(err)
		}
		//Assert bucket create successful
		var filtered []*s3.Bucket
		for _, b := range out {
			if *b.Name == bucketName {
				filtered = append(filtered, b)
			}
		}
		So(len(filtered), ShouldEqual, 1)
		// TODO: Make bucket with ACL

		// Delete the bucket
		err = sc.DeleteBucket(bucketName)
		if err != nil {
			HandleError(err)
		}

		fmt.Printf("CreateBucketSample Run Success!\n\n")
	})
}
