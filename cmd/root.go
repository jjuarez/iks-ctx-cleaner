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

func loadFile(fileName string) ([]byte, error) {
	var err error
	var fileContent []byte

	// The parameter should be a valid filename
	if fileName == "-" {
		// We have to deal with the standard input
		if fileContent, err = ioutil.ReadAll(os.Stdin); err != nil {
			return nil, fmt.Errorf("something went wrong reading from stdin")
		}

		return fileContent, nil
	}

	if _, err := os.Stat(fileName); err != nil {
		return nil, fmt.Errorf("The file: %s does not exists", fileName)
	}

	if fileContent, err = ioutil.ReadFile(fileName); err != nil {
		return nil, fmt.Errorf("something went wrong reading from file: %s", fileName)
	}

	return fileContent, nil
}

func unmarshalYAML(fileContent []byte) (*model.KubeConfig, error) {
	kubeConfig := model.KubeConfig{}
	err := yaml.Unmarshal(fileContent, &kubeConfig)
	if err != nil {
		return nil, err
	}

	return &kubeConfig, nil
}

func cleanupYAML(kubeConfig model.KubeConfig) model.KubeConfig {
	cleanKubeConfig := kubeConfig

	for i := 0; i < len(kubeConfig.Contexts); i++ {
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

func marshalYAML(kubeConfig model.KubeConfig) (string, error) {
	output, err := yaml.Marshal(kubeConfig)
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func exit(err error, exitCode codes.Code) {
	log.Println(err)
	os.Exit(int(exitCode))
}

var inputFile string

var rootCmd = &cobra.Command{
	Use:   "cat your_iks_kubeconfig.yaml|ikscc",
	Short: "IKS context cleaner",
	Long:  "Small utility to clean the IBMCloud IKS kubeconfig context names",
	Run: func(cmd *cobra.Command, args []string) {
		fileContent, err := loadFile(inputFile)
		if err != nil {
			exit(err, codes.ReadError)
		}

		kubeConfig, err := unmarshalYAML(fileContent)
		if err != nil {
			exit(err, codes.UnmarshalError)
		}

		cleanKubeConfig := cleanupYAML(*kubeConfig)
		yamlOutput, err := marshalYAML(cleanKubeConfig)
		if err != nil {
			exit(err, codes.MarshalError)
		}
		fmt.Print(yamlOutput)
	},
}

// Execute an entrypoint (TBC)
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&inputFile, "file", "f", "", "YAML k8s context file to clean")

	// Streams
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	// Adds the subcommands
	rootCmd.AddCommand(versionCmd)
}

func initConfig() {
	viper.AutomaticEnv()
}
