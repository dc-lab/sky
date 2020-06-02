package http_handles

import (
	"github.com/dc-lab/sky/cloud_manager/server/cloud_resources"
	"github.com/dc-lab/sky/cloud_manager/server/credentials"
	"github.com/dc-lab/sky/cloud_manager/server/db"
	"net/http"
)

type DataProvider struct {
	credsManager *credentials.Manager
	cloudResourcesManager *cloud_resources.Manager
}

func NewStorageProvider(dbClient *db.Client) *DataProvider {
	return &DataProvider{
		credsManager: credentials.NewManager(dbClient),
		cloudResourcesManager: cloud_resources.NewManager(dbClient),
	}
}

func (p *DataProvider) BindHandler(handler func (http.ResponseWriter, *http.Request, *DataProvider)) func (http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, p)
	}
}
