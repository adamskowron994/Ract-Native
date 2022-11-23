module github.com/joshcarp/grpc-vs-connect

go 1.14

require (
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/rs/cors v1.8.2
	github.com/stretchr/testify v1.8.1
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102 // indirect
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20201109203340-2640f1f9cdfb // indirect
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc v1.30.0 => google.golang.org/grpc v1.29.0
