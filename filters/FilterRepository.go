package filters

import (
	"github.com/astaxie/beego"
	"reflect"
	"fmt"
)

// a place to park all the filter functions.
// we do this so that we can find them via reflection

type FilterRepository struct{
	Repository map[string]beego.FilterFunc

}
var TheFilterRepository FilterRepository
func init() {
	fmt.Println("filters.Init")
	TheFilterRepository.Repository = make(map[string]beego.FilterFunc)
}


func (this *FilterRepository) FetchFilterFunc( name string ) beego.FilterFunc{

	filterMethodVal := reflect.ValueOf(this).MethodByName(name)
	// turn that into an interface{}
	filterMethodIface := filterMethodVal.Interface()
	// turn that into a function that has the expected signature
	filterMethod := filterMethodIface.(func() beego.FilterFunc)
	// call the method directly
	theFunc := filterMethod()
	return theFunc
}


