package cloud_resources

import (
	"github.com/dc-lab/sky/cloud_manager/server/db"
	"github.com/dc-lab/sky/cloud_manager/server/entity"
)

type Manager struct {
	crDao *db.PsqlCloudResourcesDao
}

func NewManager(crDao *db.PsqlCloudResourcesDao) *Manager {
	return &Manager{
		crDao: crDao,
	}
}

func (d *Manager) GetUserResources(userId string) ([]entity.CloudResource, error) {
	resourcesDtoList, err := d.crDao.GetUserResources(userId)
	if err != nil {
		return nil, err
	}

	resourcesList := make([]entity.CloudResource, 0)
	for _, resourceDto := range resourcesDtoList {
		resource := entity.CloudResourceFromDto(resourceDto)
		resourcesList = append(resourcesList, resource)
	}

	return resourcesList, nil
}
