package cfg

type Config struct {
	Debug bool `yaml:"debug"`

	HostAddress    string `yaml:"host_address"`
	AuthAddress    string `yaml:"auth_address"`
	MetricsAddress string `yaml:"metrics_address"`

	SentryDSN       string `yaml:"sentry_dsn"`
	JaegerCollector string `yaml:"jaeger_collector"`

	Profiling bool `yaml:"-"`

	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSL      string `yaml:"ssl"`
	} `yaml:"db"`
}

const (
	DbUser     = "DB_USER"
	DbPassword = "DB_PASSWORD"
)

func NewConfig(yamlFile string) (*Config, error) {
	conf := &Config{}
	err := loadYaml(yamlFile, conf)
	if err != nil {
		return nil, err
	}

	//var ok bool
	//conf.DB.User, ok = os.LookupEnv(DbUser)
	//if !ok {
	//	return nil, fmt.Errorf("$" + DbUser + " isn't set")
	//}
	//
	//conf.DB.Password, ok = os.LookupEnv(DbPassword)
	//if !ok {
	//	return nil, fmt.Errorf("$" + DbPassword + " isn't set")
	//}

	return conf, nil
}
