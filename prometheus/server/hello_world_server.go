package main

import (
	"context"
	"net"
	"net/http"

	"github.com/golang/glog"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":50051"
)

// 使用链式拦截器
// https://segmentfault.com/a/1190000016601823

// PrometheusServer 封装prometheus服务
type PrometheusServer struct {
	Registry                *prometheus.Registry
	GrpcMetrics             *grpc_prometheus.ServerMetrics
	CustomizedCounterMetric *prometheus.CounterVec
}

// CreatePrometheusServer 工厂方法
func CreatePrometheusServer(grpcServer *grpc.Server) (obj *PrometheusServer) {
	registry := prometheus.NewRegistry()
	grpcMetrics := grpc_prometheus.NewServerMetrics()
	grpcMetrics.InitializeMetrics(grpcServer)
	customizedCounterMetric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "demo_server_say_hello_method_handle_count",
			Help: "Total number of RPCs handled on the server.",
		}, []string{"name"})

	registry.MustRegister(grpcMetrics, customizedCounterMetric)
	return &PrometheusServer{
		Registry:                registry,
		GrpcMetrics:             grpcMetrics,
		CustomizedCounterMetric: customizedCounterMetric,
	}
}

// Run .
func (s *PrometheusServer) Run() {
	// 创建prometheus的HTTP server
	httpServer := &http.Server{
		Handler: promhttp.HandlerFor(s.Registry,
			promhttp.HandlerOpts{}), Addr: "0.0.0.0:50052"}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			glog.Fatalf("Unable to start a http server. err=%v", err)
		}
	}()
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (
	*pb.HelloReply, error) {

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// LoggingInterceptor 实现一元拦截器
func LoggingInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
	interface{}, error) {
	glog.V(8).Infof("gRPC method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	glog.V(8).Infof("gRPC method: %s, %v", info.FullMethod, resp)
	return resp, err
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			LoggingInterceptor,
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc_middleware.WithStreamServerChain(
			grpc_prometheus.StreamServerInterceptor,
		),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(s, &server{})

	glog.V(8).Infof("Starting prometheus server ...")
	prometheusServer := CreatePrometheusServer(s)
	prometheusServer.Run()

	glog.V(8).Infof("Starting GRPC server ...")
	if err := s.Serve(lis); err != nil {
		glog.Fatalf("failed to serve: %v", err)
	}
}
