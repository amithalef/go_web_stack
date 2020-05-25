package http_test

import (
	"fmt"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/test_utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestSampleHttp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}

func TestMain(t *testing.M) {
	container, _ := test_utils.StartMongoDbForTest()
	os.Exit(t.Run())
	container.Stop()
}

func TestHealthCheckReturns200(t *testing.T) {

}
