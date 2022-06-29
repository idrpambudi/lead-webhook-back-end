package dependency

import (
	"fmt"

	"leadwebhook/cfg"
	"leadwebhook/pkg/repository"
	mysqlrepo "leadwebhook/pkg/repository/mysql"
	"leadwebhook/pkg/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ProvideConfig() cfg.Configurations {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./cfg")
	viper.AutomaticEnv()

	var config cfg.Configurations
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading config file, %v", err))
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	fmt.Println(config)
	return config
}

func ProvideLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	slogger := logger.Sugar()
	return slogger
}

func ProvideMySqlClient(config cfg.Configurations) *gorm.DB {
	client, err := gorm.Open(mysql.Open(config.Database.MySql.URL))
	if err != nil {
		panic(err)
	}
	return client
}

func ProvideLeadRepository(client *gorm.DB) repository.LeadRepository {
	return mysqlrepo.NewMySqlLeadRepository(client)
}

func ProvideLeadService(leadRepo repository.LeadRepository) *service.LeadService {
	return service.NewLeadService(leadRepo)
}
