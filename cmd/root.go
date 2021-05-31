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

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

/*
 * Cluster configuration structure
 */
type Cluster struct {
	CertificateAuthorityData string `yaml:"certificate-authority-data,omitempty"`
	Server                   string `yaml:"server,omitempty"`
}

type Clusters []struct {
	Cluster Cluster `yaml:"cluster,omitempty"`
	Name    string  `yaml:"name,omitempty"`
}

/*
 * Context configuration structure
 */
type Context struct {
	Cluster   string `yaml:"cluster,omitempty"`
	Namespace string `yaml:"namespace,omitempty"`
	User      string `yaml:"user,omitempty"`
}

type Contexts []struct {
	Context Context `yaml:"context,omitempty"`
	Name    string  `yaml:"name,omitempty"`
}

/*
 * Users configuration structure
 */
type UsersConfig struct {
	ClientId     string `yaml:"client-id,omitempty"`
	ClientSecret string `yaml:"client-secret,omitempty"`
	IdToken      string `yaml:"id-token,omitempty"`
	IdpIssueURL  string `yaml:"idp-issuer-url,omitempty"`
	RefreshToken string `yaml:"refresh-token,omitempty"`
}

type UsersAuthProvider struct {
	Config UsersConfig `yaml:"config,omitempty"`
	Name   string      `yaml:"name,omitempty"`
}

type User struct {
	AuthProvider UsersAuthProvider `yaml:"auth-provider,omitempty"`
}

type Users []struct {
	Name string `yaml:"name,omitempty"`
	User User   `yaml:"user,omitempty"`
}
type KubeConfig struct {
	ApiVersion     string   `yaml:"apiVersion,omitempty"`
	Kind           string   `yaml:"kind,omitempty"`
	Preferences    struct{} `yaml:"preferences,omitempty"`
	Clusters       Clusters `yaml:"clusters,omitempty"`
	Contexts       Contexts `yaml:"contexts,omitempty"`
	CurrentContext string   `yaml:"current-context,omitempty"`
	Users          Users    `yaml:"users,omitempty"`
}

func cleanContextName(kubeConfig *KubeConfig) {
	contextIdIndex := strings.Index(kubeConfig.Contexts[0].Name, "/")

	if contextIdIndex != -1 {
		kubeConfig.Contexts[0].Name = kubeConfig.Contexts[0].Name[0:contextIdIndex]
	}
}

func cleanCurrentContext(kubeConfig *KubeConfig) {
	contextIdIndex := strings.Index(kubeConfig.CurrentContext, "/")

	if contextIdIndex != -1 {
		kubeConfig.CurrentContext = kubeConfig.CurrentContext[0:contextIdIndex]
	}
}

var rootCmd = &cobra.Command{
	Use:   "cat your_iks_kubeconfig.yaml|ikscc",
	Short: "IKS context cleaner",
	Long: `Small utility to clean the IBMCloud IKS context names to make easier to use them with a toolchain for humans
examples and usage of using your application. For example:

cat $HOME/.kube/config | ikscc

`,
	Run: func(cmd *cobra.Command, args []string) {
		fileContent, _ := ioutil.ReadAll(os.Stdin)
		kubeConfig := &KubeConfig{}

		err := yaml.Unmarshal([]byte(fileContent), &kubeConfig)
		if err != nil {
			log.Fatalf("error: %v", err)
		} else {
			cleanContextName(kubeConfig)
			cleanCurrentContext(kubeConfig)

			yamlOutput, err := yaml.Marshal(&kubeConfig)
			if err != nil {
				log.Fatalf("Ops! There was an error: %v", err)
			} else {
				fmt.Print(string(yamlOutput))
			}
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
