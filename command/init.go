package command

import (
	"github.com/codegangsta/cli"
)

var (
	Cat   CatCommand
	Touch TouchCommand
	Rm    RmCommand
	Mkdir MkdirCommand
	Ls    LsCommand
	Tar   TarCommand
)

func init() {
	Cat = CatCommand{
		Name:   "cat",
		Usage:  "print content of path.",
		Action: Cat.run,
	}
	Touch = TouchCommand{
		Name:   "touch",
		Usage:  "touch a key with value.",
		Action: Touch.run,
	}
	Rm = RmCommand{
		Name:  "rm",
		Usage: "remove keys or dirs",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "recursive,r,R",
				Usage: "remove directorieds and their content recursively.",
			},
		},
		Action: Rm.run,
	}
	Mkdir = MkdirCommand{
		Name:  "mkdir",
		Usage: "create directory",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "parent,p",
				Usage: "create parent directories if not exist.",
			},
		},
		Action: Mkdir.run,
	}
	Ls = LsCommand{
		Name:   "ls",
		Usage:  "recursive list item in given path",
		Action: Ls.run,
	}
	Tar = TarCommand{
		Name:  "tar",
		Usage: "archive/extract kv data",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "extract,x",
				Usage: "extract from archive file",
			},
			cli.BoolFlag{
				Name:  "create,c",
				Usage: "create an archive file",
			},
			cli.StringFlag{
				Name:  "file,f",
				Value: "-",
				Usage: "archive file, '-' for stdin/stdout",
			},
		},
		Action: Tar.run,
	}
}
