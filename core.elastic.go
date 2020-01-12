package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

//ESCLIENT ES client
var ESCLIENT *elastic.Client

func connectElastic() bool {

	ctx := context.Background()

	client, err := elastic.NewClient(elastic.SetURL(ElasticURL + ":" + ElasticPort))

	if err != nil {
		panic(err)
	}

	info, code, err := client.Ping(ElasticURL + ":" + ElasticPort).Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Elasticsearch connected at %s:%s and returned with code %d, and version %s\n", ElasticURL, ElasticPort, code, info.Version.Number)

	ESCLIENT = client

	return true

}

func createESIndexIfNotExist(index string, mapping string) {

	ctx := context.Background()

	exists, err := ESCLIENT.IndexExists(index).Do(ctx)
	if err != nil {
		panic(err)
	}

	if !exists {
		createIndex, err := ESCLIENT.CreateIndex(index).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

}
