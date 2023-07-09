package main

import (
    "fmt"
    "flag"
    "strconv"
    "regexp"
    "os"
    "os/exec"

    "github.com/luckylat/cp-template/srcModify"
)

func check(e error){
    if e != nil {
        panic(e)
    }
}

func toLine (s string) []string {
    re := regexp.MustCompile(`\n`)
    return re.Split(s, -1)
}

func toOneline (src []string) string {
    var ret string
    for _, line := range src {
        ret = ret + line + "\n"
    }
    return ret
}

//arg[0] = Folder Name
//arg[1] = File Number (A.cpp ... X.cpp)
func main(){
    flag.Parse()
    arg := flag.Args()
    
    if len(arg) != 2 {
        fmt.Println("Missing Arguments. Example: cp-template `FolderName` `Number of Problems`")
        os.Exit(1)
    }

    //get template file
    //oj-bundle 
    templatePath := os.Getenv("CP_Template")
    lawTemplate, err := exec.Command("oj-bundle", templatePath).Output()


    eachlineTemplate := toLine(string(lawTemplate))
    modifyTemplate := srcmodify.DeleteLine(eachlineTemplate)
    template := toOneline(modifyTemplate)
    check(err)

    err = os.Mkdir(arg[0], 0755)
    check(err)

    fileNum, err := strconv.Atoi(arg[1])
    check(err)
    fileName := 'A'
    for i := 0; fileNum > i; i++ {
        file, err := os.Create(arg[0] + "/" + string(fileName) + ".cpp")
        check(err)
        file.WriteString(template)
        file.Close()
        fileName++
    }
}
