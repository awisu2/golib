package io

import ()

// ValueWriter Override io.Writer interface
// write to propaty variable
type ValueWriter struct {
	p []byte
}

func (self *ValueWriter) Write(p []byte) (n int, err error) {
	if self.p == nil {
		self.p = []byte{}
	}
	self.p = append(self.p, p...)
	return len(p), nil
}

func (self *ValueWriter) Bytes() []byte {
	return self.p
}

func (self *ValueWriter) String() (s string) {
	if self.p == nil {
		return ""
	}
	return string(self.p)
}

func (self *ValueWriter) Clear() {
	self.p = nil
}
