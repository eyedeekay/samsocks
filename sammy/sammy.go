package sammy

import (
	"github.com/eyedeekay/sam3"
	"github.com/eyedeekay/sam3/i2pkeys"
	"os"
)

func Sammy() (*sam3.StreamListener, error) {
	if sam, err := sam3.NewSAM("127.0.0.1:7656"); err != nil {
		return nil, err
	} else {
		if file, err := os.Open("../keys.i2pkeys"); err == nil {
			if keys, err := i2pkeys.LoadKeysIncompat(file); err != nil {
				return nil, err
			} else {
				if stream, err := sam.NewStreamSession("serverTun", keys, sam3.Options_Fat); err != nil {
					return nil, err
				} else {
					return stream.Listen()
				}
			}
		} else {
			if keys, err := sam.NewKeys(); err != nil {
				return nil, err
			} else {
				if file, err := os.Create("../keys.i2pkeys"); err != nil {
					return nil, err
				} else {
					if err := i2pkeys.StoreKeysIncompat(keys, file); err != nil {
						return nil, err
					}
					if stream, err := sam.NewStreamSession("serverTun", keys, sam3.Options_Fat); err != nil {
						return nil, err
					} else {
						return stream.Listen()
					}
				}
			}
		}
	}

}
