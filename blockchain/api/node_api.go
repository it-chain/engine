package api

import (
	"github.com/it-chain/midgard"
	"github.com/it-chain/it-chain-Engine/blockchain"
)

type NodeApi struct {
	eventRepository midgard.Repository
	publisherId string
}

func NewNodeApi(eventRepository midgard.Repository, publisherId string) NodeApi {
	return NodeApi{
		publisherId: publisherId,
		eventRepository: eventRepository,
	}
}

/// todo
func (n *NodeApi) AddNode(node blockchain.Node) error {
	return nil
}

/// todo
func (n *NodeApi) DeleteNode(id blockchain.NodeId) error {
	return nil
}

