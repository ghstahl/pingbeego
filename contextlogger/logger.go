package contextlogger

import (
	"github.com/astaxie/beego/context"
	"github.com/ghstahl/pingbeego/contextlogger"
)

var WellKnown = struct {
		RequestIdTag string
	}{
	"RequestIdTag",
}

func PrepareLogContextTags(ctx *context.Context) map[string]interface{} {
	m2 := make(map[string]interface{})
	m2[contextlogger.WellKnown.RequestIdTag] = makeRequestIdTag(ctx)
	return m2
}

type RequestIdTag struct {
	Ctx *context.Context
}

func (self *RequestIdTag) String() string {
	result := this.Ctx.Input.Data[WellKnown.RequestId]
	if result == nil {
		return WellKnown.RequestId +":not found"
	}
	return 	result.(string)
}

func makeRequestIdTag(ctx *context.Context) RequestIdTag {
	return RequestIdTag{ctx}
}
