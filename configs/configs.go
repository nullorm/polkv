package configs

const (
	_maxKeys = 10000
)

type ChainIndexFileConfig struct {
	NumEntries int64
}

type LinearProbeIndexFileConfig struct {
	NumEntries int64
}

type Config struct {
	ChainIndexFiles      []ChainIndexFileConfig     // todo: ensure decreasing NumEntries
	DBLocation           string                     // todo: optional read from env
	LinearProbeIndexFile LinearProbeIndexFileConfig // todo: switch to picking NumEntries from ChainFiles.[0]
}

var (
	_config = Config{
		ChainIndexFiles:      []ChainIndexFileConfig{{NumEntries: _maxKeys}, {NumEntries: _maxKeys / 10}},
		DBLocation:           "db",
		LinearProbeIndexFile: LinearProbeIndexFileConfig{NumEntries: _maxKeys},
	}
)
