package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/dc-lab/sky/internal/data_manager/node/service"
	log "github.com/sirupsen/logrus"
)

func MakeRouter(srv *service.BlobsService) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.StandardLogger(), NoColor: false}))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/files", func(r chi.Router) {
			r.Route("/{file_id:[0-9a-f\\-]+}", func(r chi.Router) {
				r.Get("/data", srv.DownloadBlob)
				r.Post("/data", srv.UploadBlob)
			})
		})
	})

	return r
}
