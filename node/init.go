package node

import (
	"github.com/json-iterator/go"
	"github.com/kooksee/log"
	"github.com/kooksee/srelay/sp2p"
	"github.com/kooksee/p2pkv/node/handlers"
)

var (
	json   = jsoniter.ConfigCompatibleWithStandardLibrary
	logger = log.New("package", "sp2p")
	hm     = sp2p.GetHManager()
)

func init() {
	hm.Registry("kvSet", handlers.KVSet)
	hm.Registry("kvGet", handlers.KVGet)
}
