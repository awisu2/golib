package command

import (
	"os"
)

func GetArgs() Args {
	return Parce(os.Args)
}
