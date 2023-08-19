package querybuilder

import (
	"fmt"
)

func GeneralQuerySearch(pag string,search string) string{
  var query string=`{
    "search_type": "match",
    "query": {
        "term": "%s",
        "field": "_all"
    },
    "from": %s,
    "max_results": 10,
    "_source": [
    ]
}`
  query=fmt.Sprintf(query,search,pag)

  return query
}
