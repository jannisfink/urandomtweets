package config

type Config struct {
	Twitter twitter
}

type twitter struct {
	ConsumerKey    string `toml:"consumer_key"`
	ConsumerSecret string `toml:"consumer_secret"`

	AccessToken  string `toml:"access_token"`
	AccessSecret string `toml:"access_secret"`
}
