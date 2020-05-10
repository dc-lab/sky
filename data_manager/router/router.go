package routers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/swaggo/http-swagger"

	_ "data_manager/docs"
	"data_manager/handlers"
)

func MakeRouter(srv handlers.FilesService) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/files", func(r chi.Router) {
			r.Post("/", srv.CreateFile)
			r.Route("/{fileId:[0-9a-f\\-]+}", func(r chi.Router) {
				r.Use(srv.FileCtx)
				r.Get("/", srv.GetFile)
				r.Get("/data", srv.DownloadFileData)
				r.Post("/data", srv.UploadFileData)
			})
			r.Route("/outputs/{taskId:[0-9a-f\\-]+}/{path:*}", func(r chi.Router) {
				r.Use(srv.FileCtx)
				r.Post("/", srv.GetFile)
			})
		})
	})

	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	return r
}
