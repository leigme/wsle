package common

import (
	"fmt"
	"github.com/leigme/loki/app"
	"github.com/leigme/loki/file"
	"path/filepath"
	"time"
)

/*
Copyright © 2023 leig <leigme@gmail.com>

*/

const (
	wslWorkDir = "wsl"
	wslBackup  = "backup"
)

func ExportDefaultFile(linux string) string {
	filename := fmt.Sprintf("%s_%s.tar", linux, time.Now().Format("200601021504"))
	if err := file.CreateDir(filepath.Join(app.WorkDir(), wslWorkDir, wslBackup)); err == nil {
		filename = filepath.Join(app.WorkDir(), wslWorkDir, wslBackup, filename)
	}
	return filename
}
