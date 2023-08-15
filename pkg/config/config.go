package config

import (
	"encoding/json"
	"os"
)

// Config represents the configuration structure.
type Config struct {
	SourceDirectory             string `json:"source_directory"`
	OutputDirectory             string `json:"output_directory"`
	TemplateFilePath            string `json:"template_file_path"`
	AnnotationTitleConfigurable bool   `json:"annotation_title_configurable"`
}

// Load loads the configuration from a JSON file.
func Load() (*Config, error) {
	content, err := os.ReadFile("./config/config.json")
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
