package sammy

import (
	"github.com/eyedeekay/sam3"
	"github.com/eyedeekay/sam3/i2pkeys"
	"os"
)

func RandStringBytes() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 4)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Sammy(opts map[string]string) (*sam3.StreamListener, error) {
	sam := ""
	keypath := ""
	servertun := ""
	if opts != nil {
		if ok, _ := opts[`sam`]; !ok {
			opts[`sam`] = "127.0.0.1:7656"
		}
		if ok, _ := opts[`keypath`]; !ok {
			opts[`keypath`] = "keys.i2pkeys"
		}
		if ok, _ := opts[`servertun`]; !ok {
			opts[`servertun`] = "serv" + randStringBytes()
		}
	} else {
		opts := make(map[string]string)
		opts[`sam`] = "127.0.0.1:7656"
		opts[`keypath`] = "keys.i2pkeys"
		opts[`servertun`] = "serv" + randStringBytes()
	}
	if sam, err := sam3.NewSAM(opts[`sam`]); err != nil {
		return nil, err
	} else {
		if file, err := os.Open(opt[`keypath`]); err == nil {
			if keys, err := i2pkeys.LoadKeysIncompat(file); err != nil {
				return nil, err
			} else {
				if stream, err := sam.NewStreamSession(opt[`servertun`], keys, sam3.Options_Fat); err != nil {
					return nil, err
				} else {
					return stream.Listen()
				}
			}
		} else {
			if keys, err := sam.NewKeys(); err != nil {
				return nil, err
			} else {
				if file, err := os.Create(opt[`keypath`]); err != nil {
					return nil, err
				} else {
					if err := i2pkeys.StoreKeysIncompat(keys, file); err != nil {
						return nil, err
					}
					if stream, err := sam.NewStreamSession(opt[`servertun`], keys, sam3.Options_Fat); err != nil {
						return nil, err
					} else {
						return stream.Listen()
					}
				}
			}
		}
	}

}
