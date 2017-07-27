package config

type Config struct {
	App   App   `toml:"app"`
	Cache Cache `toml:"cache"`
	Asset Asset `toml:"asset"`
}

type App struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
	Port    int16  `toml:"port"`
}

type Cache struct {
	Host    string `toml:"host"`
	Port    int16  `toml:"port"`
	Enabled bool   `toml:"enabled"`
}

type Asset struct {
	BaseUrl string `toml:"base_url"`
	CSSPath string `toml:"css_path"`
	JSPath  string `toml:"js_path"`
	ImgPath string `toml:"img_path"`
}
