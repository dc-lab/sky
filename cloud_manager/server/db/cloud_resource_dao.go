package db

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"
)

type PsqlCloudResourcesDao struct {
	dbClient     *Client
	userTimeout time.Duration
	writeTimeout time.Duration
	readTimeout  time.Duration
}

func NewPsqlCloudResourcesDao(dbClient *Client) *PsqlCloudResourcesDao {
	return &PsqlCloudResourcesDao{
		dbClient:     dbClient,
		userTimeout: defaultOfflineTimeout,
		writeTimeout: defaultRuntimeWriteTimeout,
		readTimeout:  defaultRuntimeReadTimeout,
	}
}

func (d *PsqlCloudResourcesDao) CreateResource(resource *CloudResource) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.writeTimeout)
	defer cancel()

	err := d.dbClient.Conn.QueryRowContext(
		ctx,
		`INSERT INTO cloud_resources (
			factory_id,
			token,
			display_name,
			cpu_limit_cores,
			memory_limit_bytes,
			disk_limit_bytes,
			cpu_guarantee_cores,
			memory_guarantee_bytes,
			disk_guarantee_bytes
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`,
		resource.FactoryId,
		resource.Token,
		resource.DisplayName,
		resource.CpuLimitCores,
		resource.MemoryLimitBytes,
		resource.DiskLimitBytes,
		resource.CpuGuaranteeCores,
		resource.MemoryGuaranteeBytes,
		resource.DiskGuaranteeBytes,
	).Scan(&resource.Id)
	return err
}

func (d *PsqlCloudResourcesDao) UpdateResourceStatus(resource *CloudResource) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.writeTimeout)
	defer cancel()

	_, err := d.dbClient.Conn.ExecContext(
		ctx,
		`UPDATE cloud_resources
		SET
			status=$2
		WHERE id=$1`,
		resource.Id,
		resource.Status,
	)
	return err
}

func (d *PsqlCloudResourcesDao) GetUserResources(userId string) ([]CloudResource, error) {
	ctx, cancel := context.WithTimeout(context.Background(), d.userTimeout)
	defer cancel()

	rows, err := d.dbClient.Conn.QueryContext(
		ctx,
		`SELECT
			cloud_resources.id,
			cloud_resources.factory_id,
			cloud_resources.token,
			cloud_resources.display_name,
			cloud_resources.status,
			cloud_resources.cpu_limit_cores,
			cloud_resources.memory_limit_bytes,
			cloud_resources.disk_limit_bytes,
			cloud_resources.cpu_guarantee_cores,
			cloud_resources.memory_guarantee_bytes,
			cloud_resources.disk_guarantee_bytes
		FROM cloud_resource_factories
		JOIN cloud_resources
		ON cloud_resource_factories.owner_id=$1 AND cloud_resources.factory_id = cloud_resource_factories.id`,
		userId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	allResources := make([]CloudResource, 0)

	for rows.Next() {
		var resource CloudResource
		err := rows.Scan(
			&resource.Id,
			&resource.FactoryId,
			&resource.Token,
			&resource.DisplayName,
			&resource.Status,
			&resource.CpuLimitCores,
			&resource.MemoryLimitBytes,
			&resource.DiskLimitBytes,
			&resource.CpuGuaranteeCores,
			&resource.MemoryGuaranteeBytes,
			&resource.DiskGuaranteeBytes,
		)
		if err != nil {
			log.Tracef("Failed to get cloud resource row. Skipping... %v", err)
			continue
		}
		allResources = append(allResources, resource)
	}

	return allResources, nil
}

func (d *PsqlCloudResourcesDao) GetResourceFull(resourceId string) (*CloudResource, error) {
	ctx, cancel := context.WithTimeout(context.Background(), d.readTimeout)
	defer cancel()

	var resource CloudResource
	err := d.dbClient.Conn.QueryRowContext(
		ctx,
		`SELECT
			id,
			factory_id,
			token,
			display_name,
			status,
			cpu_limit_cores,
			memory_limit_bytes,
			disk_limit_bytes,
			cpu_guarantee_cores,
			memory_guarantee_bytes,
			disk_guarantee_bytes
		FROM cloud_resources
		WHERE id=$1`,
		resourceId,
	).Scan(
		&resource.Id,
		&resource.FactoryId,
		&resource.Token,
		&resource.DisplayName,
		&resource.Status,
		&resource.CpuLimitCores,
		&resource.MemoryLimitBytes,
		&resource.DiskLimitBytes,
		&resource.CpuGuaranteeCores,
		&resource.MemoryGuaranteeBytes,
		&resource.DiskGuaranteeBytes,
	)
	return &resource, err
}

func (d *PsqlCloudResourcesDao) DeleteResource(resourceId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.writeTimeout)
	defer cancel()

	_, err := d.dbClient.Conn.ExecContext(
		ctx,
		`DELETE FROM cloud_resources
		WHERE id=$1`,
		resourceId,
	)
	return err
}
