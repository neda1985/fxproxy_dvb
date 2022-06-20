package apis_test

import (
	"fxproxy/pkg/apis"
	"fxproxy/pkg/validator"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type TestHelper struct {
	Test *testing.T
	*require.Assertions
}

func TestCoreHandler(t *testing.T) {
	os.Setenv("PORT", ":8080")
	defer os.Unsetenv("PORT")
	os.Setenv("SCHEMA", "http")
	defer os.Unsetenv("SCHEMA")
	os.Setenv("DOWNSTREAM", "localhost:49153")
	defer os.Unsetenv("DOWNSTREAM")
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	validator.GlobalMatcher = validator.NewMatcher(validator.AllowedList)
	t.Parallel()
	tests := []struct {
		testName         string
		expectedResponse func(th *TestHelper, rec *httptest.ResponseRecorder)
		url              string
	}{
		{
			testName: "forbidden",
			expectedResponse: func(th *TestHelper, rec *httptest.ResponseRecorder) {
				th.Equal(http.StatusNotFound, rec.Code)
				th.Equal("not found", rec.Body.String())
			},
			url: "http://example.com/test",
		},
		{
			testName: "/company/",
			expectedResponse: func(th *TestHelper, rec *httptest.ResponseRecorder) {
				th.Equal(http.StatusOK, rec.Code)
				th.Equal("OK", rec.Body.String())

			},
			url: "http://example.com/company",
		},
		{
			testName: "/company/{id}",
			expectedResponse: func(th *TestHelper, rec *httptest.ResponseRecorder) {
				th.Equal(http.StatusOK, rec.Code)
				th.Equal("OK", rec.Body.String())

			},
			url: "http://example.com/company/sd45f768",
		},
		{
			testName: "/company/account",
			expectedResponse: func(th *TestHelper, rec *httptest.ResponseRecorder) {
				th.Equal(http.StatusOK, rec.Code)
				th.Equal("OK", rec.Body.String())

			},
			url: "http://example.com/company/account",
		},
		{
			testName: "/account",
			expectedResponse: func(th *TestHelper, rec *httptest.ResponseRecorder) {
				th.Equal(http.StatusOK, rec.Code)
				th.Equal("OK", rec.Body.String())

			},
			url: "http://example.com/account",
		},
		{
			testName: "/account/{id}",
			expectedResponse: func(th *TestHelper, rec *httptest.ResponseRecorder) {
				th.Equal(http.StatusOK, rec.Code)
				th.Equal("OK", rec.Body.String())

			},
			url: "http://example.com/account/acc734340",
		},
		{
			testName: "/{id}",
			expectedResponse: func(th *TestHelper, rec *httptest.ResponseRecorder) {
				th.Equal(http.StatusOK, rec.Code)
				th.Equal("OK", rec.Body.String())

			},
			url: "http://example.com/acc734340",
		},
		{
			testName: "/account/{id}/user",
			expectedResponse: func(th *TestHelper, rec *httptest.ResponseRecorder) {
				th.Equal(http.StatusOK, rec.Code)
				th.Equal("OK", rec.Body.String())

			},
			url: "http://example.com/account/acc234234/user",
		},
		{
			testName: "/tenant/account/blocked",
			expectedResponse: func(th *TestHelper, rec *httptest.ResponseRecorder) {
				th.Equal(http.StatusOK, rec.Code)
				th.Equal("OK", rec.Body.String())

			},
			url: "http://example.com/tenant/account/blocked",
		},
	}
	httpmock.RegisterResponder(http.MethodGet,
		`=~^http://localhost:49153/.*`,
		httpmock.NewStringResponder(200, `OK`))

	for i := range tests {
		test := tests[i]
		t.Run(test.testName, func(t *testing.T) {
			th := New(t)
			recorder := httptest.NewRecorder()
			newHttpTestRequest := httptest.NewRequest(http.MethodGet, test.url, nil)
			apis.CoreHandler(recorder, newHttpTestRequest)
			test.expectedResponse(th, recorder)

		})
	}

}
func New(t *testing.T) *TestHelper {
	th := TestHelper{
		Test:       t,
		Assertions: require.New(t),
	}
	return &th
}
