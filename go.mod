module github.com/journeymidnight/Yig-S3-SDK-Go

go 1.12

require (
	github.com/aws/aws-sdk-go v1.20.17
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/journeymidnight/aws-sdk-go v1.17.5
	github.com/journeymidnight/yig v2.0.0+incompatible
	github.com/spf13/viper v1.4.0
	github.com/ugorji/go v1.1.7 // indirect
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.37.4

	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.1

	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190404164418-38d8ce5564a5
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190409044807-56b785ea58b2
	golang.org/x/image => github.com/golang/image v0.0.0-20190321063152-3fc05d484e9f
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190327163128-167ebed0ec6d
	golang.org/x/net => github.com/golang/net v0.0.0-20190404232315-eb5bcb51f2a3
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190405154228-4b34438f7a67
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190408220357-e5b8258f4918
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.3.2
	google.golang.org/appengine => github.com/golang/appengine v1.5.0

	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190404172233-64821d5d2107
	google.golang.org/grpc => github.com/grpc/grpc-go v1.19.1
)
