package configs

import (
	"os"
	"strconv"

	"github.com/nullorm/polkv/app"
	"go.uber.org/zap"
)

func (c Config) coldStartData() {
	ensureFile(c.DBLocation, "data.80b.0", _maxKeys*10)
	ensureFile(c.DBLocation, "data.80b.1", _maxKeys*10)

	ensureFile(c.DBLocation, "data.800b.0", _maxKeys*100)
	ensureFile(c.DBLocation, "data.800b.1", _maxKeys*100)

	ensureFile(c.DBLocation, "data.8kb.0", _maxKeys*1000)
	ensureFile(c.DBLocation, "data.8kb.1", _maxKeys*1000)

	ensureFile(c.DBLocation, "data.80kb.0", _maxKeys*10000)
	ensureFile(c.DBLocation, "data.80kb.1", _maxKeys*10000)

	ensureFile(c.DBLocation, "data.800kb.0", _maxKeys*100000)
	ensureFile(c.DBLocation, "data.800kb.1", _maxKeys*100000)
}

func (c Config) coldStartIndex() {
	for i, config := range c.ChainIndexFiles {
		ensureFile(c.DBLocation, "index.chain."+strconv.Itoa(i), config.NumEntries*8)
	}

	ensureFile(c.DBLocation, "index.linear-probe.0", c.LinearProbeIndexFile.NumEntries*8)
}

func (c Config) coldStart() {
	// todo - refactor when implementing warmstart
	if err := os.RemoveAll(c.DBLocation); err != nil {
		app.Log.Fatal("unable to remove dir", zap.String("dblocaltion", c.DBLocation), zap.Error(err))
	}

	ensureDir(c.DBLocation)

	c.coldStartIndex()
	c.coldStartData()

}
