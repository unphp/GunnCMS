// utils.go
package utils

import (
	"log"
	"runtime"
)

var (
	SysName string
)

//系统信息
type Platform struct {
	OS, ID, Release, Codename, Description string
}

func NewPlatform() *Platform {
	return new(Platform)
}

func (p *Platform) GetAll() *Platform {
	return p.get_Plat_info()
}

//end time 2014-03-03 15:50
func (p *Platform) get_Plat_info() *Platform {
	var (
		file = "/etc/lsb-release"
		rs   []string
	)

	//检查文件是否存在
	if IsExist(file) {
		Rs, err := ReadFile(file)
		if err != nil {
			log.Panicln(err.Error())
		}
		rs = Explode(string(Rs), "\n")
	}
	//得到文件信息给 Platform
	p.OS = runtime.GOOS
	for _, v := range rs {
		vv := Explode(v, "=")
		k := vv[0]
		switch k {
		case "DISTRIB_ID":
			p.ID = StrReplace(Trim(vv[1]), "\"", "")
		case "DISTRIB_RELEASE":
			p.Release = StrReplace(Trim(vv[1]), "\"", "")
		case "DISTRIB_CODENAME":
			p.Codename = StrReplace(Trim(vv[1]), "\"", "")
		case "DISTRIB_DESCRIPTION":
			p.Description = StrReplace(Trim(vv[1]), "\"", "")
		}
	}
	return p
}
