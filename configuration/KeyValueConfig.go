package configuration

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

/*
{
    "KeyValues":
    [
        {
            "Key":"key1",
            "Value":"value1"
        },
        {
            "Key":"key2",
            "Value":"value2"
        },
        {
            "Key":"key3",
            "Value":"value3"
        }
    ]
}
*/


type KeyValueType struct {
	Key string
	Value 	string
}

type KeyValueConfigJsonObject struct {
	KeyValues   []KeyValueType
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
	fmt.Printf("KeyValueConfigs Configuration Load\n")
	this.loadJsonConfig()
	fmt.Println(fmt.Sprintf("item: %v",this.KeyValueConfigLookup))
	for _,item:= range this.Config.KeyValues{
		fmt.Println(fmt.Sprintf("item: %v",item))
	}
}
