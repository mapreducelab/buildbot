package models

// A Component represent smallest working unit.
type Component struct {
	ID     string `yaml:"id"`
	GitUrl string `yaml:"git_url"`
}
