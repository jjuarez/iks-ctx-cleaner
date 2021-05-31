package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jjuarez/iks-ctx-cleaner/codes"
	"github.com/jjuarez/iks-ctx-cleaner/model"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func loadFile(fileName string) []byte {
	var err error
	var fileContent []byte

	if fileName == "-" || fileName == "" {
		fileContent, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			exit("Something went wrong reading from stdin", codes.ReadError)
		}

		return fileContent
	} else {
		fileContent, err = ioutil.ReadFile(fileName)
		if err != nil {
			exit(fmt.Sprintf("Something went wrong reading from file: %s", fileName), codes.ReadError)
		}

		return fileContent
	}
}

func unmarshalYAML(fileContent []byte) model.KubeConfig {
	kubeConfig := model.KubeConfig{}
	err := yaml.Unmarshal(fileContent, &kubeConfig)
	if err != nil {
		exit(fmt.Sprintf("Error: %v", err), codes.UnmarshalError)
	}

	return kubeConfig
}

func cleanupYAML(kubeConfig model.KubeConfig) model.KubeConfig {
	cleanKubeConfig := kubeConfig
	maxContexts := len(kubeConfig.Contexts)

	for i := 0; i < maxContexts; i++ {
		idx := strings.Index(kubeConfig.Contexts[i].Name, "/")
		if idx != -1 {
			cleanKubeConfig.Contexts[i].Name = kubeConfig.Contexts[i].Name[0:idx]
		}
	}

	if kubeConfig.CurrentContext != "" {
		idx := strings.Index(kubeConfig.CurrentContext, "/")
		if idx != -1 {
			cleanKubeConfig.CurrentContext = kubeConfig.CurrentContext[0:idx]
		}
	}

	return cleanKubeConfig
}

func marshalYAML(kubeConfig model.KubeConfig) string {
	output, err := yaml.Marshal(kubeConfig)
	if err != nil {
		exit(fmt.Sprintf("Error: %v", err), codes.MarshalError)
	}
	return string(output)
}

func exit(message string, exitCode codes.Code) {
	log.Println(message)
	os.Exit(int(exitCode))
}

var InputFile string

var rootCmd = &cobra.Command{
	Use:   "cat your_iks_kubeconfig.yaml|ikscc",
	Short: "IKS context cleaner",
	Long:  "Small utility to clean the IBMCloud IKS kubeconfig context names",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(marshalYAML(cleanupYAML(unmarshalYAML(loadFile(InputFile)))))
		os.Exit(0)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&InputFile, "file", "f", "", "Input file to clean")
}

func initConfig() {
	viper.AutomaticEnv()
}
