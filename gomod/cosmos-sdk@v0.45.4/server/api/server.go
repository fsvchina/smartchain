package api

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gogo/gateway"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/tendermint/tendermint/libs/log"
	tmrpcserver "github.com/tendermint/tendermint/rpc/jsonrpc/server"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/telemetry"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"github.com/cosmos/cosmos-sdk/types/rest"

	
	_ "github.com/cosmos/cosmos-sdk/client/docs/statik"
)


type Server struct {
	Router            *mux.Router
	GRPCGatewayRouter *runtime.ServeMux
	ClientCtx         client.Context

	logger  log.Logger
	metrics *telemetry.Metrics
	
	
	
	
	mtx      sync.Mutex
	listener net.Listener
}






func CustomGRPCHeaderMatcher(key string) (string, bool) {
	switch strings.ToLower(key) {
	case grpctypes.GRPCBlockHeightHeader:
		return grpctypes.GRPCBlockHeightHeader, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func New(clientCtx client.Context, logger log.Logger) *Server {
	
	
	marshalerOption := &gateway.JSONPb{
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
		AnyResolver:  clientCtx.InterfaceRegistry,
	}

	return &Server{
		Router:    mux.NewRouter(),
		ClientCtx: clientCtx,
		logger:    logger,
		GRPCGatewayRouter: runtime.NewServeMux(
			
			runtime.WithMarshalerOption(runtime.MIMEWildcard, marshalerOption),

			
			
			runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),

			
			
			runtime.WithIncomingHeaderMatcher(CustomGRPCHeaderMatcher),
		),
	}
}





func (s *Server) Start(cfg config.Config) error {
	s.mtx.Lock()
	if cfg.Telemetry.Enabled {
		m, err := telemetry.New(cfg.Telemetry)
		if err != nil {
			s.mtx.Unlock()
			return err
		}

		s.metrics = m
		s.registerMetrics()
	}

	tmCfg := tmrpcserver.DefaultConfig()
	tmCfg.MaxOpenConnections = int(cfg.API.MaxOpenConnections)
	tmCfg.ReadTimeout = time.Duration(cfg.API.RPCReadTimeout) * time.Second
	tmCfg.WriteTimeout = time.Duration(cfg.API.RPCWriteTimeout) * time.Second
	tmCfg.MaxBodyBytes = int64(cfg.API.RPCMaxBodyBytes)

	listener, err := tmrpcserver.Listen(cfg.API.Address, tmCfg)
	if err != nil {
		s.mtx.Unlock()
		return err
	}

	s.registerGRPCGatewayRoutes()

	s.listener = listener
	var h http.Handler = s.Router

	if cfg.API.EnableUnsafeCORS {
		allowAllCORS := handlers.CORS(handlers.AllowedHeaders([]string{"Content-Type"}))
		s.mtx.Unlock()
		return tmrpcserver.Serve(s.listener, allowAllCORS(h), s.logger, tmCfg)
	}

	s.logger.Info("starting API server...")
	s.mtx.Unlock()
	return tmrpcserver.Serve(s.listener, s.Router, s.logger, tmCfg)
}


func (s *Server) Close() error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.listener.Close()
}

func (s *Server) registerGRPCGatewayRoutes() {
	s.Router.PathPrefix("/").Handler(s.GRPCGatewayRouter)
}

func (s *Server) registerMetrics() {
	metricsHandler := func(w http.ResponseWriter, r *http.Request) {
		format := strings.TrimSpace(r.FormValue("format"))

		gr, err := s.metrics.Gather(format)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to gather metrics: %s", err))
			return
		}

		w.Header().Set("Content-Type", gr.ContentType)
		_, _ = w.Write(gr.Metrics)
	}

	s.Router.HandleFunc("/metrics", metricsHandler).Methods("GET")
}
