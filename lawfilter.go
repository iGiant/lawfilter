package lawfilter

import (
	"net/http"
)

func LawFilter(url string) (bool, error) {
	response, err := http.Head(url)
	if err != nil {
		return false, err
	}
	if lawfilter, ok := response.Header["X-Project"]; ok {
		if len(lawfilter) > 0 {
			return lawfilter[0] == "lawfilter", nil
		}
	}
	return false, nil
}
