package data

type ChainFileConfig struct {
	NumEntries int64
}

type LinearProbeFileConfig struct {
	NumEntries int64
}

type IndexConfig struct {
	ChainFiles      []ChainFileConfig     // todo: ensure decreasing NumEntries
	DBLocation      string                // todo: optional read from env
	LinearProbeFile LinearProbeFileConfig // todo: switch to picking NumEntries from ChainFiles.[0]
}

var (
	_config = IndexConfig{
		ChainFiles:      []ChainFileConfig{{NumEntries: 10000}, {NumEntries: 1000}},
		DBLocation:      "db",
		LinearProbeFile: LinearProbeFileConfig{NumEntries: 10000},
	}
)
