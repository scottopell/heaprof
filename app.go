package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type PprofLoadedData struct {
	ProfileId string `json:"fileName"`
}

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
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// SelectFile loads the file if it is a valid pprof file and returns
// a frontend summary
func (a *App) SelectFile() (PprofLoadedData, error) {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return PprofLoadedData{}, err
	}
	return PprofLoadedData{ProfileId: file}, nil
}
