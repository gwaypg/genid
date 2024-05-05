package etc

import (
	"github.com/gwaylib/conf"
	"github.com/gwaylib/conf/ini"
)

var (
	Ini = ini.NewIni(conf.RootDir() + "/etc/")
	Etc = Ini.GetFile("etc.cfg")
)
