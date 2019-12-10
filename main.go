package main

import (
	"fmt"
	"github.com/journeymidnight/Yig-S3-SDK-Go/sample"
)

func main() {
	// Read S3 config
	sample.ReadConfig()

	sample.MakeBucketSample()
	sample.ListBucketsSample()
	sample.BucketACLSample()
	sample.BucketLifecycleSample()
	sample.BucketRefererSample()
	sample.BucketLoggingSample()
	//sample.BucketCORSSample()
	//
	//sample.PutObjectSample()
	//sample.GetObjectSample()
	sample.ListObjectsSample()
	//sample.DeleteObjectSample()
	//sample.AppendObjectSample()
	//sample.ObjectACLSample()
	//sample.ObjectMetaSample()


	//sample.CopyObjectSample()

	//sample.ArchiveSample()
	//sample.MySample()

	fmt.Println("All samples completed !")
}
