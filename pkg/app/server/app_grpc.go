package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/BetaLixT/gosearch/pkg/app/server/common"
	domcom "github.com/BetaLixT/gosearch/pkg/domain/common"

	"github.com/betalixt/gorr"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	servicePrefix = domcom.ServiceName + ":"
)

func (a *app) startGRPC(portStr string) {
	opts := []grpc.ServerOption{}
	opts = append(
		opts,
		grpc.UnaryInterceptor(func(
			c context.Context,
			req interface{},
			info *grpc.UnaryServerInfo,
			handler grpc.UnaryHandler,
		) (resp interface{}, err error) {
			start := time.Now()

			agent := ""
			t := grpc.ServerTransportStreamFromContext(c)
			path := t.Method()

			p, _ := peer.FromContext(c)
			ip := p.Addr.String()

			md, ok := metadata.FromIncomingContext(c)
			if !ok {
				return nil, fmt.Errorf("empty context")
			}

			temp := md["traceparent"]
			traceparent := ""
			if len(temp) > 0 {
				traceparent = temp[0]
			}
			temp = md["user-agent"]
			if len(temp) > 0 {
				agent = temp[0]
			}

			ctx := a.ctxf.Create(traceparent)
			resp, err = handler(ctx, req)
			end := time.Now()
			statusCode := 200
			if err != nil {
				if gerr, ok := err.(*gorr.Error); ok {
					statusCode = gerr.StatusCode
					err = status.Error(
						mapGRPCCode(statusCode),
						servicePrefix+strconv.Itoa(gerr.ErrorCode.Code),
					)
				} else {
					statusCode = 500
					err = status.Error(mapGRPCCode(statusCode), servicePrefix+"10000")
				}
			}

			a.traceRequest(
				ctx,
				common.GRPCLable,
				path,
				"",
				agent,
				ip,
				statusCode,
				0,
				start,
				end,
				common.GRPCLable,
			)
			return
		}),
	)

	_, err := os.Stat(common.CertKeyLocation)
	if err == nil {
		_, err = os.Stat(common.CertPEMLocation)
	}

	if err == nil {
		a.lgr.Info("found certificates for grpc server")
		creds, err := loadTLSCredentials()
		if err != nil {
			a.lgr.Error("failed to load tls credentials", zap.Error(err))
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	a.registerGRPCHandlers(s)
	a.registerCloser(s.GracefulStop)
	reflection.Register(s)

	port, err := strconv.Atoi(portStr)
	if err != nil {
		a.lgr.Warn(
			"unable to parse provided port, setting port to default",
			zap.String("portConfig", portStr),
		)
		port = common.GRPCDefaultPort
	}
	if port < 0 {
		a.lgr.Warn(
			"port was configured was invalid, setting port to default",
		)
		port = common.GRPCDefaultPort
	}

	a.lgr.Info("grpc listening", zap.Int("port", port))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(common.CertPEMLocation, common.CertKeyLocation)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func mapGRPCCode(statusCode int) codes.Code {
	switch statusCode {
	case 400:
		return codes.Internal
	case 401:
		return codes.Unauthenticated
	case 403:
		return codes.PermissionDenied
	case 404:
		return codes.NotFound
	case 429, 502, 503, 504:
		return codes.Unavailable
	default:
		return codes.Unknown
	}
}
