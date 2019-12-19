package testsdk

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	endpoint      string
	accessKey     string
	secretKey     string
	bucketName    string
	objectKey     string
	localFilePath string
)

func ReadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err:=viper.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
	}
	endpoint = viper.GetString("s3.endpoint")
	accessKey = viper.GetString("s3.accessKey")
	secretKey = viper.GetString("s3.secretKey")
	bucketName = viper.GetString("s3.bucketName")
	objectKey = viper.GetString("s3.objectKey")
	localFilePath = viper.GetString("s3.localFilePath")
}
