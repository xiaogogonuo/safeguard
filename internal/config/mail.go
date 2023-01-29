package config

type Mail struct {
	From    string   `yaml:"from"`
	To      []string `yaml:"to"`
	Cc      []string `yaml:"cc"`
	Service `yaml:"service"`
	Auth    `yaml:"auth"`
}

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
