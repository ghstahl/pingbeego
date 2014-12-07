package configuration

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

/*
{
    "KeyValueConfig":
    [
          {
            "key":"key1",
            "value":"value1"
        }
    ]
}
*/


type KeyValueType struct {
	Key string
	Value 	string
}

type KeyValueConfigJsonObject struct {
	KeyValueConfig   []KeyValueType
}

type KeyValueConfigs struct{
	Config KeyValueConfigJsonObject
	KeyValueConfigLookup  map[string]string
}

var TheKeyValueConfigs = &KeyValueConfigs{

}

func init() {
	fmt.Println("configuration.Init")
	TheKeyValueConfigs.KeyValueConfigLookup = map[string]string{
		"a":"a",
		"b":"b",
		"c":"c",
	}
	fmt.Printf("KeyValueConfigLookup: %v\n", TheKeyValueConfigs.KeyValueConfigLookup)
}


func (this *KeyValueConfigs) loadJsonConfig() {
	file, e := ioutil.ReadFile("conf/KeyValueConfig.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))
	json.Unmarshal(file, &this.Config)
	fmt.Printf("Results: %v\n", this.Config)
}

func (this *KeyValueConfigs) Load() {
	this.loadJsonConfig()
	fmt.Println(fmt.Sprintf("item: %v",this.KeyValueConfigLookup))
	for _,item:= range this.Config.KeyValueConfig{
		fmt.Println(fmt.Sprintf("item: %v",item))
	}
}
