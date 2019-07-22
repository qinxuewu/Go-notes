package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var configFile []byte

// 配置文件
type Config struct {
	ServerConfig     Server     `yaml:"listen"`
	SeoConfig        Seo        `yaml:"seo"`
	DatasourceConfig Datasource `yaml:"datasource"`
	UploadConfig     Upload     `yaml:"upload"`
}

type Server struct {
	Port         string `yaml:"port"`
	ReadTimeout  int    `yaml:"readtimeout"`
	WriteTimeout int    `yaml:"writetimeout"`
}

// 站点seo
type Seo struct {
	Title       string `yaml:"title"`
	Keywords    string `yaml:"keywords"`
	Description string `yaml:"description"`
}

// 数据源
type Datasource struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	User        string `yaml:"user"`
	Pwd         string `yaml:"pwd"`
	Dbname      string `yaml:"dbname"`
	Connections int    `yaml:"connections"`
	Idles       int    `yaml:"idles"`
	TablePrefix string `yaml:"table_prefix"`
}

// 文件上传配置
type Upload struct {
	ImagePrefixUrl  string   `yaml:"ImagePrefixUrl"`
	ImageSavePath   string   `yaml:"ImageSavePath"`
	ImageMaxSize    int      `yaml:"ImageMaxSize"`
	ImageAllowExts  []string `yaml:"ImageAllowExts"`
	RuntimeRootPath string   `yaml:"RuntimeRootPath"`
}

var UploadSetting = &Upload{}

// 外部方法调用
func GetConfig() (e *Config, err error) {
	err = yaml.Unmarshal(configFile, &e)
	return e, err
}

// 初始配置文件
func init() {
	var err error
	configFile, err = ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("【初始配置文件失败..】 %v ", err)
	}
}
