package elastic

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func NewClient(conf Config) *elasticsearch.Client {
	var address []string
	address = append(address, conf.host+":"+conf.port)

	elasticConf := elasticsearch.Config{
		Addresses: address,
	}

	elasticClient, err := elasticsearch.NewClient(elasticConf)
	if err != nil {
		log.Println("Error in creating client")
	}

	return elasticClient
}

type Config struct {
	host string
	port string
}
