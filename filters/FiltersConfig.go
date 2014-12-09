package filters

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/astaxie/beego"
)

/*
{
    "filters":
    [
        {
          	"struct":"UserFilterType",
			"pattern":"*",
			"position":"beego.AfterStatic"

        }
    ]
}
*/
func init() {
	fmt.Println("filtersconfig filters.Init")
}

type FilterType struct {
	Struct		string
	Pattern 	string
	Position    string
}
type FilterConfigJsonObject struct {
	Filters   []FilterType
}

type FilterConfigs struct{
	Config FilterConfigJsonObject
	beegoFilterPositionLookup  map[string]int
}

var TheFilterConfigs = &FilterConfigs{

}

func init() {
	fmt.Println("FiltersConfig.Init")
	TheFilterConfigs.beegoFilterPositionLookup = map[string]int{
		"beego.BeforeRouter":beego.BeforeRouter,
		"beego.BeforeExec":beego.BeforeExec,
		"beego.FinishRouter":beego.FinishRouter,
		"beego.BeforeStatic":beego.BeforeStatic,
	}
}


func (this *FilterConfigs) loadJsonConfig() {
	file, e := ioutil.ReadFile("conf/filterconfigs.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))
	json.Unmarshal(file, &this.Config)
	fmt.Printf("Results: %v\n", this.Config)
}

func (this *FilterConfigs) Load() {
	fmt.Printf("beegoFilterPositionLookup:%v\n",this.beegoFilterPositionLookup)
	this.loadJsonConfig()

	for _,item:= range this.Config.Filters{
		position := this.beegoFilterPositionLookup[item.Position]
		fmt.Println(fmt.Sprintf("item: %v position: %v",item,position))
		theFilterFunc := FetchFilterFunc(item.Struct)
		beego.InsertFilter(item.Pattern, position, theFilterFunc)
	}
}
