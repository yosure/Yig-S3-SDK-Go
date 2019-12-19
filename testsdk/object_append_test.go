package testsdk

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/s3lib"
	."github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"strings"
	"testing"
)

func TestAppendObject(t*testing.T){
	Convey("Append Object", t, func() {
		DeleteTestBucketAndObject()
		defer DeleteTestBucketAndObject()
		sc := s3lib.NewS3(endpoint, accessKey, secretKey)
		// Create a bucket
		err := sc.MakeBucket(bucketName)
		if err != nil {
			HandleError(err)
		}
		var nextPos int64

		// 1. Append strings to an object
		strs := []string{"yig1", "yig2", "yig3"}
		for _, s := range strs {
			//fmt.Println("Append String:", s)
			nextPos, err = sc.AppendObject(bucketName, objectKey, strings.NewReader(s), nextPos)
			if err != nil {
				HandleError(err)
			}
		}
		out, err := sc.GetObject(bucketName, objectKey)
		b, _ := ioutil.ReadAll(out)
		So(string(b), ShouldEqual, "yig1yig2yig3")
		//fmt.Println("Get appended string:", string(b))
		out.Close()
		// TODO 2. Append files to an object

		// TODO 3. Get Next Append Position

		// TODO 4. Append With ACL And Meta

		fmt.Printf("AppendObjectSample Run Success !\n\n")
	})
}