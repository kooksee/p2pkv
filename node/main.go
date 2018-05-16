package node

import (
	"github.com/kooksee/srelay/sp2p"
	"github.com/kooksee/srelay/types"
	"github.com/kooksee/p2pkv/protocol"
	"github.com/kooksee/uspnet/common"
)

type KVNode struct {
	*sp2p.SP2p
}

func (n *KVNode) Set(req *protocol.KVGetReq) {
	nodes := n.GetTable().FindNodeWithTargetBySelf(common.StringToHash(req.K))
	if len(nodes) == 0 {
		return
	}

	for _, node := range nodes {
		if err := n.Write(&types.KMsg{
			Event: "kvSet",
			Data:  req,
			TAddr: node.Addr().String(),
		}); err != nil {
			logger.Error(err.Error())
		}
	}
}

func (n *KVNode) Get(req protocol.KVGetReq) {
	nodes := n.GetTable().FindNodeWithTargetBySelf(common.StringToHash(req.K))
	if len(nodes) == 0 {
		return
	}

	for _, node := range n.GetTable().FindNodeWithTargetBySelf(common.StringToHash(req.K)) {
		if err := n.Write(&types.KMsg{
			Event: "kvGet",
			Data:  req,
			TAddr: node.Addr().String(),
		}); err != nil {
			logger.Error(err.Error())
		}
	}
}
