package setting

type Config struct {
	Mysql MysqlSetting `mapstructure:"mysql"`
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