package gateway

import (
	"context"
	"io"
	"net/http"

	gw "github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/markbates/pkger"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	httpSwagger "github.com/swaggo/http-swagger"

	pb "github.com/dc-lab/sky/api/proto"
	lg "github.com/pressly/lg"
)

type App struct {
	config *Config
}

func NewApp(config *Config) (*App, error) {
	return &App{config}, nil
}

func (a *App) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcHandler, err := a.makeGrpcGatewayMux(ctx)
	if err != nil {
		return err
	}

	mux := a.makeHTTPMux(grpcHandler)

	log.Infoln("Starting grpc-gateway server at", a.config.BindAddress)
	return http.ListenAndServe(a.config.BindAddress, mux)
}

func (a *App) makeGrpcGatewayMux(ctx context.Context) (http.Handler, error) {
	mux := gw.NewServeMux(
		gw_runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			switch key {
			case "User-Id":
				return key, true
			default:
				return gw_runtime.DefaultHeaderMatcher(key)
			}
		}),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}

	var err error
	err = pb.RegisterDataManagerHandlerFromEndpoint(ctx, mux, a.config.DataManagerAddress, opts)
	if err != nil {
		return nil, err
	}

	err = pb.RegisterJobManagerHandlerFromEndpoint(ctx, mux, a.config.JobManagerAddress, opts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func (a *App) makeHTTPMux(grpcHandler http.Handler) http.Handler {
	lg.RedirectStdlogOutput(log.StandardLogger())
	lg.DefaultLogger = log.StandardLogger()

	r := chi.NewMux()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(lg.RequestLogger(log.StandardLogger()))

	log.Infoln("Attached swagger ui at /swagger")

	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})

	r.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		f, err := pkger.Open("/docs/swagger.json")
		if err != nil {
			log.WithError(err).Errorln("Failed to get swagger.json")
			return
		}
		defer f.Close()

		io.Copy(w, f)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	r.Mount("/", grpcHandler)

	return r
}
