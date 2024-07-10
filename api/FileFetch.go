package api

import (
	"errors"
	"fmt"
	"github.com/wallet/base"
	"github.com/wallet/service/apiConnect"
	"github.com/wallet/service/universeCourier"
	"os"
	"path/filepath"
)

const (
	defaultlndpath  = ".lnd"
	defaultlitpath  = ".lit"
	defaulttapdpath = ".tapd"
)

const (
	UniverseHostMainnet = "universerpc://132.232.109.84:8444"
	UniverseHostTestnet = "universerpc://127.0.0.1:1235"
	UniverseHostRegtest = "universerpc://132.232.109.84:8443"
)
const (
	BtlServerMainnet = "132.232.109.84:8095"
	BtlServerTestNet = ""
	BtlServerRegTest = "132.232.109.84:8090"
)

type Config struct {
	Network        string `json:"network"`
	PostServiceUrl string `json:"postServiceUrl"`
	BtlServerHost  string `json:"btlServerHost"`
}

var Cfg Config

func SetPath(path string, network string) error {
	err := base.SetFilePath(path)
	if err != nil {
		return errors.New("path not exist")
	}
	if network != "mainnet" && network != "testnet" && network != "regtest" {
		return errors.New("network not exist")
	}
	base.SetNetwork(network)

	err = Cfg.loadConfig()
	if err != nil {
		return errors.New("load config error ")
	}
	return nil
}
func (c *Config) loadConfig() error {
	//load service config
	err := c.loadServiceConfig()
	if err != nil {
		return errors.New("load config error ")
	}
	//load config from file
	c.Network = base.NetWork
	//load config other
	switch {
	case base.NetWork == "mainnet":
		c.PostServiceUrl = UniverseHostMainnet
		c.BtlServerHost = BtlServerMainnet
	case base.NetWork == "testnet":
		c.PostServiceUrl = UniverseHostTestnet
		c.BtlServerHost = BtlServerTestNet
	case base.NetWork == "regtest":
		c.PostServiceUrl = UniverseHostRegtest
		c.BtlServerHost = BtlServerRegTest
	default:
		return errors.New("network not exist")
	}
	return nil
}

func (c *Config) loadServiceConfig() error {
	apiConnect.LoadConnectConfig()
	universeCourier.LoadUniverseCourierConfig()
	return nil
}

func GetPath() string {
	return base.GetFilePath()
}

const defaultbitcoinpath = "data/chain/bitcoin"

// CheckDir Check the integrity of the directory
func CheckDir(dir string) error {
	baseDir := dir
	//Check whether the snapshot file location exists
	neutrinoPath := filepath.Join(baseDir, defaultlndpath, defaultbitcoinpath, base.NetWork)
	fmt.Println(neutrinoPath)
	if !fileExists(neutrinoPath) {
		if err := os.MkdirAll(neutrinoPath, 0700); err != nil {
			return err
		}
	}
	return nil
}

// fileExists reports whether the named file or directory exists.
func fileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
