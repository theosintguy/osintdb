package stores

import (
	"encoding/json"
	"github.com/wailsapp/wails"
	"osintdb/internal/model"
)

type Resources struct {
	runtime *wails.Runtime
	store   *wails.Store
}

func NewResources() *Resources {
	return &Resources{}
}

func (t *Resources) WailsInit(runtime *wails.Runtime) error {
	t.runtime = runtime
	t.store = runtime.Store.New("resourcesStore", "")
	return nil
}

//func (t *Resources) Get() string {
//	return t.store.Get().(string)
//}

func (t *Resources) Update(data []*model.Resource) {
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