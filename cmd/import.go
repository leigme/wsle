package cmd

/*
Copyright © 2023 leig <leigme@gmail.com>

*/

import (
	"fmt"
	loki "github.com/leigme/loki/cobra"
	"github.com/spf13/cobra"
)

func init() {
	loki.Add(rootCmd, &wslImport{})
}

type wslImport struct {
}

func (wi *wslImport) Execute() loki.Exec {
	return func(cmd *cobra.Command, args []string) {

	}
}

// wsl --import [Ubuntu-22.04] [C:\Users\leig\.wsl\] [C:\Users\leig\.wsl\backup\Ubuntu-22.04_20060102150405.tar
func generateImportCmd(linux, importDir, importFile string) string {
	return fmt.Sprintf("wsl --import %s %s %s", linux, importDir, importFile)
}
