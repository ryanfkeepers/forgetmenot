package backend

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

type data struct {
	Groups         map[string]group    `json:"groups"`
	Entries        map[string]entry    `json:"entries"`
	GroupedEntries map[string][]string `json:"groupedEntries"`
}

type group struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type entry struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	URL   string `json:"url"`
	Deets string `json:"deets"`
}

func (a *App) Data() data {
	return data{
		Groups: map[string]group{
			"1": {"1", "G1", "blue"},
			"2": {"2", "G2", "pink"},
		},
		Entries: map[string]entry{
			"a": {"a", "ea", "", "ea deets"},
			"b": {"b", "eb", "", "eb deets"},
			"c": {"c", "ec", "", "ec deets"},
		},
		GroupedEntries: map[string][]string{
			"1": {"a", "b", "c"},
			"2": {"c"},
		},
	}
}
