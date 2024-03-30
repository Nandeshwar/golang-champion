package router

import (
	"context"
	"fmt"
	"io"
	pb "myprj/pkg/proto"
	"net"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"

	jwt "github.com/dgrijalva/jwt-go"
)

type LoginServer struct {
	pb.LoginServer
}

func (l *LoginServer) Token(ctx context.Context, req *pb.GetTokenReq) (*pb.GetTokenResponse, error) {
	fmt.Println("login page")

	jm := NewJwtManager("myscret", time.Duration(5)*time.Minute)

	var token string
	var err error
	if req.GetUsername() == "nks" && req.GetPassword() == "nks" {
		token, err = jm.Generate(req.GetUsername(), "admin")
		if err != nil {
			return nil, err
		}
	}

	response := &pb.GetTokenResponse{
		Token: token,
	}

	return response, nil
}

func (l *LoginServer) mustEmbedUnimplementedLoginServer() {

}

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

	l := &LoginServer{}

	pb.RegisterBookServer(g, s)
	pb.RegisterLoginServer(g, l)

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

	if strings.Contains(strings.ToLower(info.FullMethod), "token") {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]

	jm := NewJwtManager("myscret", time.Duration(5)*time.Minute)
	_, err1 := jm.Verify(accessToken)
	if err1 != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is invalid")
	}

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

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJwtManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey: secretKey, tokenDuration: tokenDuration}
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username`
	Role     string `json:"role"`
}

func (j *JWTManager) Generate(username string, role string) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.tokenDuration).Unix(),
		},
		Username: username,
		Role:     role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return []byte(j.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token=%s", err.Error())
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token cliams")
	}

	return claims, nil
}
