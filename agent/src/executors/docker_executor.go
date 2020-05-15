package executors

import (
	"context"
	pb "github.com/dc-lab/sky/api/proto/common"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"io"
	"os"
	"sync/atomic"
)

type DockerExecutor struct {
	Image                    string       // don't change
	ExecutionDir             string       // don't change
	RequirementsShellCommand string       // don't change
	ExecutionShellCommand    string       // don't change
	ContainerID              atomic.Value // string
	DockerClient             *client.Client
}

func (e *DockerExecutor) SetContainerID(containerID string) {
	e.ContainerID.Store(containerID)
}

func (e *DockerExecutor) GetContainerID() string {
	return e.ContainerID.Load().(string)
}

func (e *DockerExecutor) Prepare() {
	ctx := context.Background()
	var err error
	e.DockerClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	reader, err := e.DockerClient.ImagePull(ctx, e.Image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)
	resp, err := e.DockerClient.ContainerCreate(ctx, &container.Config{
		Image:      e.Image,
		Cmd:        []string{"/bin/sh", "-c", e.ExecutionShellCommand},
		WorkingDir: e.ExecutionDir,
	}, &container.HostConfig{
		Mounts: []mount.Mount{{Type: mount.TypeBind, Source: e.ExecutionDir, Target: e.ExecutionDir}},
	}, nil, "")
	if err != nil {
		panic(err)
	}
	e.SetContainerID(resp.ID)
}

func (e *DockerExecutor) Run(
	quiteChannel <-chan struct{},
	beforeExecution func(result *pb.TResult),
	afterExecution func(err error),
) {
	ctx := context.Background()
	containerID := e.GetContainerID()
	if err := e.DockerClient.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := e.DockerClient.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := e.DockerClient.ContainerLogs(ctx, containerID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(os.Stdout, out)
	//_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	if err != nil {
		panic(err)
	}
}
