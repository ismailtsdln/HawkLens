package db

import (
	"context"

	"github.com/olivere/elastic/v7"
)

type ElasticDB struct {
	Client *elastic.Client
}

func NewElasticDB(url string) (*ElasticDB, error) {
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	return &ElasticDB{Client: client}, nil
}

func (edb *ElasticDB) IndexResult(ctx context.Context, res *ScanResult) error {
	_, err := edb.Client.Index().
		Index("hawklens_results").
		BodyJson(res).
		Do(ctx)
	return err
}

func (edb *ElasticDB) Search(ctx context.Context, query string) (*elastic.SearchResult, error) {
	return edb.Client.Search().
		Index("hawklens_results").
		Query(elastic.NewMultiMatchQuery(query, "query", "platform", "data_type")).
		Do(ctx)
}
