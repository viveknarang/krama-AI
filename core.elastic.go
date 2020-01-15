package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

//ESCLIENT ES client
var ESCLIENT *elastic.Client

func connectElastic() bool {

	client, err := elastic.NewClient(elastic.SetURL(ElasticURL + ":" + ElasticPort))

	if err != nil {
		panic(err)
	}

	ESCLIENT = client

	return pingES(false)

}

func pingES(silent bool) bool {

	if ESCLIENT == nil || !ESCLIENT.IsRunning() {
		return false
	}

	var isESUp bool

	ctx := context.Background()

	info, code, err := ESCLIENT.Ping(ElasticURL + ":" + ElasticPort).Do(ctx)

	if err != nil {

		isESUp = false
		panic(err)

	} else {

		isESUp = true

	}

	if !silent {
		fmt.Printf("ACTIVE PING FOR ES: Elasticsearch responding at %s:%s and returned with code %d, and version %s\n", ElasticURL, ElasticPort, code, info.Version.Number)
	}

	return isESUp

}

func createESIndexIfNotExists(index string, mapping string) {

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

func indexES(index string, mapping string, document interface{}, id string) bool {

	createESIndexIfNotExists(index, mapping)

	ctx := context.Background()

	_, err := ESCLIENT.Index().Index(index).Id(id).BodyJson(document).Do(ctx)

	if err != nil {
		panic(err)
	}

	return true

}

func queryES(index string, from int, to int, q string, fields []string) *elastic.SearchResult {

	ctx := context.Background()

	multiQuery := elastic.NewMultiMatchQuery(q, fields...)

	searchResult, err := ESCLIENT.Search().
		Index(index).
		Query(multiQuery).
		Aggregation("Brands", elastic.NewTermsAggregation().Field("Brands").Size(100)).
		Aggregation("Colors", elastic.NewTermsAggregation().Field("Colors").Size(100)).
		Aggregation("Sizes", elastic.NewTermsAggregation().Field("Sizes").Size(100)).
		From(from).Size(to).
		Pretty(true).
		Do(ctx)
	if err != nil {
		panic(err)
	}

	return searchResult

}

func deleteESDocumentByID(index string, id string) bool {

	ctx := context.Background()
	_, err := ESCLIENT.Delete().Index(index).Id(id).Do(ctx)

	if err != nil {
		panic(err)
	}

	return true
}
