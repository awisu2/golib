package local

import ()

// local writer
type LocalWriter struct {
	local []byte
}

func (self *LocalWriter) Write(p []byte) (n int, err error) {
	if self.local == nil {
		self.local = []byte{}
	}
	self.local = append(self.local, p...)
	return len(self.local), nil
}

func (self *LocalWriter) String() string { return string(self.local) }
