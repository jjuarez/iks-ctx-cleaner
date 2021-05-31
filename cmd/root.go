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

func loadYAML(fileName string, kubeConfig *model.KubeConfig) {
	var fileContent []byte
	var fileError error

	if fileName == "-" || fileName == "" {
		fileContent, fileError = ioutil.ReadAll(os.Stdin)
		if fileError != nil {
			exit("Something went wrong reading from stdin", errors.ReadError)
		}
	} else {
		fileContent, fileError = ioutil.ReadFile(fileName)
		if fileError != nil {
			exit(fmt.Sprintf("Something went wrong reading from file: %s", fileName), errors.ReadError)
		}
	}

	err := yaml.Unmarshal(fileContent, &kubeConfig)
	if err != nil {
		exit(fmt.Sprintf("Error: %v", err), errors.UnmarshalError)
	}
}

func cleanUp(kubeConfig *model.KubeConfig) {
	maxContexts := len(kubeConfig.Contexts)

	for i := 0; i < maxContexts; i++ {
		idx := strings.Index(kubeConfig.Contexts[i].Name, "/")

		if idx != -1 {
			kubeConfig.Contexts[i].Name = kubeConfig.Contexts[i].Name[0:idx]
		}
	}

	if kubeConfig.CurrentContext != "" {
		idx := strings.Index(kubeConfig.CurrentContext, "/")

		if idx != -1 {
			kubeConfig.CurrentContext = kubeConfig.CurrentContext[0:idx]
		}
	}
}

func showYAML(kubeConfig model.KubeConfig) {
	output, err := yaml.Marshal(kubeConfig)
	if err != nil {
		exit(fmt.Sprintf("Error: %v", err), errors.MarshalError)
	}

	fmt.Print(string(output))
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
		kubeConfig := model.KubeConfig{}

		loadYAML(InputFile, &kubeConfig)
		cleanUp(&kubeConfig)
		showYAML(kubeConfig)
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
