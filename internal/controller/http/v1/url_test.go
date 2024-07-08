package v1_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	v1 "url-shortener/internal/controller/http/v1"
	"url-shortener/internal/service"
	"url-shortener/internal/service/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateURLHandler(t *testing.T) {
	cases := []struct {
		name      string
		alias     string
		url       string
		respError string
	}{
		{
			name:  "Success",
			alias: "test_alias",
			url:   "https://google.com",
		},
		{
			name:      "Invalid URL",
			alias:     "test_alias",
			url:       "google.com",
			respError: "Invalid URL",
		},
		{
			name:      "Empty URL",
			alias:     "test_alias",
			url:       "",
			respError: "Empty URL",
		},
		{
			name:      "Empty Alias",
			alias:     "",
			url:       "https://google.com",
			respError: "Empty Alias",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			urlMock := mocks.NewUrl(t)

			if tc.respError == "" {
				urlMock.On("CreateURL", tc.url, mock.AnythingOfType("string")).Return(int(1), nil).Once()
			}

			service := &service.Service{
				Url: urlMock,
			}

			handler := echo.New()
			v1.NewRouter(handler, service)

			input := fmt.Sprintf(`{"url": "%s", "alias": "%s"}`, tc.url, tc.alias)

			fmt.Printf("Test: %s\n", tc.name)
			fmt.Printf("%s\n", input)

			req, err := http.NewRequest(http.MethodPost, "/api/v1/urls/create", bytes.NewReader([]byte(input)))
			req.Header.Set("Content-Type", "application/json")
			require.NoError(t, err)

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			fmt.Println(rr.Code)
			require.Equal(t, http.StatusOK, rr.Code)

			body := rr.Body.String()

			var resp v1.CreateURLResponse

			require.NoError(t, json.Unmarshal([]byte(body), &resp))
			require.Equal(t, tc.respError, resp.Error)

		})
	}
}
