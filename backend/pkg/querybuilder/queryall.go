package querybuilder

import "fmt"

func AllQuerySearch(pag string) string{
  var query string=`{
    "search_type": "alldocuments",
    "from":%s,
    "max_results": 10,
    "_source": [
    ]
  }`
  query=fmt.Sprintf(query,pag)

  return query
}
