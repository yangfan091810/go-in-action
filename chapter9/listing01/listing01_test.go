package listing01

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T){
	url := "http://www.baidu.com/"
	statusCode := 200
	t.Log("Given the need to test downloading content")
	{
		t.Logf("\twhen checking %s for status code %d", url, statusCode )
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\t should be able  to make the Get call ", ballotX, err)
			}
			t.Log("\t\t make the get call", checkMark)
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\t should receive a %d status %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\t should receive a %d status %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
