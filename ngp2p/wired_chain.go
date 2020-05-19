package ngp2p

import (
	"github.com/libp2p/go-libp2p-core/network"

	"github.com/ngchain/ngcore/ngtypes"
	"github.com/ngchain/ngcore/utils"
)

// chain will send peer the specific vault's chain, which's len is not must be full BlockCheckRound num
func (w *wiredProtocol) chain(uuid []byte, stream network.Stream, blocks ...*ngtypes.Block) bool {
	if len(blocks) == 0 {
		return false
	}

	log.Debugf("replying chain to %s. Message id: %x, chain from block@%d to",
		stream.Conn().RemotePeer(), uuid, blocks[0].GetHeight(), blocks[len(blocks)-1].GetHeight(),
	)

	payload, err := utils.Proto.Marshal(&ChainPayload{
		Blocks: blocks,
	})
	if err != nil {
		log.Errorf("failed to sign pb data: %s", err)
		return false
	}

	// create message data
	resp := &Message{
		Header:  w.node.NewHeader(uuid, MessageType_CHAIN),
		Payload: payload,
	}

	// sign the data
	signature, err := signMessage(w.node.PrivKey(), resp)
	if err != nil {
		log.Errorf("failed to sign pb data")
		return false
	}

	// add the signature to the message
	resp.Header.Sign = signature

	err = w.node.replyToStream(stream, resp)
	if err != nil {
		log.Debugf("chain to: %s was sent. Message Id: %x", stream.Conn().RemotePeer(), resp.Header.MessageId)
		return false
	}

	log.Debugf("chain to: %s was sent. Message Id: %x", stream.Conn().RemotePeer(), resp.Header.MessageId)

	return true
}

// DecodeChainPayload unmarshal the raw and return the *pb.ChainPayload.
func DecodeChainPayload(rawPayload []byte) (*ChainPayload, error) {
	payload := &ChainPayload{}

	err := utils.Proto.Unmarshal(rawPayload, payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
