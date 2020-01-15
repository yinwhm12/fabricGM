package factory

import (
	"github.com/hyperledger/fabric/bccsp"
	"github.com/pkg/errors"
	"github.com/hyperledger/fabric/bccsp/gm"
)

const GuomiBasedFactoryName  = "GM"

type GMFactory struct {}

func (f *GMFactory) Name() string  {
	return GuomiBasedFactoryName
}

func (f *GMFactory) Get(config *FactoryOpts) (bccsp.BCCSP, error)  {
	if config == nil || config.SwOpts == nil{
		return nil, errors.New("Invalid config. It must not be nil.")
	}

	swOpts := config.SwOpts

	var ks bccsp.KeyStore
	if swOpts.Ephemeral == true{
		ks = gm.NewDummyKeyStore()
	} else if swOpts.FileKeystore != nil{
		//logger.Infof("GMFactory Get swOpts.FileKeystore.KeyStorePath=%s",swOpts.FileKeystore.KeyStorePath)
		fks, err := gm.NewFileBasedKeyStore(nil, swOpts.FileKeystore.KeyStorePath, false)
		if err != nil{
			return nil, errors.Wrapf(err, "Failed to initialize software key store")
		}
		ks = fks
	} else {
		// Default to ephemeral key store
		ks = gm.NewDummyKeyStore()
	}
	//todo 是否需要内存存储
	// else if swOpts.InmemKeystore != nil {
	//	ks = sw.NewInMemoryKeyStore()
	return gm.NewWithParams(swOpts.SecLevel, swOpts.HashFamily, ks)
}
