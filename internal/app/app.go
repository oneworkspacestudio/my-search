package app

import (
	"context"
	"encoding/json"
	"github.com/oneworkspacestudio/one-workspace/pkg/apphost"
	"net/http"
)

type App struct{ H *apphost.Host }
type Settings struct {
	Recent bool     `json:"recent"`
	Terms  []string `json:"terms"`
}

func New(h *apphost.Host) (*App, error) {
	if e := h.Migrate(1, `CREATE TABLE search_preferences(user_id TEXT PRIMARY KEY,value TEXT NOT NULL);`); e != nil {
		return nil, e
	}
	a := &App{h}
	h.Mux.HandleFunc("GET /api/preferences", h.Require(a.get))
	h.Mux.HandleFunc("PUT /api/preferences", h.Require(a.put))
	return a, nil
}
func (a *App) Run(ctx context.Context) error { <-ctx.Done(); return nil }
func (a *App) get(w http.ResponseWriter, r *http.Request) {
	var value string
	p := Settings{Terms: []string{}}
	if a.H.DB.QueryRow("SELECT value FROM search_preferences WHERE user_id=?", apphost.Current(r).User.ID).Scan(&value) == nil {
		json.Unmarshal([]byte(value), &p)
	}
	apphost.JSON(w, 200, p)
}
func (a *App) put(w http.ResponseWriter, r *http.Request) {
	var p Settings
	if !apphost.Decode(w, r, &p) {
		return
	}
	if len(p.Terms) > 20 {
		apphost.Problem(w, 400, "Keep up to 20 recent searches")
		return
	}
	for _, term := range p.Terms {
		if len(term) > 200 {
			apphost.Problem(w, 400, "Search is too long")
			return
		}
	}
	if !p.Recent {
		p.Terms = []string{}
	}
	v, _ := json.Marshal(p)
	_, e := a.H.DB.Exec("INSERT INTO search_preferences VALUES(?,?) ON CONFLICT(user_id)DO UPDATE SET value=excluded.value", apphost.Current(r).User.ID, string(v))
	if e != nil {
		a.H.DBError(w, e)
		return
	}
	apphost.JSON(w, 200, p)
}
