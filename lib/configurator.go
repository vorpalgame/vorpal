package lib

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Initial setup for configuration files with yaml.
// Will be imporant for state/scene transions.
func StartConfiguration(path string) {
	viper.SetDefault("ContentDir", "resources")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	viper.SetConfigName("config.yaml") // name of config file (without extension)
	viper.SetConfigType("yaml")        // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)          // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	//viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
}
