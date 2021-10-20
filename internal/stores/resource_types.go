package stores

import (
	"encoding/json"
	"github.com/wailsapp/wails"
	"osintdb/internal/model"
)

type ResourceTypes struct {
	runtime *wails.Runtime
	store   *wails.Store
}

func NewResourceTypes() *Resources {
	return &Resources{}
}

func (t *ResourceTypes) WailsInit(runtime *wails.Runtime) error {
	t.runtime = runtime
	t.store = runtime.Store.New("tableDataStore", "")
	return nil
}

func (t *ResourceTypes) Update(data []*model.ResourceType) {
	jsonData, _ := json.Marshal(data)
	jsonDataStr := string(jsonData)
	if jsonDataStr == "[]" { // TODO: find better solution here
		jsonDataStr = "[{}]"
	}
	err := t.store.Set(jsonDataStr)
	if err != nil {
		panic(err) // TODO: remove panic here
	}
}
