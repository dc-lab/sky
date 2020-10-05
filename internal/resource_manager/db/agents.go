package db

import (
	pb "github.com/dc-lab/sky/api/proto"
	"github.com/dc-lab/sky/internal/resource_manager/app"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type HardwareData struct {
	CoresCount  float64
	MemoryBytes uint64
	DiskBytes   uint64
}

type AgentConnection struct {
	Status        string
	MessageQueue  chan pb.ToAgentMessage
	TotalHardware HardwareData
	FreeHardware  HardwareData
	LastUpdate    time.Time
}

type AgentMap struct {
	mu         sync.Mutex
	agents     map[string]AgentConnection
	lastUpdate time.Time
}

func NewAgentMap() *AgentMap {
	return &AgentMap{
		agents: make(map[string]AgentConnection),
	}
}

func (am *AgentMap) AddAgent(resourceId string) {
	am.mu.Lock()
	defer am.mu.Unlock()
	am.agents[resourceId] = AgentConnection{Status: "online", MessageQueue: make(chan pb.ToAgentMessage)}
}

func (am *AgentMap) RemoveAgent(resourceId string) {
	am.mu.Lock()
	defer am.mu.Unlock()
	delete(am.agents, resourceId)
}

func (am *AgentMap) AddMessage(resourceId string, message *pb.ToAgentMessage) error {
	am.mu.Lock()
	defer am.mu.Unlock()
	if connection, ok := am.agents[resourceId]; ok {
		connection.MessageQueue <- *message
		return nil
	}
	return &app.ResourceNotFound{}
}

func (am *AgentMap) GetMessage(resourceId string) *pb.ToAgentMessage {
	am.mu.Lock()
	defer am.mu.Unlock()
	var message *pb.ToAgentMessage = nil
	if connection, ok := am.agents[resourceId]; ok {
		var x pb.ToAgentMessage
		select {
		case x = <-connection.MessageQueue:
			log.Printf("Got one message from channel for resource %s\n", resourceId)
			message = &x
		default:
		}
	}
	return message
}

func (am *AgentMap) GetLastUpdate(resourceId string) *time.Time {
	am.mu.Lock()
	defer am.mu.Unlock()
	if connection, ok := am.agents[resourceId]; ok {
		return &connection.LastUpdate
	}
	return nil
}

func (am *AgentMap) GetResourceStatus(resourceId string) string {
	am.mu.Lock()
	defer am.mu.Unlock()
	if connection, ok := am.agents[resourceId]; ok {
		if time.Since(connection.LastUpdate).Seconds() < 10 {
			return "online"
		}
	}
	return "offline"
}

func (am *AgentMap) AddHardwareData(resourceId string, total, free *pb.HardwareData) error {
	am.mu.Lock()
	defer am.mu.Unlock()
	if connection, ok := am.agents[resourceId]; ok {
		connection.LastUpdate = time.Now()
		connection.TotalHardware.CoresCount = total.GetCoresCount()
		connection.TotalHardware.MemoryBytes = total.GetMemoryBytes()
		connection.TotalHardware.DiskBytes = total.GetMemoryBytes()

		connection.FreeHardware.CoresCount = free.GetCoresCount()
		connection.FreeHardware.MemoryBytes = free.GetMemoryBytes()
		connection.FreeHardware.DiskBytes = free.GetMemoryBytes()
		return nil
	}
	return &app.ResourceNotFound{}
}

func makeProtoHardwareData(data *HardwareData) *pb.HardwareData {
	return &pb.HardwareData{
		CoresCount:  data.CoresCount,
		MemoryBytes: data.MemoryBytes,
		DiskBytes:   data.DiskBytes,
	}
}

func (am *AgentMap) GetHardwareData(resourceId string) (*pb.HardwareData, *pb.HardwareData) {
	am.mu.Lock()
	defer am.mu.Unlock()

	if connection, ok := am.agents[resourceId]; ok {
		return makeProtoHardwareData(&connection.TotalHardware), makeProtoHardwareData(&connection.FreeHardware)
	}

	return nil, nil
}

var ConnectedAgents = NewAgentMap()
