/*
Copyright © 2021 Javier Juarez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jjuarez/iks-ctx-cleaner/errors"
	"github.com/jjuarez/iks-ctx-cleaner/model"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func readFile(fileName string) []byte {
	if fileName == "-" || fileName == "" {
		fileContent, fileError := ioutil.ReadAll(os.Stdin)
		if fileError != nil {
			exit("Something went wrong reading from stdin", errors.ReadError)
		}

		return fileContent
	} else {
		fileContent, fileError := ioutil.ReadFile(fileName)
		if fileError != nil {
			exit(fmt.Sprintf("Something went wrong reading from file: %s", fileName), errors.ReadError)
		}

		return fileContent
	}
}

func cleanUp(kubeConfig *model.KubeConfig) {
	for i := 0; i < len(kubeConfig.Contexts); i++ {
		contextIdIndex := strings.Index(kubeConfig.Contexts[i].Name, "/")

		if contextIdIndex != -1 {
			kubeConfig.Contexts[i].Name = kubeConfig.Contexts[i].Name[0:contextIdIndex]
		}
	}

	if kubeConfig.CurrentContext != "" {
		contextIdIndex := strings.Index(kubeConfig.CurrentContext, "/")

		if contextIdIndex != -1 {
			kubeConfig.CurrentContext = kubeConfig.CurrentContext[0:contextIdIndex]
		}
	}
}

func exit(message string, exitCode errors.Code) {
	log.Println(message)
	os.Exit(int(exitCode))
}

var InputFile string

var rootCmd = &cobra.Command{
	Use:   "cat your_iks_kubeconfig.yaml|ikscc",
	Short: "IKS context cleaner",
	Long:  "Small utility to clean the IBMCloud IKS kubeconfig context names",
	Run: func(cmd *cobra.Command, args []string) {
		fileContent := readFile(InputFile)
		kubeConfig := &model.KubeConfig{}

		err := yaml.Unmarshal(fileContent, &kubeConfig)
		if err != nil {
			exit(fmt.Sprintf("Error: %v", err), errors.UnmarshalError)
		}

		cleanUp(kubeConfig)

		cleanYaml, err := yaml.Marshal(&kubeConfig)
		if err != nil {
			exit(fmt.Sprintf("Error: %v", err), errors.MarshalError)
		}

		fmt.Print(string(cleanYaml))
		os.Exit(0)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&InputFile, "file", "f", "-", "Input file to clean")
}

func initConfig() {
	viper.AutomaticEnv()
}
