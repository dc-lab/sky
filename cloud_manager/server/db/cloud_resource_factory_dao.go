package db

import (
	"context"
	"time"
)

type PsqlCloudResourceFactoriesDao struct {
	dbClient     *Client
	userTimeout time.Duration
	writeTimeout time.Duration
	readTimeout  time.Duration
}

func NewPsqlCloudResourceFactoriesDao(dbClient *Client) *PsqlCloudResourceFactoriesDao {
	return &PsqlCloudResourceFactoriesDao{
		dbClient:     dbClient,
		userTimeout: defaultOfflineTimeout,
		writeTimeout: defaultRuntimeWriteTimeout,
		readTimeout:  defaultRuntimeReadTimeout,
	}
}

func (d *PsqlCloudResourceFactoriesDao) CreateFactory(factory *CloudResourceFactory) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.userTimeout)
	defer cancel()

	err := d.dbClient.Conn.QueryRowContext(
		ctx,
		`INSERT INTO cloud_resource_factories (
		    owner_id,
			token,
			display_name,
			provider,
			agent_docker_version,
			type,
			cpu_inst_limit_cores,
			memory_inst_limit_bytes,
			disk_inst_limit_bytes,
			cpu_fact_limit_cores,
			memory_fact_limit_bytes,
			disk_fact_limit_bytes,
			cpu_used_cores,
			memory_used_bytes,
			disk_used_bytes,
			inst_count_limit,
			inst_count_used,
			aws_cluster,
			aws_vpc,
			aws_tag,
			aws_nets,
			aws_sec_groups
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
		RETURNING id`,
		factory.OwnerId,
		factory.Token,
		factory.DisplayName,
		factory.Provider,
		factory.AgentDockerVersion,
		factory.Type,
		factory.InstanceCpuLimitCores,
		factory.InstanceMemoryLimitBytes,
		factory.InstanceDiskLimitBytes,
		factory.FactoryCpuLimitCores,
		factory.FactoryMemoryLimitBytes,
		factory.FactoryDiskLimitBytes,
		factory.UsedCpuCores,
		factory.UsedMemoryBytes,
		factory.UsedDiskBytes,
		factory.InstanceLimitCount,
		factory.UsedInstanceCount,
		factory.AwsCluster,
		factory.AwsVpc,
		factory.AwsTag,
		factory.AwsNets,
		factory.AwsSecGroups,
	).Scan(&factory.Id)
	return err
}

func (d *PsqlCloudResourceFactoriesDao) UpdateFactoryFull(factory *CloudResourceFactory) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.userTimeout)
	defer cancel()

	_, err := d.dbClient.Conn.ExecContext(
		ctx,
		`UPDATE cloud_resource_factories
		SET
			owner_id=$2,
			token=$3,
			display_name=$4,
			provider=$5,
			agent_docker_version=$6,
			type=$7,
			cpu_inst_limit_cores=$8,
			memory_inst_limit_bytes=$9,
			disk_inst_limit_bytes=$10,
			cpu_fact_limit_cores=$11,
			memory_fact_limit_bytes=$12,
			disk_fact_limit_bytes=$13,
			cpu_used_cores=$14,
			memory_used_bytes=$15,
			disk_used_bytes=$16,
			inst_count_limit=$17,
			inst_count_used=$18,
			aws_cluster=$19,
			aws_vpc=$20,
			aws_tag=$21,
			aws_nets=$22,
			aws_sec_groups=$23
		WHERE id=$1`,
		factory.Id,
		factory.OwnerId,
		factory.Token,
		factory.DisplayName,
		factory.Provider,
		factory.AgentDockerVersion,
		factory.Type,
		factory.InstanceCpuLimitCores,
		factory.InstanceMemoryLimitBytes,
		factory.InstanceDiskLimitBytes,
		factory.FactoryCpuLimitCores,
		factory.FactoryMemoryLimitBytes,
		factory.FactoryDiskLimitBytes,
		factory.UsedCpuCores,
		factory.UsedMemoryBytes,
		factory.UsedDiskBytes,
		factory.InstanceLimitCount,
		factory.UsedInstanceCount,
		factory.AwsCluster,
		factory.AwsVpc,
		factory.AwsTag,
		factory.AwsNets,
		factory.AwsSecGroups,
	)
	return err
}

func (d *PsqlCloudResourceFactoriesDao) UpdateFactoryUsage(factory *CloudResourceFactory) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.writeTimeout)
	defer cancel()

	_, err := d.dbClient.Conn.ExecContext(
		ctx,
		`UPDATE cloud_resource_factories
		SET
			cpu_used_cores=$2,
			memory_used_bytes=$3,
			disk_used_bytes=$4,
			inst_count_used=$5
		WHERE id=$1`,
		factory.Id,
		factory.UsedCpuCores,
		factory.UsedMemoryBytes,
		factory.UsedDiskBytes,
		factory.UsedInstanceCount,
	)
	return err
}

// TODO: make lightweight select query for runtime
func (d *PsqlCloudResourceFactoriesDao) GetFactoryFull(factoryId string) (*CloudResourceFactory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), d.readTimeout)
	defer cancel()

	var factory CloudResourceFactory
	err := d.dbClient.Conn.QueryRowContext(
		ctx,
		`SELECT
			id,
			owner_id,
			token,
			display_name,
			provider,
			agent_docker_version,
			type,
			cpu_inst_limit_cores,
			memory_inst_limit_bytes,
			disk_inst_limit_bytes,
			cpu_fact_limit_cores,
			memory_fact_limit_bytes,
			disk_fact_limit_bytes,
			cpu_used_cores,
			memory_used_bytes,
			disk_used_bytes,
			inst_count_limit,
			inst_count_used,
			aws_cluster,
			aws_vpc,
			aws_tag,
			aws_nets,
			aws_sec_groups
		FROM cloud_resource_factories
		WHERE id=$1`,
		factoryId,
	).Scan(
		&factory.Id,
		&factory.OwnerId,
		&factory.Token,
		&factory.DisplayName,
		&factory.Provider,
		&factory.AgentDockerVersion,
		&factory.Type,
		&factory.InstanceCpuLimitCores,
		&factory.InstanceMemoryLimitBytes,
		&factory.InstanceDiskLimitBytes,
		&factory.FactoryCpuLimitCores,
		&factory.FactoryMemoryLimitBytes,
		&factory.FactoryDiskLimitBytes,
		&factory.UsedCpuCores,
		&factory.UsedMemoryBytes,
		&factory.UsedDiskBytes,
		&factory.InstanceLimitCount,
		&factory.UsedInstanceCount,
		&factory.AwsCluster,
	)
	return &factory, err
}

// TODO: select query with conditions
func (d *PsqlCloudResourcesDao) DeleteFactory(factoryId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), d.writeTimeout)
	defer cancel()

	_, err := d.dbClient.Conn.ExecContext(
		ctx,
		`DELETE FROM cloud_resource_factories
		WHERE id=$1`,
		factoryId,
	)
	return err
}
