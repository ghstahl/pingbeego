package filters

import (
	"github.com/astaxie/beego/context"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ghstahl/pingbeego/reflection"
	"github.com/ghstahl/pingbeego/contextmanager"
	"reflect"
 	"code.google.com/p/go-uuid/uuid"
	"github.com/jtolds/gls"
)

type RequestIdFilterType struct {
}

func (v RequestIdFilterType) FilterFunc() beego.FilterFunc{
	return func(ctx *context.Context) {
		fmt.Println(fmt.Sprintf(">>>>>>>>>>>>>>RequestIdFilterType Enter"))
		requestId := uuid.New();
		ctx.Input.Data[WellKnown.RequestId] = requestId;
		contextmanager.WellKnown.ContextManager.SetValues(gls.Values{contextmanager.WellKnown.BeegoContextKey: ctx}, func() {});
		fmt.Println(fmt.Sprintf("$$$$$$$$$$$$$$$ %v $$$$$$$$$$$$$$$",ctx.Input.Data[WellKnown.RequestId]))
//		ctx.Input.Data[WellKnown.RequestId] = requestId
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

