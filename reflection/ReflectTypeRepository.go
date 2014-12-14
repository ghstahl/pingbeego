package reflection
/*
this is a parking lot for types that we later want to reference through reflection.
we need to put these things here because the GO compiler will optimize out unused code, but since we are going to configure
in use of the code, we need to tell GO to keep this stuff in here.

usage:

func init() {
	var x RequestIdFilterType
	reflection.TheReflectRepository.ValueRepository["RequestIdFilterType"] = reflect.ValueOf(x)
	reflection.TheReflectRepository.TypeRepository["RequestIdFilterType"] = reflect.TypeOf(x)
}

 */
import (
	"fmt"
	"reflect"
)

type ReflectRepositoryType struct{
	ValueRepository map[string]reflect.Value
	TypeRepository map[string]reflect.Type
}

var TheReflectRepository ReflectRepositoryType

func init() {
	fmt.Println("reflection.ReflectTypeRepsoitory.Init")
	TheReflectRepository.ValueRepository = make(map[string]reflect.Value)
	TheReflectRepository.TypeRepository = make(map[string]reflect.Type)
}

func (this *ReflectRepositoryType) GetInterface(structTypeName string, methodName string ) interface{}{
	var reflectValue reflect.Value = this.ValueRepository[structTypeName]
	filterMethodVal := reflectValue.MethodByName(methodName)
	// turn that into an interface{}
	filterMethodIface := filterMethodVal.Interface()
	return filterMethodIface
}
