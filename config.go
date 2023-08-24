package main

import (
	"fmt"
	"os"

	"github.com/d8x/sgw/providers"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Storage struct {
	Name            string
	Type            string
	StorageProvider providers.Provider
}

type MinioMeta struct {
	Endpoint string `yaml:"endpoint"`
	Bucket   string `yaml:"bucket"`
}

type MinioAuth struct {
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
}

type ScalewayMeta struct {
	Endpoint string
	Bucket   string
}

type ScalewayAuth struct {
	AccessKey string
	SecretKey string
}

type Config struct {
	ListenPort string
	Debug      bool
	Providers  map[string]Storage
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ReadConfig(configFilePath string) error {
	logrus.Debugf("config file path: %s", configFilePath)
	viper.AddConfigPath(configFilePath)
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	viper.AddConfigPath(os.Getenv("SGW_CONFIG_PATH"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("sgw")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	logrus.Infof("using config file: %s", viper.ConfigFileUsed())

	providersS3 := viper.GetStringMap("storages.providers")

	storageProviders := make(map[string]Storage, len(providersS3))

	for name := range providersS3 {
		if !providers.IsSupported(providers.ProviderType(viper.GetString("storages.providers." + name + ".type"))) {
			logrus.Errorf("provider %s is not supported", name)
			continue
		}
		provider := Storage{
			Name: name,
			Type: viper.GetString("storages.providers." + name + ".type"),
		}
		var err error
		switch provider.Type {
		case "minio":
			auth := MinioAuth{}
			meta := MinioMeta{}
			if err := viper.UnmarshalKey("storages.providers."+name+".auth", &auth); err != nil {
				return fmt.Errorf("failed to parse auth for provider %s: %v", name, err)
			}

			if err := viper.UnmarshalKey("storages.providers."+name+".meta", &meta); err != nil {
				return fmt.Errorf("failed to parse meta for provider %s: %v", name, err)
			}

			provider.StorageProvider, err = providers.NewMinio(meta.Endpoint, meta.Bucket, auth.AccessKey, auth.SecretKey)
			if err != nil {
				return fmt.Errorf("failed to create minio provider %s: %v", name, err)
			}
		case "scaleway":
			auth := ScalewayAuth{}
			meta := ScalewayMeta{}
			if err := viper.UnmarshalKey("storages.providers."+name+".auth", &auth); err != nil {
				return fmt.Errorf("failed to parse auth for provider %s: %v", name, err)
			}
			if err := viper.UnmarshalKey("storages.providers."+name+".meta", &meta); err != nil {
				return fmt.Errorf("failed to parse meta for provider %s: %v", name, err)
			}

			provider.StorageProvider, err = providers.NewScaleway(meta.Endpoint, meta.Bucket, auth.AccessKey, auth.SecretKey)
			if err != nil {
				return fmt.Errorf("failed to create scaleway provider %s: %v", name, err)
			}
		}

		storageProviders[name] = provider
	}

	c.ListenPort = viper.GetString("listenPort")
	c.Debug = viper.GetBool("debug")
	c.Providers = storageProviders

	return nil
}

func (c *Config) WatchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("config file changed: %s", e.Name)
		if err := c.ReadConfig(viper.ConfigFileUsed()); err != nil {
			logrus.Errorf("failed to reload config: %v", err)
		}
	})
}
