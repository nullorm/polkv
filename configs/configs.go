package configs

import (
	"path"
)

const (
	_maxKeys    = 10000
	_dbLocation = "db"
)

type ChainIndexFileConfig struct {
	FilePath   string // todo- calculate on init
	NumEntries int64
}

type LinearProbeIndexFileConfig struct {
	FilePath   string // todo- use constant?
	NumEntries int64
}

type DBConfig struct {
	ChainIndexFiles      []ChainIndexFileConfig     // todo- ensure decreasing NumEntries
	DBLocation           string                     // todo- optional read from env
	LinearProbeIndexFile LinearProbeIndexFileConfig // todo- switch to picking NumEntries from ChainFiles.[0]
}

var (
	DB = DBConfig{
		ChainIndexFiles: []ChainIndexFileConfig{
			{NumEntries: _maxKeys, FilePath: path.Join(_dbLocation, "index.chain.0")},
			{NumEntries: _maxKeys / 10, FilePath: path.Join(_dbLocation, "index.chain.1")}},
		DBLocation:           _dbLocation,
		LinearProbeIndexFile: LinearProbeIndexFileConfig{NumEntries: _maxKeys, FilePath: path.Join(_dbLocation, "index.linear-probe.0")},
	}
)
