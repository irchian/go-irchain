// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package huc

import (
	"math/big"

	"github.com/happyuc-project/happyuc-go/common"
	"github.com/happyuc-project/happyuc-go/common/hexutil"
	"github.com/happyuc-project/happyuc-go/consensus/huchash"
	"github.com/happyuc-project/happyuc-go/core"
	"github.com/happyuc-project/happyuc-go/huc/downloader"
	"github.com/happyuc-project/happyuc-go/huc/gasprice"
)

var _ = (*configMarshaling)(nil)

func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               uint64
		SyncMode                downloader.SyncMode
		LightServ               int  `toml:",omitempty"`
		LightPeers              int  `toml:",omitempty"`
		SkipBcVersionCheck      bool `toml:"-"`
		DatabaseHandles         int  `toml:"-"`
		DatabaseCache           int
		Coinbase                common.Address `toml:",omitempty"`
		MinerThreads            int            `toml:",omitempty"`
		ExtraData               hexutil.Bytes  `toml:",omitempty"`
		GasPrice                *big.Int
		Huchash                 huchash.Config
		TxPool                  core.TxPoolConfig
		GPO                     gasprice.Config
		EnablePreimageRecording bool
		DocRoot                 string `toml:"-"`
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.LightServ = c.LightServ
	enc.LightPeers = c.LightPeers
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.Coinbase = c.Coinbase
	enc.MinerThreads = c.MinerThreads
	enc.ExtraData = c.ExtraData
	enc.GasPrice = c.GasPrice
	enc.Huchash = c.Huchash
	enc.TxPool = c.TxPool
	enc.GPO = c.GPO
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.DocRoot = c.DocRoot
	return &enc, nil
}

func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               *uint64
		SyncMode                *downloader.SyncMode
		LightServ               *int  `toml:",omitempty"`
		LightPeers              *int  `toml:",omitempty"`
		SkipBcVersionCheck      *bool `toml:"-"`
		DatabaseHandles         *int  `toml:"-"`
		DatabaseCache           *int
		Coinbase                *common.Address `toml:",omitempty"`
		MinerThreads            *int            `toml:",omitempty"`
		ExtraData               *hexutil.Bytes  `toml:",omitempty"`
		GasPrice                *big.Int
		Huchash                 *huchash.Config
		TxPool                  *core.TxPoolConfig
		GPO                     *gasprice.Config
		EnablePreimageRecording *bool
		DocRoot                 *string `toml:"-"`
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.LightServ != nil {
		c.LightServ = *dec.LightServ
	}
	if dec.LightPeers != nil {
		c.LightPeers = *dec.LightPeers
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.Coinbase != nil {
		c.Coinbase = *dec.Coinbase
	}
	if dec.MinerThreads != nil {
		c.MinerThreads = *dec.MinerThreads
	}
	if dec.ExtraData != nil {
		c.ExtraData = *dec.ExtraData
	}
	if dec.GasPrice != nil {
		c.GasPrice = dec.GasPrice
	}
	if dec.Huchash != nil {
		c.Huchash = *dec.Huchash
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	return nil
}
