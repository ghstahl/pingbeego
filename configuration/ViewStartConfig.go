package configuration

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

/*
{
    "ViewStartTypes":
        [
            {
                "Name":"Primary",
                "SharedHead":"shared/_head.tpl",
                "Header":"shared/_header.tpl",
                "Footer":"shared/_footer.tpl",
                "HtmlHead":"index.htmlhead.tpl",
                "Layout":"shared/_layout.tpl"
            }
        ]
}
*/


type ViewStartType struct {
	Name string
	SharedHead string
	Header 	string
	Footer 	string
	HtmlHead string
	Layout 	string
}

 
type ViewStartConfigJsonObject struct {
	ViewStartTypes   []ViewStartType
}

type ViewStartConfigs struct{
	Config ViewStartConfigJsonObject
	ViewStartConfigMap  map[string]*ViewStartType
}

var TheViewStartConfigs = &ViewStartConfigs{

}

func init() {
	fmt.Println("configuration.Init")
	TheViewStartConfigs.ViewStartConfigMap = map[string]*ViewStartType{
	}
}


func (this *ViewStartConfigs) loadJsonConfig() {
	file, e := ioutil.ReadFile("conf/ViewStartConfig.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))
	json.Unmarshal(file, &this.Config)
	fmt.Printf("Results: %v\n", this.Config)

}

func (this *ViewStartConfigs) Load() {
	fmt.Printf("ViewStart Configuration Load\n")
	this.loadJsonConfig()

	for _,item:= range this.Config.ViewStartTypes{
		this.ViewStartConfigMap[item.Name] = &item
		fmt.Println(fmt.Sprintf("item: %v",item))
	}
	fmt.Printf("ViewStartConfigLookup: %v\n", TheViewStartConfigs.ViewStartConfigMap)
}
