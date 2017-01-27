package sql

import (
	"errors"
)

const (
	CONFIG_DEFAULT = "default"
)

// db config
type Config struct {
	Host     string
	Database string
}

// db session
type Session struct {
	configs map[string]*Config
	dbs     map[string]*DB
}

// open with session
func (self *Session) Open(name string) (db *DB, err error) {
	if name == "" {
		name = CONFIG_DEFAULT
	}

	// exist check
	db = self.dbs[name]
	if db != nil {
		return
	}

	// get config
	config := self.configs[name]
	if config == nil {
		err = errors.New("no db config. name " + name)
		return
	}

	// open
	db, err = OpenByConfig(config)
	if err != nil {
		return
	}

	// set
	self.dbs[name] = db
	return
}

// close db and delete instance
func (self *Session) Close(name string) (err error) {
	if name == "" {
		name = CONFIG_DEFAULT
	}

	db := self.dbs[name]
	if db == nil {
		return
	}

	err = db.DB.Close()
	if err != nil {
		return err
	}

	delete(self.dbs, name)
	return
}

// close db all
func (self *Session) CloseAll() (err error) {
	for name, _ := range self.dbs {
		err = self.Close(name)
		if err != nil {
			return err
		}
	}
	return
}

// set configs direct
func (self *Session) SetConfigs(configs map[string]*Config) {
	self.configs = configs
}

// add config to session
func (self *Session) AddConfig(name string, config *Config) {
	self.configs[name] = config
}

// create session
func NewSession() *Session {
	return &Session{
		configs: map[string]*Config{},
		dbs:     map[string]*DB{},
	}
}
