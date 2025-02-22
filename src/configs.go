package main

// Структуры для работы с файлом opencd.yaml.
type OpenCd struct {
	Environments []Environments `yaml:"environments"`
	Settings     Settings       `yaml:"settings"`
}

// Раздел environments: в файле opencd.yaml
type Environments struct {
	Name   string `yaml:"name"`
	Local  string `yaml:"local"`
	Remote string `yaml:"remote"`
	Docker string `yaml:"docker"`
}

// Раздел environments: в файле opencd.yaml
type Settings struct {
	Cache struct {
		Delete bool `yaml:"delete"`
	}
}
