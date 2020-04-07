package ngp2p

import (
	"io/ioutil"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	core "github.com/libp2p/go-libp2p-core"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/ngchain/ngcore/ngp2p/pb"
	"github.com/ngchain/ngcore/ngtypes"
)

func (w *Wired) Ping(remotePeerID peer.ID) bool {
	payload, err := proto.Marshal(&pb.PingPongPayload{
		BlockHeight:     w.node.consensus.GetLatestBlockHeight(),
		VaultHeight:     w.node.consensus.GetLatestVaultHeight(),
		LatestBlockHash: w.node.consensus.GetLatestBlockHash(),
		LatestVaultHash: w.node.consensus.GetLatestVaultHash(),
	})
	if err != nil {
		log.Errorf("failed to sign pb data")
		return false
	}

	// create message data
	req := &pb.Message{
		Header:  w.node.NewHeader(uuid.New().String()),
		Payload: payload,
	}

	// sign the data
	signature, err := w.node.signMessage(req)
	if err != nil {
		log.Errorf("failed to sign pb data")
		return false
	}

	// add the signature to the message
	req.Header.Sign = signature

	ok := w.node.sendProtoMessage(remotePeerID, pingMethod, req)
	if !ok {
		log.Infof("failed sending ping to: %s.", remotePeerID)
		return false
	}

	// store ref request so response handler has access to it
	w.requests.Store(req.Header.Uuid, req)
	log.Debugf("Sent ping to: %s was sent. Message Id: %s.", remotePeerID, req.Header.Uuid)
	return true
}

// remote peer requests handler
func (w *Wired) onPing(s network.Stream) {
	// get request data
	buf, err := ioutil.ReadAll(s)
	if err != nil {
		_ = s.Reset()
		log.Error(err)
		return
	}

	remotePeerID := s.Conn().RemotePeer()
	_ = s.Close()

	// unmarshal it
	var data = &pb.Message{}
	err = proto.Unmarshal(buf, data)
	if err != nil {
		log.Error(err)
		go w.Reject(remotePeerID, data.Header.Uuid)

		return
	}

	if !w.node.authenticateMessage(s.Conn().RemotePeer(), data) {
		log.Errorf("Failed to authenticate message")
		return
	}

	var ping = &pb.PingPongPayload{}
	err = proto.Unmarshal(data.Payload, ping)
	if err != nil {
		log.Error(err)
		go w.Reject(remotePeerID, data.Header.Uuid)

		return
	}

	log.Debugf("Received ping request from %s. Remote height: %d", s.Conn().RemotePeer(), ping.BlockHeight)

	// Pong
	w.node.Peerstore().AddAddrs(s.Conn().RemotePeer(), []core.Multiaddr{s.Conn().RemoteMultiaddr()}, ngtypes.TargetTime*ngtypes.BlockCheckRound*ngtypes.BlockCheckRound)

	go w.Pong(remotePeerID, data.Header.Uuid)
}
