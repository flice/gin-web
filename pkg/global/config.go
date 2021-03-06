package global

import "go.uber.org/zap/zapcore"

// 系统配置, 配置字段可参见yml注释
// viper内置了mapstructure, yml文件用"-"区分单词, 转为驼峰方便
type Configuration struct {
	System    SystemConfiguration    `mapstructure:"system" json:"system"`
	Logs      LogsConfiguration      `mapstructure:"logs" json:"logs"`
	Mysql     MysqlConfiguration     `mapstructure:"mysql" json:"mysql"`
	Redis     RedisConfiguration     `mapstructure:"redis" json:"redis"`
	Casbin    CasbinConfiguration    `mapstructure:"casbin" json:"casbin"`
	Jwt       JwtConfiguration       `mapstructure:"jwt" json:"jwt"`
	RateLimit RateLimitConfiguration `mapstructure:"rate-limit" json:"rateLimit"`
	Upload    UploadConfiguration    `mapstructure:"upload" json:"upload"`
}

type SystemConfiguration struct {
	UrlPathPrefix   string `mapstructure:"url-path-prefix" json:"urlPathPrefix"`
	Port            int    `mapstructure:"port" json:"port"`
	UseRedis        bool   `mapstructure:"use-redis" json:"useRedis"`
	Transaction     bool   `mapstructure:"transaction" json:"transaction"`
	InitData        bool   `mapstructure:"init-data" json:"initData"`
	OperationLogKey string `mapstructure:"operation-log-key" json:"operationLogKey"`
	RSAPublicKey    string `mapstructure:"rsa-public-key" json:"rsaPublicKey"`
	RSAPrivateKey   string `mapstructure:"rsa-private-key" json:"rsaPrivateKey"`
	RSAPublicBytes  []byte `mapstructure:"-" json:"-"`
	RSAPrivateBytes []byte `mapstructure:"-" json:"-"`
}

type LogsConfiguration struct {
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"maxSize"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}

type MysqlConfiguration struct {
	Username    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	Database    string `mapstructure:"database" json:"database"`
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Query       string `mapstructure:"query" json:"query"`
	LogMode     bool   `mapstructure:"log-mode" json:"logMode"`
	TablePrefix string `mapstructure:"table-prefix" json:"tablePrefix"`
	Charset     string `mapstructure:"charset" json:"charset"`
	Collation   string `mapstructure:"collation" json:"collation"`
}

type RedisConfiguration struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	Password  string `mapstructure:"password" json:"password"`
	Database  int    `mapstructure:"database" json:"database"`
	BinlogPos string `mapstructure:"binlog-pos" json:"binlogPos"`
}

type CasbinConfiguration struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath"`
}

type JwtConfiguration struct {
	Realm      string `mapstructure:"realm" json:"realm"`
	Key        string `mapstructure:"key" json:"key"`
	Timeout    int    `mapstructure:"timeout" json:"timeout"`
	MaxRefresh int    `mapstructure:"max-refresh" json:"maxRefresh"`
}

type RateLimitConfiguration struct {
	Max int64 `mapstructure:"max" json:"max"`
}

type UploadConfiguration struct {
	SaveDir                      string `mapstructure:"save-dir" json:"saveDir"`
	SingleMaxSize                uint   `mapstructure:"single-max-size" json:"singleMaxSize"`
	MergeConcurrentCount         uint   `mapstructure:"merge-concurrent-count" json:"mergeConcurrentCount"`
	CompressImageCronTask        string `mapstructure:"compress-image-cron-task" json:"compressImageCronTask"`
	CompressImageRootDir         string `mapstructure:"compress-image-root-dir" json:"compressImageRootDir"`
	CompressImageOriginalSaveDir string `mapstructure:"compress-image-original-save-dir" json:"compressImageOriginalSaveDir"`
}
