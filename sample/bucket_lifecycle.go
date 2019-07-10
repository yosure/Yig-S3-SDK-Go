package sample

import "fmt"

func BucketLifecycleSample() {
	DeleteTestBucketAndObject()

	// TODO: Support LC ID

	DeleteTestBucketAndObject()
	fmt.Printf("BucketLifecycleSample Run Success !\n\n")
}
