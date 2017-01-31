package command

import (
	"strconv"
)

type Args []string

// get arg by index
func (self Args) Get(i int, def string) string {
	if len(self) > i {
		return self[i]
	}
	return def
}

// get arg by index
func (self Args) Int(i int, def int) (int, error) {
	s := ""
	if len(self) > i {
		s = self[i]
		i, err := strconv.Atoi(s)
		if err != nil {
			return def, err
		}
		return i, nil
	}
	return def, nil
}

// parce os.Args to Args
func Parce(args []string) Args {
	return Args(args)
}
