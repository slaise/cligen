/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var name string
var version string
var help string
var author string

// kubectlgenCmd represents the kubectlgen command
var kubectlgenCmd = &cobra.Command{
	Use:   "kubectlgen",
	Short: "kubectlgen is a code generator for a kubectl plugin with client-go",
	Long: `kubectlgen generates code with a cobra command and client-go setup 
	to speed up the kubectl plugin development.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(name) < 4 {
			log.Fatalln("Failed to create the code, because name is too short. length(name) >= 4")
		}
		log.Println("Create directories ...")
		if err := os.MkdirAll(name+"/cmd", os.ModePerm); err != nil {
			log.Fatalf("Failed to create %s dir, %v", name, err)
		}
		if err := os.MkdirAll(name+"/pkg/client", os.ModePerm); err != nil {
			log.Fatalf("Failed to create %s/pkg/client dir, %v", name, err)
		}
		if err := os.MkdirAll(name+"/pkg/config", os.ModePerm); err != nil {
			log.Fatalf("Failed to create %s/pkg/config dir, %v", name, err)
		}
		log.Println("Generate go.mod file ...")
		templateGenerate("tpl/go.mod.tpl", name+"/go.mod", map[string]string{"name": name})
		log.Println("Generate main.go ...")
		templateGenerate("tpl/main.go.tpl", name+"/main.go", map[string]string{"name": name, "author": author})
		log.Println("Generate version.go ...")
		templateGenerate("tpl/cmd/version.go.tpl", name+"/cmd/version.go", map[string]string{"name": name, "version": version})
		if len(help) > 0 {
			log.Println("Generate cmd/help.go ...")
			templateGenerate("tpl/cmd/help.go.tpl", name+"/cmd/help.go", map[string]string{"name": name, "help": help, "version": version})
		}
		log.Println("Generate plugincmd.go ...")
		templateGenerate("tpl/cmd/plugincmd.go.tpl", name+"/cmd/plugincmd.go", map[string]string{"name": name})
		log.Println("Generate client.go ...")
		templateGenerate("tpl/pkg/client/client.go.tpl", name+"/pkg/client/client.go", map[string]string{})
		log.Println("Generate config.go ...")
		templateGenerate("tpl/pkg/config/config.go.tpl", name+"/pkg/config/config.go", map[string]string{})
	},
}

func templateGenerate(templateFileName string, oFileName string, data any) {
	t, f := createFileFromTemplate(templateFileName, oFileName)
	err := t.Execute(f, data)
	if err != nil {
		log.Fatalf("Failed to parse %s template file, %v \n", oFileName, err)
	}
}

func createFileFromTemplate(templateFilename string, oFileName string) (*template.Template, *os.File) {
	modT := template.Must(template.ParseFiles(templateFilename))
	modF, err := os.Create(oFileName)
	if err != nil {
		log.Fatalf("Failed to create %s file, %v \n", oFileName, err)
	}
	return modT, modF
}

func init() {
	rootCmd.AddCommand(kubectlgenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kubectlgenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kubectlgenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	kubectlgenCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the new kubectl plugin")
	kubectlgenCmd.Flags().StringVar(&version, "v", "0.0.1", "Version of the new kubectl plugin")
	kubectlgenCmd.Flags().StringVar(&help, "hi", "", "Help info of the new kubectl plugin")
	kubectlgenCmd.Flags().StringVar(&author, "a", "", "Author of the new kubectl plugin")
}
