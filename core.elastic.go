package main

import (
	"context"
	"encoding/json"
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

func basicSearch(index string, from int, to int, query string, queryFields []string, responseFields []string) *elastic.SearchHits {

	ctx := context.Background()

	multiQuery := elastic.NewMultiMatchQuery(query, queryFields...)
	ss := elastic.NewSearchSource()

	ss.FetchSourceIncludeExclude(responseFields, nil)

	searchResult, err := ESCLIENT.Search().
		Index(index).
		SearchSource(ss).
		Query(multiQuery).
		From(from).Size(to).
		Pretty(true).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	return searchResult.Hits

}

func facetedSearch(index string, from int, to int, q string, queryFields []string, responseFields []string, termFacetFields []string, rangeFacetFields []map[string][]map[string]interface{}) *elastic.SearchResult {

	m0 := make(map[string]interface{})
	m1 := make(map[string]map[string]interface{})
	m2 := make(map[string]map[string]interface{})
	m3 := make(map[string]interface{})
	var m4 []interface{}
	m5 := make(map[string]interface{})

	// Range Aggregation for rangeFacetFields parameter
	for _, rangeF := range rangeFacetFields {

		for key, value := range rangeF {

			m3 = make(map[string]interface{})
			m4 = nil

			for _, ar := range value {

				m5 = make(map[string]interface{})
				m5["from"] = ar["from"]
				m5["to"] = ar["to"]
				m4 = append(m4, m5)

			}

			m3["field"] = key
			m3["ranges"] = m4

			m1[key] = make(map[string]interface{})
			m1[key]["range"] = m3

		}

	}

	// Term Aggregation for termFacetFields parameter
	for _, term := range termFacetFields {

		m3 = make(map[string]interface{})
		m3["field"] = term
		m1[term] = make(map[string]interface{})
		m1[term]["terms"] = m3

	}

	m2["multi_match"] = make(map[string]interface{})
	m2["multi_match"]["fields"] = queryFields
	m2["multi_match"] = make(map[string]interface{})
	m2["multi_match"]["query"] = q

	m0["query"] = m2
	m0["aggs"] = m1

	m0["from"] = from
	m0["size"] = to

	if len(responseFields) == 0 {
		responseFields = append(responseFields, "*")
	}

	m6 := make(map[string][]string)
	m6["includes"] = responseFields
	m0["_source"] = m6

	json, err := json.Marshal(m0)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	so := elastic.NewSearchService(ESCLIENT)
	so.Source(string(json))
	searchResult, err := so.Do(ctx)

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
