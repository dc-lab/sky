package credentials

import (
	"github.com/dc-lab/sky/cloud_manager/server/db"
	"github.com/dc-lab/sky/cloud_manager/server/entity"
	"github.com/google/uuid"
)

type Manager struct {
 	credsDao *db.PsqlCredsDao
}

func NewManager(credsDao *db.PsqlCredsDao) *Manager {
	return &Manager{
		credsDao: credsDao,
	}
}

func (m *Manager) CreateCreds(userId string, credsInput *entity.CreateCredentialsInput) (*entity.Credentials, error) {
	credsDto := db.Credentials{
		OwnerId: userId,
		DisplayName: credsInput.DisplayName,
		Provider: credsInput.Provider,
		AwsAccessKey: credsInput.AwsAccessKey,
		AwsAccessKeyId: credsInput.AwsAccessKeyId,
	}
	if len(credsDto.DisplayName) == 0 {
		credsDto.DisplayName = uuid.New().String()
	}

	err := m.credsDao.Create(&credsDto)
	if err != nil {
		return nil, err
	}

	creds := entity.CredentialsFromDto(credsDto)
	return &creds, nil
}

func (m *Manager) GetAllUserCreds(userId string) ([]entity.Credentials, error) {
	credsDtoList, err := m.credsDao.GetAll(userId)
	if err != nil {
		return nil, err
	}

	credsList := make([]entity.Credentials, 0)
	for _, credsDto := range credsDtoList {
		creds := entity.CredentialsFromDto(credsDto)
		credsList = append(credsList, creds)
	}

	return credsList, nil
}

func (m *Manager) GetUserCreds(userId string, credsId string) (*entity.Credentials, error) {
	credsDto, err := m.credsDao.Get(userId, credsId)
	if err != nil {
		return nil, err
	}

	creds := entity.CredentialsFromDto(*credsDto)
	return &creds, nil
}

func (m *Manager) DeleteUserCreds(userId string, credsId string) error {
	return m.credsDao.Delete(userId, credsId)
}
