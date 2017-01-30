package command

import ()

type Args []string

// get arg by index
func (self Args) Get(i int, def string) string {
	if len(self) > i {
		return self[i]
	}
	return def
}

// parce os.Args to Args
func Parce(args []string) Args {
	return Args(args)
}
