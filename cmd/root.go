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
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type KubeConfig struct {
	ApiVersion  string   `yaml:"apiVersion"`
	Kind        string   `yaml:"kind"`
	Preferences struct{} `yaml:"preferences,omitempty"`
	Clusters    []struct {
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
			Server                   string `yaml:"server"`
		}
		Name string `yaml:"name"`
	}
	Contexts []struct {
		Context struct {
			Cluster   string `yaml:"cluster"`
			Namespace string `yaml:"namespace"`
			User      string `yaml:"user"`
		}
		Name string `yaml:"name"`
	}
	CurrentContext string `yaml:"current-context"`
}

func dumpData(kubeConfig *KubeConfig) {
	log.Printf("Cluster.name: %s", kubeConfig.Clusters[0].Name)
	log.Printf("Cluster.certificate-authority-data: %s", kubeConfig.Clusters[0].Cluster.CertificateAuthorityData)
	log.Printf("Cluster.server: %s", kubeConfig.Clusters[0].Cluster.Server)
	log.Printf("Context.name: %s", kubeConfig.Contexts[0].Name)
	log.Printf("Context.cluster: %s", kubeConfig.Contexts[0].Context.Cluster)
	log.Printf("Context.namespace: %s", kubeConfig.Contexts[0].Context.Namespace)
	log.Printf("Context.user: %s", kubeConfig.Contexts[0].Context.User)
}

var rootCmd = &cobra.Command{
	Use:   "iksctxcleaner",
	Short: "IKS context cleaner",
	Long: `Small utility to clean the IBMCloud k8s context names to make easier to use them with a toolchain for humans
examples and usage of using your application. For example:

cat $HOME/.kube/config | iksctxcleaner

`,
	Run: func(cmd *cobra.Command, args []string) {
		fileContent, _ := ioutil.ReadAll(os.Stdin)
		kubeConfig := &KubeConfig{}

		err := yaml.Unmarshal([]byte(fileContent), &kubeConfig)
		if err != nil {
			log.Fatalf("error: %v", err)
		} else {
			dumpData(kubeConfig)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv()
}
