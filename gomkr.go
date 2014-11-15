package main

import (
	"os"
	"runtime"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gomkr"
	app.Version = Version
	app.Usage = "A CLI tool for mackerel.io"
	app.Author = "Hatena Co., Ltd."
	app.Email = "y_uuki@hatena.ne.jp"
	app.Commands = Commands

	cpu := runtime.NumCPU()
	runtime.GOMAXPROCS(cpu)

	app.Run(os.Args)
}
