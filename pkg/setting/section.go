package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Mysql MysqlSetting `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
}
type ServerSetting struct {
	Port int `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MysqlSetting struct {
	Host	 string `mapstructure:"host"`
	Port	 string `mapstructure:"port"`
	User	 string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName	 string `mapstructure:"dbname"`
	MaxIdleConns	 int `mapstructure:"maxIdleConns"`
	MaxOpenConns	 int `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime	 int `mapstructure:"connMaxLifeTime"`
}

type LoggerSetting struct {
	Log_level string `mapstructure:"log_level"`
	File_log_path string `mapstructure:"file_log_path"`
	Max_size int `mapstructure:"max_size"`
	Max_backups int `mapstructure:"max_backups"`
	Max_age int `mapstructure:"max_age"`
	Compress bool `mapstructure:"compress"`
}