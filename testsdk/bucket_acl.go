package testsdk

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBucketAcl(t *testing.T) {
	Convey("Bucket acl", t, func() {
		DeleteTestBucketAndObject()
		defer DeleteTestBucketAndObject()

		sc := s3lib.NewS3(endpoint, accessKey, secretKey)
		err := sc.MakeBucket(bucketName)
		if err != nil {
			HandleError(err)
		}

		// TODO : Surport Set Bucket CannedACL 'PublicRead'

		out, err := sc.GetBucketAcl(bucketName)
		if err != nil {
			HandleError(err)
		}
		fmt.Println("Get Bucket ACL:", out)

		fmt.Printf("BucketACLSample Run Success!\n\n")
	})
}