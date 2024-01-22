module go.study.src

go 1.21.1

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/go-sql-driver/mysql v1.7.1
	github.com/pochard/commons v1.1.2
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
	github.com/cj/go.study v1.0.0      //并不是github官网上的，而是我们本地的
)
replace github.com/cj/go.study => /home/cj/go/src/go.study.src_2 //并不是github官网上的，而是我们本地的,这就是replace的用处，用于替换
require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
)
