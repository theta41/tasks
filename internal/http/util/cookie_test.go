package util

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/g6834/team41/tasks/internal/models"
)

func TestGetTokens(t *testing.T) {
	recorder := httptest.NewRecorder()

	http.SetCookie(recorder, &http.Cookie{Name: cookieAccessToken, Value: "123"})
	http.SetCookie(recorder, &http.Cookie{Name: cookieRefreshToken, Value: "456"})

	request := &http.Request{Header: http.Header{"Cookie": recorder.Header()["Set-Cookie"]}}

	expected := models.TokenPair{
		AccessToken:  "123",
		RefreshToken: "456",
	}

	tokens := GetTokensFromCookie(request)

	require.Equal(t, expected, tokens)
}

func TestGetTokensError(t *testing.T) {
	recorder := httptest.NewRecorder()

	request := &http.Request{Header: http.Header{"Cookie": recorder.Header()["Set-Cookie"]}}

	tokens := GetTokensFromCookie(request)

	require.Empty(t, tokens)
}

func TestGetLogin(t *testing.T) {
	recorder := httptest.NewRecorder()

	http.SetCookie(recorder, &http.Cookie{Name: cookieLogin, Value: "test@example.org"})

	request := &http.Request{Header: http.Header{"Cookie": recorder.Header()["Set-Cookie"]}}

	expected := "test@example.org"

	login := GetLoginFromCookie(request)

	require.Equal(t, expected, login)
}

func TestGetLoginError(t *testing.T) {
	recorder := httptest.NewRecorder()

	request := &http.Request{Header: http.Header{"Cookie": recorder.Header()["Set-Cookie"]}}

	login := GetLoginFromCookie(request)

	require.Empty(t, login)
}

func TestPutTokens(t *testing.T) {
	recorder := httptest.NewRecorder()

	PutTokensToCookie(recorder, models.TokenPair{
		AccessToken:  "123",
		RefreshToken: "456",
	})

	request := &http.Request{Header: http.Header{"Cookie": recorder.Header()["Set-Cookie"]}}

	cookieAccess, err := request.Cookie(cookieAccessToken)
	require.NoError(t, err)
	require.NotNil(t, cookieAccess)
	require.Equal(t, "123", cookieAccess.Value)

	cookieRefresh, err := request.Cookie(cookieRefreshToken)
	require.NoError(t, err)
	require.NotNil(t, cookieRefresh)
	require.Equal(t, "456", cookieRefresh.Value)
}
