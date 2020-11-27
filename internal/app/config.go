package app

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Configs struct {
	vn    *viper.Viper
	State string

	Port    int     `mapstructure:"port"`
	MongoDB MongoDB `mapstructure:"mongodb"`
	Movie   Movie   `mapstructure:"movie"`

	BangkokLocation *time.Location
}

func NewConfig() *Configs {
	return &Configs{}
}

func (c *Configs) Init(s string) error {
	if s == "" {
		s = "dev"
	}
	name := fmt.Sprintf("config.%s", s)

	vn := viper.New()
	vn.AddConfigPath("./config")
	vn.SetConfigName(name)
	c.vn = vn
	c.State = s

	if err := vn.ReadInConfig(); err != nil {
		return err
	}

	if err := c.binding(); err != nil {
		return err
	}
	return nil
}

func (c *Configs) binding() error {
	if err := c.vn.Unmarshal(&c); err != nil {
		return err
	}

	// binding additional data
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return err
	}
	c.BangkokLocation = loc

	if err := c.MongoDB.binding(); err != nil {
		return err
	}

	return nil
}

type MongoDB struct {
	ConnectionString string `mapstructure:"conn"`

	Client *mongo.Client
}

func (m *MongoDB) binding() error {
	ctx := context.Background()
	client, err := mongo.NewClient(options.Client().ApplyURI(m.ConnectionString))
	if err != nil {
		return err
	}
	if err := client.Connect(ctx); err != nil {
		return err
	}

	// check connection
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}
	m.Client = client

	return nil
}

type Movie struct {
	Enable bool `mapstructure:"enable"`
}
