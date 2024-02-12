package main

import (
    "fmt"
    "flag"
    "regexp"
    "os"
    "os/exec"

    "github.com/luckylat/cp-template/srcModify"
    "github.com/luckylat/cp-template/readConfig"
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
    flag.CommandLine.Usage = func() {
        o := flag.CommandLine.Output()
        fmt.Fprintf(o, "\nUsage: %s\n", flag.CommandLine.Name())
        fmt.Fprintf(o, "cp-template --folder `FolderName` --number `Number of Problems` (--style `TemplateStyle`)\n")
        fmt.Fprintf(o, "docs: https://github.com/luckylat/cp-template/blob/main/README.md#how-to-use")

    }

    var (
        folder string
        problemNumber int
        templateStyle string
    )
    //enhance: set short alias command !/w duplicate command name
    flag.StringVar(&folder, "folder", "", "Folder Name")
    flag.IntVar(&problemNumber, "number", 0, "Problem Number")
    flag.StringVar(&templateStyle, "style", "default", "Template Style")

    flag.Parse()

    if folder == "" || problemNumber == 0 {
        fmt.Println("Missing Arguments. Example: cp-template --folder `FolderName` --number `Number of Problems` (--style `TemplateStyle`)")
        fmt.Println("docs: https://github.com/luckylat/cp-template/blob/main/README.md#how-to-use")
        os.Exit(1)
    }

    //get template file
    //oj-bundle 
    templatePath, err := readConfig.GetTemplatePath(templateStyle)
    check(err)
    lawTemplate, err := exec.Command("oj-bundle", templatePath).Output()


    eachlineTemplate := toLine(string(lawTemplate))
    modifyTemplate := srcmodify.DeleteLine(eachlineTemplate)
    template := toOneline(modifyTemplate)
    check(err)

    err = os.Mkdir(folder, 0755)
    check(err)

    check(err)
    fileName := 'A'
    for i := 0; problemNumber > i; i++ {
        file, err := os.Create(folder + "/" + string(fileName) + ".cpp")
        check(err)
        file.WriteString(template)
        file.Close()
        fileName++
    }
}
