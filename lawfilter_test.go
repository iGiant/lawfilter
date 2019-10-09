package lawfilter

import "testing"

func TestLawFilter(t *testing.T) {
	var samples = []struct {
		url    string
		result bool
	}{
		{
			"http://rutor.org",
			true,
		},
		{
			"http://yandex.ru",
			false,
		},
		{
			"http://ffff",
			false,
		},
	}
	for _, sample := range samples {
		result, err := LawFilter(sample.url)
		if err != nil && result {
			t.Error(err)
		}
		if result != sample.result {
			t.Errorf("url=%s, excepted %v, got %v", sample.url, sample.result, result)
		}
	}

}
