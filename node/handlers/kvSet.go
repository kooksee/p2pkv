package handlers

import (
	"github.com/kooksee/srelay/sp2p"
	"github.com/kooksee/srelay/types"
	"github.com/kooksee/p2pkv/protocol"
)

func KVSet(p *sp2p.SP2p, msg *types.KMsg) {
	req, ok := msg.Data.(protocol.KVGetReq)
	if !ok {
		return
	}

	p.GetTable()
}
