package Routes

import (
	"mymodule/Controllers"
	"net/http"
)

func LoadRoutes() {

	http.HandleFunc("/", Controllers.Index)
	http.HandleFunc("/new", Controllers.New)
	http.HandleFunc("/insert", Controllers.Insert)
}
