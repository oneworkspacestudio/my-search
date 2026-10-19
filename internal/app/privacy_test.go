package app

import (
	"bytes"
	"encoding/json"
	"github.com/oneworkspacestudio/one-workspace/pkg/apphost"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

func TestRecentSearchIsOptInScopedAndErased(t *testing.T) {
	h, e := apphost.New(apphost.Config{ID: "search", Name: "My Search", Data: filepath.Join(t.TempDir(), "data"), Origin: "http://localhost:29011", Listen: "127.0.0.1:29011"}, nil)
	if e != nil {
		t.Fatal(e)
	}
	defer h.Close()
	if _, e = New(h); e != nil {
		t.Fatal(e)
	}
	if e = h.Bootstrap("alice@example.test", "fixture-password"); e != nil {
		t.Fatal(e)
	}
	call := func(method, path, body string, cookie *http.Cookie, csrf string) *httptest.ResponseRecorder {
		r := httptest.NewRequest(method, h.Config.Origin+path, bytes.NewBufferString(body))
		r.Header.Set("Origin", h.Config.Origin)
		r.Header.Set("X-CSRF-Token", csrf)
		if cookie != nil {
			r.AddCookie(cookie)
		}
		w := httptest.NewRecorder()
		h.Handler().ServeHTTP(w, r)
		return w
	}
	login := func(email string) (*http.Cookie, string) {
		w := call("POST", "/api/login", `{"email":"`+email+`","password":"fixture-password"}`, nil, "")
		if w.Code != 200 {
			t.Fatal(w.Code, w.Body)
		}
		c := w.Result().Cookies()[0]
		w = call("GET", "/api/me", "", c, "")
		var s apphost.Session
		json.Unmarshal(w.Body.Bytes(), &s)
		return c, s.CSRF
	}
	alice, csrf := login("alice@example.test")
	w := call("GET", "/api/preferences", "", alice, csrf)
	if w.Code != 200 || !strings.Contains(w.Body.String(), `"recent":false`) {
		t.Fatal("recent enabled by default")
	}
	w = call("POST", "/api/local-users", `{"email":"bob@example.test","password":"fixture-password"}`, alice, csrf)
	if w.Code != 201 {
		t.Fatal(w.Code, w.Body)
	}
	w = call("PUT", "/api/preferences", `{"recent":true,"terms":["private fixture"]}`, alice, csrf)
	if w.Code != 200 {
		t.Fatal(w.Code, w.Body)
	}
	bob, bcsrf := login("bob@example.test")
	w = call("GET", "/api/preferences", "", bob, bcsrf)
	if strings.Contains(w.Body.String(), "private fixture") {
		t.Fatal("recent leaked")
	}
	w = call("PUT", "/api/preferences", `{"recent":false,"terms":["private fixture"]}`, alice, csrf)
	if w.Code != 200 || strings.Contains(w.Body.String(), "private fixture") {
		t.Fatal("disabled history retained")
	}
	if call("GET", "/api/preferences", "", nil, "").Code != 401 {
		t.Fatal("anonymous preferences exposed")
	}
}
