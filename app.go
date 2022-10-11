package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/remeh/diago/pprof"
	"github.com/remeh/diago/proto"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type PprofLoadedData struct {
	ProfileId string    `json:"fileName"`
	Tree      TraceTree `json:"root"`
}

// App struct
type App struct {
	ctx        context.Context
	profileMap map[string]*pprof.Profile
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.profileMap = map[string]*pprof.Profile{}
	a.ctx = ctx
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}

// SelectFile loads the file if it is a valid pprof file and returns
// a frontend summary
func (a *App) SelectFile() (PprofLoadedData, error) {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return PprofLoadedData{}, err
	}

	prof, err := proto.ReadProtoFile(file)
	if err != nil {
		return PprofLoadedData{}, err
	}
	a.profileMap[file] = prof

	treeRepresentation := loadTreeProf(prof)

	return PprofLoadedData{ProfileId: file, Tree: treeRepresentation}, nil
}
