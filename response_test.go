package adyen

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// TestResponseErrorResponseStatus - response is valid JSON, but status is > 299, error should be returned
func TestResponseErrorResponseStatus(t *testing.T) {
	t.Parallel()

	responseJSON := `
{
	"errorType" : "authorise",
	"errorCode" : "501",
	"message"   : "sample error",
	"status"    : 501
}
`
	body := strings.NewReader(responseJSON)

	resp := &http.Response{
		Status:        "OK 200",
		StatusCode:    200,
		ContentLength: int64(body.Len()),
		Body:          ioutil.NopCloser(body),
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	providerResponse := &Response{
		Response: resp,
		Body:     buf.Bytes(),
	}

	err = providerResponse.handleHTTPError()

	if err == nil {
		t.Fatal("Response should raise an error - got nil")
	}

	errorContent := "[authorise][501]: (501) sample error"

	if err.Error() != errorContent {
		t.Fatal("Expected error message ", errorContent, " got ", err.Error())
	}
}

// TestResponseNotValidJson - response is an empty script, error should be returned
func TestResponseNotValidJson(t *testing.T) {
	t.Parallel()

	body := strings.NewReader("")

	resp := &http.Response{
		Status:        "503",
		StatusCode:    503,
		ContentLength: int64(body.Len()),
		Body:          ioutil.NopCloser(body),
	}

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	providerResponse := &Response{
		Response: resp,
		Body:     buf.Bytes(),
	}

	err = providerResponse.handleHTTPError()

	if err == nil {
		t.Fatal("Response should raise an error - got nil")
	}
}
