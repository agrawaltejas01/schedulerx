package db

func GetDBConfig(opts *DBConfig) DBConfig {

	if opts == nil {
		opts = &DBConfig{}
	}

	if opts.Username == "" {
		opts.Username = "user"
	}
	if opts.Password == "" {
		opts.Password = "password"
	}
	if opts.Host == "" {
		opts.Host = "localhost"
	}
	if opts.Port == "" {
		opts.Port = "3306"
	}
	if opts.DatabaseName == "" {
		opts.DatabaseName = "db"
	}

	return DBConfig{
		Username:     opts.Username,
		Password:     opts.Password,
		Host:         opts.Host,
		Port:         opts.Port,
		DatabaseName: opts.DatabaseName,
	}

}
