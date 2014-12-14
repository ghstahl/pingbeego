package filters

import (
	"github.com/astaxie/beego/context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ghstahl/pingbeego/reflection"
	"reflect"
 	"code.google.com/p/go-uuid/uuid"
)

type RequestIdFilterType struct {
}

func (v RequestIdFilterType) FilterFunc() beego.FilterFunc{
	return func(ctx *context.Context) {
		fmt.Println(fmt.Sprintf(">>>>>>>>>>>>>>RequestIdFilterType Enter"))
		ctx.Input.Data[WellKnown.RequestId] = uuid.New();
		fmt.Println(fmt.Sprintf("$$$$$$$$$$$$$$$ %v $$$$$$$$$$$$$$$",ctx.Input.Data[WellKnown.RequestId]))
	}
}

func GetCurrentRequestId(ctx *context.Context) string {

	ok := ctx.Input.Data[WellKnown.RequestId]
	if(ok == nil){
		return "RequestId NOT Here"
	}
	return ok.(string)
}

func init() {
	var x RequestIdFilterType
	reflection.TheReflectRepository.ValueRepository["RequestIdFilterType"] = reflect.ValueOf(x)
	reflection.TheReflectRepository.TypeRepository["RequestIdFilterType"] = reflect.TypeOf(x)
}

