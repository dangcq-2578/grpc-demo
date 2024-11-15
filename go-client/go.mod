module github.com/dangcq-2578/grpc-demo

go 1.18

require (
	github.com/framgia/apollo-worker/github/rpc v0.0.0
	github.com/framgia/apollo-worker/githublanguage/rpc v0.0.0
	github.com/framgia/apollo-worker/synchistory/rpc v0.0.0
	github.com/lib/pq v1.10.9
	github.com/spaolacci/murmur3 v1.1.0
	github.com/zeromicro/go-queue v1.1.8
	github.com/zeromicro/go-zero v1.5.4
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)

replace github.com/framgia/apollo-worker/githublanguage/rpc v0.0.0 => ../../githublanguage/rpc

replace github.com/framgia/apollo-worker/synchistory/rpc v0.0.0 => ../../synchistory/rpc

replace github.com/framgia/apollo-worker/github/rpc v0.0.0 => ../../github/rpc

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
