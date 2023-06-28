package druid

import (
	"net/http"

	"github.com/jonah-rankin/go-druid/builder"
	"github.com/jonah-rankin/go-druid/builder/query"
)

const (
	NativeQueryEndpoint     = "druid/v2"
	SQLQueryEndpoint        = "druid/v2/sql"
	PolarisSQLQueryEndpoint = "v1/query/sql"
)

type QueryService struct {
	client *Client
}

func (q *QueryService) Execute(qry builder.Query, result interface{}, headers ...http.Header) (*Response, error) {
	var path string
	switch qry.Type() {
	case "sql":
		if q.client.polarisConnection {
			path = PolarisSQLQueryEndpoint
		} else {
			path = SQLQueryEndpoint
		}
	default:
		path = NativeQueryEndpoint
	}
	r, err := q.client.NewRequest("POST", path, qry)
	if err != nil {
		return nil, err
	}
	if len(headers) >= 1 {
		for k, v := range headers[0] {
			for _, vv := range v {
				r.Header.Set(k, vv)
			}
		}
	}
	resp, err := q.client.Do(r, result)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//func (q *QueryService) Cancel(query builder.Query) () {}

//func (q *QueryService) Candidates(query builder.Query, result interface{}) (*Response, error) {}

func (q *QueryService) Load(data []byte) (builder.Query, error) {
	return query.Load(data)
}
