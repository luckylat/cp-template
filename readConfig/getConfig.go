package readConfig

import (
	"fmt"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"
)


func GetTemplatePath(templateStyle string) (string, error) {
    configFolder := func() string {
        if(os.Getenv("XDG_CONFIG_HOME") != ""){
            return os.Getenv("XDG_CONFIG_HOME")
        }else{
            return os.Getenv("HOME") + "/.config"
        }
    }()
    configFilePath := configFolder + "/cp-template/settings.toml"
    
    fp, err := os.Open(configFilePath)

    defer fp.Close()

    buf := make([]byte, 1024)
    var configByteFile []byte
    for {
        n, err := fp.Read(buf)

        if n == 0 { break }
        if err != nil { return "", err }
        configByteFile = append(configByteFile, buf[:n]...)

    }

    var tomlData map[string]interface{}
    err = toml.Unmarshal(configByteFile, &tomlData)
    templates, ok := tomlData["templates"].(map[string]interface{})
    if !ok {
        return "", fmt.Errorf("toml parsing error")
    }
    if err != nil {
        fmt.Printf("%s\n",err)
        return "", err
    }
    if templates[templateStyle] == "" {
        return "", fmt.Errorf("a template of the template style is nothing")
    }
    return templates[templateStyle].(string), nil
}
