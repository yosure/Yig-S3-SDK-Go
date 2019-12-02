package sample

import "fmt"

func BucketLifecycleSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()
	// TODO: Support LC ID


	fmt.Printf("BucketLifecycleSample Run Success !\n\n")
}
