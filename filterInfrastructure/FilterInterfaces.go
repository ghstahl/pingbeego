package filterInfrastructure

import (
	"github.com/astaxie/beego"
)

type  FilterFuncer interface {
	FilterFunc() beego.FilterFunc
}

