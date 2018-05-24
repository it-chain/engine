package rabbitmq

import (
	"encoding/json"

	"github.com/it-chain/it-chain-Engine/common"
	"github.com/it-chain/it-chain-Engine/messaging/rabbitmq/event"
	"github.com/it-chain/it-chain-Engine/messaging/rabbitmq/topic"
	"github.com/it-chain/it-chain-Engine/peer/api"
	"github.com/it-chain/it-chain-Engine/peer/domain/model"
	"github.com/it-chain/it-chain-Engine/peer/domain/repository"
	"github.com/it-chain/it-chain-Engine/peer/domain/service"
	"github.com/streadway/amqp"
)

type MessageListener struct {
	leaderSelectionApi *api.LeaderSelection
	peerRepository     *repository.Peer
	peerTable          *service.PeerTable
	messageProducer    *service.MessageProducer
}

func NewMessageListener(leaderSelectionApi *api.LeaderSelection, repository *repository.Peer, table *service.PeerTable, producer *service.MessageProducer) *MessageListener {

	return &MessageListener{
		leaderSelectionApi: leaderSelectionApi,
		peerRepository:     repository,
		peerTable:          table,
		messageProducer:    producer,
	}
}

// connection이 발생하면 처리하는 메소드이다.
func (ml MessageListener) HandleConnCreateEvent(amqpMessage amqp.Delivery) {
	connCreateEevent := &event.ConnCreateEvent{}
	err := json.Unmarshal(amqpMessage.Body, connCreateEevent)

	if err != nil {
		// todo amqp error handle
	}
	newPeer := model.NewPeer(connCreateEevent.Address, model.PeerId(connCreateEevent.Id))
	(*ml.peerRepository).Save(*newPeer)
	if ml.peerTable.GetLeader() == nil {
		err = (*ml.messageProducer).RequestLeaderInfo(*newPeer)
		if err != nil {
			// todo amqp error handle
		}
	}
}

func (ml MessageListener) HandleMessageReceiveEvent(amqpMessage amqp.Delivery) {
	receiveEvent := &event.MessageReceiveEvent{}
	err := json.Unmarshal(amqpMessage.Body, receiveEvent)
	if err != nil {
		// todo amqp error handle
	}
	// handle 해야될거만 확인 아니면 버려~
	if receiveEvent.Protocol == topic.LeaderInfoRequestCmd.String() {
		curLeader := ml.peerTable.GetLeader()
		if curLeader == nil {
			curLeader = &model.Peer{
				IpAddress: "",
				Id:        "",
			}
		}
		// todo error handle
		toPeer, _ := (*ml.peerRepository).FindById(model.PeerId(receiveEvent.SenderId))
		// todo error handle
		err = (*ml.messageProducer).DeliverLeaderInfo(*toPeer, *curLeader)

	} else if receiveEvent.Protocol == topic.LeaderInfoPublishEvent.String() {
		eventBody := &event.LeaderInfoPublishEvent{}
		// todo error handle
		err = common.Deserialize(receiveEvent.Body, eventBody)
		leader := model.NewPeer(eventBody.Address, model.PeerId(eventBody.LeaderId))
		ml.peerTable.SetLeader(leader)
	}

}
