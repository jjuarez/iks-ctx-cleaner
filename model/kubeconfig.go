package model

//
// Cluster configuration structure
//
type Cluster struct {
	CertificateAuthorityData string `yaml:"certificate-authority-data,omitempty"`
	Server                   string `yaml:"server,omitempty"`
}

type Clusters []struct {
	Cluster Cluster `yaml:"cluster,omitempty"`
	Name    string  `yaml:"name,omitempty"`
}

//
// Context configuration structure
//
type Context struct {
	Cluster   string `yaml:"cluster,omitempty"`
	Namespace string `yaml:"namespace,omitempty"`
	User      string `yaml:"user,omitempty"`
}

type Contexts []struct {
	Context Context `yaml:"context,omitempty"`
	Name    string  `yaml:"name,omitempty"`
}

//
// Users configuration structure
//
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

//
// KubeConfig structure
//
type KubeConfig struct {
	ApiVersion     string   `yaml:"apiVersion,omitempty"`
	Kind           string   `yaml:"kind,omitempty"`
	Preferences    struct{} `yaml:"preferences,omitempty"`
	Clusters       Clusters `yaml:"clusters,omitempty"`
	Contexts       Contexts `yaml:"contexts,omitempty"`
	CurrentContext string   `yaml:"current-context,omitempty"`
	Users          Users    `yaml:"users,omitempty"`
}
