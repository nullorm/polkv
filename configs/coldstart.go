package configs

import (
	"os"
	"path"

	"github.com/nullorm/polkv/app"
	"go.uber.org/zap"
)

func (c DBConfig) coldStartData() {
	// todo- move to config
	ensureFile(path.Join(c.DBLocation, "data.80b.0"), _maxKeys*10)
	ensureFile(path.Join(c.DBLocation, "data.80b.1"), _maxKeys*10)

	ensureFile(path.Join(c.DBLocation, "data.800b.0"), _maxKeys*100)
	ensureFile(path.Join(c.DBLocation, "data.800b.1"), _maxKeys*100)

	ensureFile(path.Join(c.DBLocation, "data.8kb.0"), _maxKeys*1000)
	ensureFile(path.Join(c.DBLocation, "data.8kb.1"), _maxKeys*1000)

	ensureFile(path.Join(c.DBLocation, "data.80kb.0"), _maxKeys*10000)
	ensureFile(path.Join(c.DBLocation, "data.80kb.1"), _maxKeys*10000)

	ensureFile(path.Join(c.DBLocation, "data.800kb.0"), _maxKeys*100000)
	ensureFile(path.Join(c.DBLocation, "data.800kb.1"), _maxKeys*100000)
}

func (c DBConfig) coldStartIndex() {
	for _, config := range c.ChainIndexFiles {
		ensureFile(config.FilePath, config.NumEntries*8)
	}

	ensureFile(c.LinearProbeIndexFile.FilePath, c.LinearProbeIndexFile.NumEntries*8)
}

func (c DBConfig) coldStart() {
	// todo- refactor when implementing warmstart
	if err := os.RemoveAll(c.DBLocation); err != nil {
		app.Log.Fatal("unable to remove dir", zap.String("dblocaltion", c.DBLocation), zap.Error(err))
	}

	ensureDir(c.DBLocation)

	c.coldStartIndex()
	c.coldStartData()

}
