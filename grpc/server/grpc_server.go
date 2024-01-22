package main

// 1.需要监听
// 2.需要实例化gRPC服务端
// 3.在gRPC商注册微服务
// 4.启动服务端
import (
   "context"
   "fmt"
   pb "go.study.src/grpc/proto"  //注意这个路径；这里pb是别名，作用：当有多个引用且包名一样时，这样进行区分调用的是哪个引用的包
   "google.golang.org/grpc"
   "net"
)

// 定义空接口
type UserInfoService struct{
   pb.UserInfoServiceServer
}
//var u = UserInfoService{}

// 实现方法
// 第一个参数是上下文参数，所有接口默认都要必填
// 第二个参数是我们定义的UserRequest消息
// 返回值是我们定义的UserReply消息，error返回值也是必须的。
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
   // 通过用户名查询用户信息
   name := req.Name
   // 数据里查用户信息
   if name == "zs" {
      resp = &pb.UserResponse{
         Id:    1,
         Name:  name,
         Age:   22,
         Hobby: []string{"Sing", "Run"},
      }
   }
   return
}

func main() {
   // 地址
   addr := "127.0.0.1:4363"
   // 1.监听
   listener, err := net.Listen("tcp", addr)
   if err != nil {
      fmt.Printf("监听异常:%s\n", err)
   }
   fmt.Printf("监听端口：%s\n", addr)
   // 2.实例化gRPC
   s := grpc.NewServer()
   // 3.在gRPC上注册微服务
   //pb.RegisterUserInfoServiceServer(s, &u)
   pb.RegisterUserInfoServiceServer(s, &UserInfoService{})
   // 4.启动服务端
   s.Serve(listener)
}