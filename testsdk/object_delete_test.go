package testsdk

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
	."github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestDeleteObject(t*testing.T){
	Convey("Delete Object", t, func() {
		DeleteTestBucketAndObject()
		defer DeleteTestBucketAndObject()
		// Delete single key
		sc := s3lib.NewS3(endpoint, accessKey, secretKey)
		// Create a bucket
		err := sc.MakeBucket(bucketName)
		if err != nil {
			HandleError(err)
		}

		// 1. Delete an object
		err = sc.PutObject(bucketName, objectKey, strings.NewReader("NewBucketAndObjectSample"))
		if err != nil {
			HandleError(err)
		}

		err = sc.DeleteObject(bucketName, objectKey)
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
		So(len(filtered), ShouldEqual, 0)

		// Delete file that not exists will not failed
		err = sc.DeleteObject(bucketName, objectKey)
		if err != nil {
			HandleError(err)
		}

		// 2. Delete multiple objects
		//err = sc.PutObject(bucketName, objectKey+"1", strings.NewReader("NewBucketAndObjectSample"))
		//if err != nil {
		//	HandleError(err)
		//}
		//
		//err = sc.PutObject(bucketName, objectKey+"2", strings.NewReader("NewBucketAndObjectSample"))
		//if err != nil {
		//	HandleError(err)
		//}

		fmt.Printf("DeleteObjectSample Run Success !\n\n")
	})
}