package adapter

import (
	"encoding/json"
	"errors"

	"github.com/it-chain/it-chain-Engine/blockchain"
)

var ErrAddBlock = errors.New("error when adding block")

type SyncBlockApi interface {
	SyncedCheck(block blockchain.Block) error
	SaveBlock(block blockchain.Block) error
}

type SyncCheckGrpcCommandService interface {
	SyncCheckResponse(block blockchain.Block) error
	ResponseBlock(peerId blockchain.PeerId, block blockchain.Block) error
}

type GrpcCommandHandler struct {
	blockApi           SyncBlockApi
	blockQueryApi      blockchain.BlockQueryApi
	grpcCommandService SyncCheckGrpcCommandService
}

func NewGrpcCommandHandler(blockApi SyncBlockApi, blockQueryService blockchain.BlockQueryApi, grpcCommandService SyncCheckGrpcCommandService) *GrpcCommandHandler {
	return &GrpcCommandHandler{
		blockApi:           blockApi,
		blockQueryApi:      blockQueryService,
		grpcCommandService: grpcCommandService,
	}
}

func (g *GrpcCommandHandler) HandleGrpcCommand(command blockchain.GrpcReceiveCommand) error {
	switch command.Protocol {
	case "SyncCheckRequestProtocol":
		//TODO: 상대방의 SyncCheck를 위해서 자신의 last block을 보내준다.
		block, err := g.blockQueryApi.GetLastCommitedBlock()
		if err != nil {
			return ErrGetLastBlock
		}

		err = g.grpcCommandService.SyncCheckResponse(block)
		if err != nil {
			return ErrSyncCheckResponse
		}
		break

	case "SyncCheckResponseProtocol":
		//TODO: 상대방의 last block을 받아서 SyncCheck를 시작한다.
		break

	case "BlockRequestProtocol":
		var height uint64
		err := json.Unmarshal(command.Body, &height)
		if err != nil {
			return ErrBlockInfoDeliver
		}

		block, err := g.blockQueryApi.GetCommitedBlockByHeight(height)
		if err != nil {
			return ErrGetBlock
		}

		err = g.grpcCommandService.ResponseBlock(command.FromPeer.PeerId, block)
		if err != nil {
			return ErrResponseBlock
		}
		break

	case "BlockResponseProtocol":
		//TODO: Construct 과정에서 block을 받는다.
		var block blockchain.DefaultBlock
		err := json.Unmarshal(command.Body, &block)
		if err != nil {
			return ErrBlockInfoDeliver
		}

		err = g.blockApi.SaveBlock(&block)
		if err != nil {
			return ErrAddBlock
		}

		break
	}

	return nil
}
