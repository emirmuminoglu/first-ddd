//go:generate easyjson -output_filename easyjson.go ./

package response

import "sync"

var (
	registerPool    sync.Pool
	zeroRegisterReq = &Register{}
)

func AcquireRegister() *Register {
	v := registerPool.Get()
	if v == nil {
		return new(Register)
	}

	return v.(*Register)
}

//easyjson:json
type Register struct {
	Reason string `json:"reason"`
}

func ReleaseRegister(req *Register) {
	*req = *zeroRegisterReq

	registerPool.Put(req)
}
