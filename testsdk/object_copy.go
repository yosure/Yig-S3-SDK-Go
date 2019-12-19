package testsdk

import "fmt"

func CopyObjectSample() {
	DeleteTestBucketAndObject()
	defer DeleteTestBucketAndObject()

	fmt.Printf("CopyObjectSample Run Success !\n\n")
}
