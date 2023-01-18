package configs

import (
	"os"
	"path"

	"github.com/nullorm/polkv/app"
	"go.uber.org/zap"
)

func ensureDir(dirpath string) {
	if fs, err := os.Stat(dirpath); err != nil {
		if os.IsNotExist(err) {
			// create new dir
			if err := os.Mkdir(dirpath, os.ModePerm); err != nil {
				app.Log.Fatal("unable to create dir", zap.String("dblocaltion", dirpath), zap.Error(err))
			}
			app.Log.Debug("created dir", zap.String("dblocaltion", dirpath))
		} else {
			app.Log.Fatal("unable to verify dir", zap.String("dblocaltion", dirpath), zap.Error(err))
		}
	} else if !fs.IsDir() {
		app.Log.Fatal("file exists but is not a dir", zap.String("dblocaltion", dirpath))
	}
}

func ensureFile(dirpath, filename string, size int64) {
	filepath := path.Join(dirpath, filename)
	if fs, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			// create new file
			f, err := os.Create(filepath)
			if err != nil {
				app.Log.Fatal("unable to create file", zap.String("filepath", filepath), zap.Error(err))
			}
			if err := f.Truncate(size); err != nil {
				app.Log.Fatal("unable to truncate file", zap.String("filepath", filepath), zap.Error(err))
			}
			app.Log.Debug("created file", zap.String("filepath", filepath))
		} else {
			app.Log.Fatal("unable to verify file", zap.String("filepath", filepath), zap.Error(err))
		}
	} else if fs.IsDir() {
		app.Log.Fatal("file exists but is a dir", zap.String("filepath", filepath))
	}
}
