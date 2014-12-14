package filterInfrastructure

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ghstahl/pingbeego/reflection"
)

func init() {
	fmt.Println("filters.Init")
}


func FetchFilterFunc(structTypeName string ) beego.FilterFunc{

	filterMethodIface := reflection.TheReflectRepository.GetInterface(structTypeName,"FilterFunc")
	// turn that into a function that has the expected signature
	filterMethod := filterMethodIface.(func() beego.FilterFunc)
	// call the method directly
	theFunc := filterMethod()
	return theFunc
}


