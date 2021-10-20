package db

import (
	"encoding/json"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/wailsapp/wails"
	"os"
	"osintdb/internal/data"
	"osintdb/internal/model"
	"osintdb/internal/stores"
	"osintdb/internal/utils"
	"path/filepath"
)

func InitObjectBox() *objectbox.ObjectBox {
	// get osintDB dir
	configDir, err := utils.FindConfigDir()
	if err != nil {
		// use current working directory if all else fails
		configDir = os.Getenv("PWD")
	}
	configDir = filepath.Join(configDir, "objectbox")
	// general initialization
	objectBox, err := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Directory(configDir).Build()
	if err != nil {
		panic(err)
	}
	return objectBox
}

func loadDefaultDataIfEmpty(ob *objectbox.ObjectBox) {
	// check if db is empty and if so load initial data
	dbIsEmpty, err := model.BoxForResource(ob).IsEmpty()
	if err != nil {
		panic(err)
	}
	if dbIsEmpty {
		resourcesFile, err := utils.FindInitData()
		if err != nil {
			// don't load any data on error
			// TODO: log error for debugging
			return
		}
		err = importData(ob, resourcesFile)
		if err != nil {
			// don't load any data on error
			// TODO: log error for debugging
			return
		}
	}
}

func importData(ob *objectbox.ObjectBox, resourcesFile string) error {
	resources, err := data.Import(resourcesFile)
	if err != nil {
		return err
	}
	// TODO: "dedupe" (remove duplicates of new resources and pre-existing ones)
	resourceBox := model.BoxForResource(ob)
	_, err = resourceBox.PutMany(resources)
	return err
}

func exportData(ob *objectbox.ObjectBox, resourcesFile string) error {
	resourceBox := model.BoxForResource(ob)
	resources, err := resourceBox.GetAll()
	if err != nil {
		return err
	}
	err = data.Export(resources, resourcesFile)
	return err
}

type DB struct {
	ob      *objectbox.ObjectBox
	stores  stores.Stores
	runtime *wails.Runtime
}

func New() *DB {
	db := &DB{
		ob: InitObjectBox(),
	}
	loadDefaultDataIfEmpty(db.ob)
	return db
}

func (db *DB) WailsInit(runtime *wails.Runtime) error {
	db.runtime = runtime
	return nil
}

func (db *DB) Close() {
	db.ob.Close()
}

func (db *DB) RegisterStores(stores stores.Stores) {
	db.stores = stores
}

func (db *DB) updateResources() error {
	resourceBox := model.BoxForResource(db.ob)
	resources, err := resourceBox.GetAll()
	if err != nil {
		return err
	}
	db.stores.Table.Update(resources)
	return nil
}

func (db *DB) UpdateAll() error {
	return db.updateResources()
}

func (db *DB) GetResourceTypes() []string {
	rtv := model.ResourceTypeValues()
	rt := make([]string, len(rtv))
	for i, v := range rtv {
		rt[i] = v.String()
	}
	return rt
}

func (db *DB) GetTags() ([]string, error) {
	resources, err := db.getResources()
	if err != nil {
		return nil, err
	}
	// make slice of all tags combined
	tags := make([]string, 0)
	for _, resource := range resources {
		tags = append(tags, resource.Tags...)
	}
	// make slice unique
	tags = utils.UniqueStrings(tags)
	return tags, nil
}

func (db *DB) GetRegions() ([]string, error) {
	resources, err := db.getResources()
	if err != nil {
		return nil, err
	}
	// make slice of all regions combined
	regions := make([]string, 0)
	for _, resource := range resources {
		regions = append(regions, resource.Regions...)
	}
	// make slice unique
	regions = utils.UniqueStrings(regions)
	return regions, nil
}

func (db *DB) GetLicenses() ([]string, error) {
	resources, err := db.getResources()
	if err != nil {
		return nil, err
	}
	// make slice of all licenses combined
	licenses := make([]string, 0)
	for _, resource := range resources {
		licenses = append(licenses, resource.License...)
	}
	// make slice unique
	licenses = utils.UniqueStrings(licenses)
	return licenses, nil
}

func (db *DB) GetBrowsers() ([]string, error) {
	resources, err := db.getResources()
	if err != nil {
		return nil, err
	}
	// make slice of all browsers combined
	browsers := make([]string, 0)
	for _, resource := range resources {
		browsers = append(browsers, resource.Browser...)
	}
	// make slice unique
	browsers = utils.UniqueStrings(browsers)
	return browsers, nil
}

func (db *DB) getResources() ([]*model.Resource, error) {
	resourceBox := model.BoxForResource(db.ob)
	resources, err := resourceBox.GetAll()
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func (db *DB) AddResource(jsonData string) error {
	return db.putResource(jsonData)
}

func (db *DB) ModifyResource(jsonData string) error {
	return db.putResource(jsonData)
}

func (db *DB) GetResourceTemplate() (string, error) {
	jsonData, err := json.Marshal(model.GetResourceTemplate())
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (db *DB) putResource(jsonData string) error {
	resource := model.Resource{}
	err := json.Unmarshal([]byte(jsonData), &resource)
	if err != nil {
		return err
	}

	resourceBox := model.BoxForResource(db.ob)
	_, err = resourceBox.Put(&resource)
	if err != nil {
		return err
	}
	return db.updateResources()
}

func (db *DB) DeleteResource(id uint64) error {
	resourceBox := model.BoxForResource(db.ob)
	err := resourceBox.RemoveId(id)
	if err != nil {
		return err
	}
	return db.updateResources()
}


func (db *DB) ImportData() error {
	file := db.runtime.Dialog.SelectFile("Select the data to import", "*.json")
	if err := importData(db.ob, file); err != nil {
		return err
	}
	return db.UpdateAll()
	//TODO: allow importing from other formats than JSON by using a different extension
}

func (db *DB) ExportData() error {
	file := db.runtime.Dialog.SelectSaveFile("Select the export location", "*.json")
	return exportData(db.ob, file)
	//TODO: allow exporting filtered data instead of whole set
	//TODO: allow exporting to other formats than JSON by using a different extension
}

func (db *DB) ClearData() error {
	resourceBox := model.BoxForResource(db.ob)
	err := resourceBox.RemoveAll()
	if err != nil {
		return err
	}
	return db.UpdateAll()
	//TODO: fix bug where the frontend does not notice that the data got wiped (only when built as desktop app)
}