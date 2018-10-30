package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/gokaykucuk/quotecommit/data"
	"go/build"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"strings"
)

func main() {
	if len(os.Args) > 1 &&  os.Args[1] == "update_data"{
		// This should only go run environment in projects source folder.
		updateData()
		return
	}
	pushAfterCommit := false
	if len(os.Args) > 1 &&  os.Args[1] == "p"{
		pushAfterCommit = true
	}
	chosen_message := getRandomQuote()
	splitted_message := strings.Split(chosen_message, "|")
	color.Green("Adding and commiting your files with commit message...\n")
	fmt.Println(splitted_message[0])
	color.New(color.Bold, color.BgCyan, color.Italic).Println("      " + splitted_message[1]+ "   ")
	// I like how it does warn you..
	commit_cmd := exec.Command("git", "commit", "-a", "-m", getRandomQuote())
	commit_cmd.Run()
	fmt.Println()
	if pushAfterCommit{
		color.New(color.BgHiBlue).Print("! p flag found! Also pushing...\n")
		push_cmd := exec.Command("git", "push")
		push_cmd.Run()
		color.New(color.FgYellow, color.Bold).Println("Done!")
	}
	color.New(color.BlinkSlow,color.FgRed).Println("!!! Please do this as rarely as possible !!!")
}

func getRandomQuote() (quote string){
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(data.Data)
	quote = data.Data[n]
	return
}

// Build time function, not so important!
func updateData(){
	// Make this able to process all files in folder
	text, err := ioutil.ReadFile(projectFolder()+"/data/data.raw")
	// Process the text so it represents an array
	scanner := bufio.NewScanner(strings.NewReader(string(text)))
	var processed_text string
	for scanner.Scan(){
		processed_text += "\""+scanner.Text()+"\",\n"
	}
	go_data_template, err := ioutil.ReadFile(projectFolder()+"/data/template.go")
	go_data_replaced := strings.Replace(string(go_data_template),"//+",processed_text, -1)
	go_data_replaced2 := strings.Replace(go_data_replaced,"template_data","Data", -1)
	err = ioutil.WriteFile(projectFolder()+"/data/data.go", []byte(go_data_replaced2),0)
	if err != nil{
		panic(err)
	}
	fmt.Println("data/data.go updated!")
}
// Build time function, not important!
func projectFolder()(gopath string){
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	gopath += "/src/"
	if os.Getenv("PROJECT_URI") != ""{
		gopath += os.Getenv("PROJECT_URI")
	}else{
		gopath += "github.com/gokaykucuk/quotecommit"
	}
	return
}