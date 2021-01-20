// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"github.com/vithubati/http-cache/http"
	"github.com/vithubati/http-cache/memo"
	"github.com/vithubati/http-cache/test_data"
	"testing"

)

var httpGetBody = http.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	test_data.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	test_data.Concurrent(t, m)
}

