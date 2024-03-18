package router

import (
	"context"
	"fmt"
	"io"
	pb "myprj/pkg/proto"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	pb.BookServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ServeGrpc() {
	listner, err := net.Listen("tcp", ":"+strconv.Itoa(9901))
	if err != nil {
		fmt.Println("error=", err.Error())
	}

	g := grpc.NewServer(
		grpc.UnaryInterceptor(s.unaryInterceptor),
		grpc.StreamInterceptor(s.streamInterceptor),
	)

	pb.RegisterBookServer(g, s)
	reflection.Register(g)
	if err := g.Serve(listner); err != nil {
		fmt.Println("GRPC server error", err.Error())
	}
}

func (s *Server) ServeRestApi() {
	conn, err := grpc.DialContext(
		context.Background(),
		":9901",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("error=", err.Error())
	}

	grpcMux := runtime.NewServeMux()
	err = pb.RegisterBookHandler(context.Background(), grpcMux, conn)
	if err != nil {
		fmt.Println("error=", err.Error())
		return
	}

	restAddress := ":9902"
	gwServer := &http.Server{Addr: restAddress, Handler: grpcMux}
	if err := gwServer.ListenAndServe(); err != nil {
		fmt.Println("error=", err.Error())
		return
	}

}

func (s *Server) unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	fmt.Println("unary interceptor called")

	return handler(ctx, req)
}

func (s *Server) streamInterceptor(
	req interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	fmt.Println("stream interceptor called")
	return handler(req, stream)
}

func (s *Server) GetBookInfo(c context.Context, req *pb.GetBookInfoReq) (*pb.GetBookInfoResponse, error) {
	res := &pb.GetBookInfoResponse{
		Name:      req.GetName(),
		Auther:    "Ram",
		Publisher: "Shyam",
	}
	return res, nil
}

func (s *Server) GetBookInfoBidirectional(stream pb.Book_GetBookInfoBidirectionalServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		var name string
		if req != nil {
			name = req.GetName()
			fmt.Println("Got request from client. name=", name)
		}

		res := &pb.GetBookInfoResponse{
			Name:      name,
			Auther:    "nks",
			Publisher: "Gopal",
		}

		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) mustEmbedUnimplementedBookServer() {

}
