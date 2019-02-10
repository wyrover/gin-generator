package main

import (
    "os"
    "flag"
    "fmt"
    "regexp"
    "strings"
    "log"

    "github.com/wyrover/gin-generator/static"
)






type Config struct {
	appName      string
	modelName           string
	
}

func (c *Config) Setup() {	
    flag.StringVar(&c.appName, "a", "XXXXX", "X@@@@ name")
	flag.StringVar(&c.modelName, "m", "@@@@@", "model name")
}


func main() {    

    c := Config{}
	c.Setup()
   
    flag.Parse()

    
    content, err := static.ReadFile("template/controller_template.go")
	if err != nil {
		log.Fatal(err)
	}

    controller_template := string(content)

    content, err = static.ReadFile("template/model_template.go")
    if err != nil {
        log.Fatal(err)
    }

    model_template := string(content)

    
    re1 := regexp.MustCompile("@@@@@")
    re2 := regexp.MustCompile("X@@@@")
    re3 := regexp.MustCompile("XXXXX")
    
    controller_content := re1.ReplaceAllString(controller_template, c.modelName)
    controller_content = re2.ReplaceAllString(controller_content, strings.ToLower(c.modelName[0:1]) + c.modelName[1:])
    controller_content = re3.ReplaceAllString(controller_content, c.appName)

    fmt.Println(controller_content)

    model_content := re1.ReplaceAllString(model_template, c.modelName)
    model_content = re2.ReplaceAllString(model_content, strings.ToLower(c.modelName[0:1]) + c.modelName[1:])
    model_content = re3.ReplaceAllString(model_content, c.appName)

    os.MkdirAll("ginApp/controllers", os.ModePerm)
    os.MkdirAll("ginApp/models", os.ModePerm)

    writefile("ginApp/controllers/" + strings.ToLower(c.modelName[0:1]) + c.modelName[1:] + "_controller.go", controller_content)
    writefile("ginApp/models/" + strings.ToLower(c.modelName[0:1]) + c.modelName[1:] + ".go", model_content)
	
}

func writefile(filename string, content string) {
    f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}
}