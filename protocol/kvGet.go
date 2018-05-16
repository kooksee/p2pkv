package protocol

type KVGetReq struct {
	K string
	V interface{}
}

type KVGetResp struct {
	K string
}
