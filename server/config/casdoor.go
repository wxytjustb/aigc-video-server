package config

// Casdoor 登录配置
type Casdoor struct {
	Endpoint     string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	ClientID     string `mapstructure:"client-id" json:"clientId" yaml:"client-id"`
	ClientSecret string `mapstructure:"client-secret" json:"clientSecret" yaml:"client-secret"`
	//RedirectURL      string `mapstructure:"redirect-url" json:"redirectUrl" yaml:"redirect-url"`
	Certificate      string `mapstructure:"certificate" json:"certificate" yaml:"certificate"`
	OrganizationName string `mapstructure:"organization-name" json:"organizationName" yaml:"organization-name"`
	ApplicationName  string `mapstructure:"application-name" json:"applicationName" yaml:"application-name"`
}
