package sample

import "fmt"

func BucketLoggingSample() {
	DeleteTestBucketAndObject()

	// TODO: SetBucketLogging(bucketName, logBucketName, "prefix")

	// TODO: GetBucketLogging(bucketName)

	// TODO: DeleteBucketLogging(bucketName)

	DeleteTestBucketAndObject()
	fmt.Printf("BucketLoggingSample Run Success !\n\n")
}
