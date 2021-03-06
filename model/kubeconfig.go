package model

//
// Cluster configuration structure
//
type Cluster struct {
	CertificateAuthorityData string `yaml:"certificate-authority-data,omitempty"`
	Server                   string `yaml:"server,omitempty"`
}

//
// Clusters the collection of the cluster included in a context
//
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
	AuthProvider          UsersAuthProvider `yaml:"auth-provider,omitempty"`
	ClientCertificateDAta string            `yaml:"client-certificate-data,omitempty"`
	ClientKeyData         string            `yaml:"client-key-data,omitempty"`
}

type Users []struct {
	Name string `yaml:"name,omitempty"`
	User User   `yaml:"user,omitempty"`
}

type Preferences struct {
}

//
// KubeConfig structure
//
type KubeConfig struct {
	ApiVersion     string      `yaml:"apiVersion"`
	Kind           string      `yaml:"kind"`
	Preferences    Preferences `yaml:"preferences,omitempty"`
	Clusters       Clusters    `yaml:"clusters"`
	Contexts       Contexts    `yaml:"contexts"`
	CurrentContext string      `yaml:"current-context,omitempty"`
	Users          Users       `yaml:"users,omitempty"`
}
