package sammy

import (
	"github.com/eyedeekay/sam3"
	"github.com/eyedeekay/sam3/i2pkeys"

	"math/rand"
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
	if opts != nil {
		if _, ok := opts[`sam`]; !ok {
			opts[`sam`] = "127.0.0.1:7656"
		}
		if _, ok := opts[`servertun`]; !ok {
			opts[`servertun`] = "serv" + RandStringBytes()
		}
		if _, ok := opts[`keypath`]; !ok {
			opts[`keypath`] = opts[`servertun`] + ".i2pkeys"
		}
	} else {
		opts := make(map[string]string)
		opts[`sam`] = "127.0.0.1:7656"
		opts[`servertun`] = "serv" + RandStringBytes()
		opts[`keypath`] = opts[`servertun`] + ".i2pkeys"
	}
	if sam, err := sam3.NewSAM(opts[`sam`]); err != nil {
		return nil, err
	} else {
		if file, err := os.Open(opts[`keypath`]); err == nil {
			if keys, err := i2pkeys.LoadKeysIncompat(file); err != nil {
				return nil, err
			} else {
				if stream, err := sam.NewStreamSession(opts[`servertun`], keys, sam3.Options_Fat); err != nil {
					return nil, err
				} else {
					return stream.Listen()
				}
			}
		} else {
			if keys, err := sam.NewKeys(); err != nil {
				return nil, err
			} else {
				if file, err := os.Create(opts[`keypath`]); err != nil {
					return nil, err
				} else {
					if err := i2pkeys.StoreKeysIncompat(keys, file); err != nil {
						return nil, err
					}
					if stream, err := sam.NewStreamSession(opts[`servertun`], keys, sam3.Options_Fat); err != nil {
						return nil, err
					} else {
						return stream.Listen()
					}
				}
			}
		}
	}

}
