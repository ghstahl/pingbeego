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
                "TplNames":"index.tpl",
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
	TplNames string
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
	ViewStartConfigLookup  map[string]string
}

var TheViewStartConfigs = &ViewStartConfigs{

}

func init() {
	fmt.Println("configuration.Init")
	TheViewStartConfigs.ViewStartConfigLookup = map[string]string{
		"a":"a",
		"b":"b",
		"c":"c",
	}
	fmt.Printf("ViewStartConfigLookup: %v\n", TheViewStartConfigs.ViewStartConfigLookup)
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
	this.loadJsonConfig()
	fmt.Println(fmt.Sprintf("item: %v",this.ViewStartConfigLookup))
	for _,item:= range this.Config.ViewStartTypes{
		fmt.Println(fmt.Sprintf("item: %v",item))
	}
}
