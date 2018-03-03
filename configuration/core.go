package configuration

var Configs []Configuration

type Configuration struct {
	Location  string
	Teamwork  TeamworkConfig   `mapstructure:"teamwork"`
	Favorites []FavoriteConfig `mapstructure:"favorites"`
}

type TeamworkConfig struct {
	SiteName string `mapstructure:"siteName"`
	ApiKey   string `mapstructure:"apiKey"`
}

type FavoriteConfig struct {
	Name   string `mapstructure:"name"`
	TaskId string `mapstructure:"taskId"`
}

const FILENAME = ".teamworkgo"


// TODO write getters and setters