package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

// Resource represents an OSINT resource in the database
type Resource struct {
	ID           uint64       `json:"id"`
	Title        string       `json:"title"`
	ResourceType ResourceType `json:"resource_type" objectbox:"type:int64 converter:resourceTypeInt64"`
	Regions      []string     `json:"regions"`
	Comment      string       `json:"comment"`
	URL          string       `json:"url"`
	Description  string       `json:"description"`
	Tags         []string     `json:"tags"`
	License      []string     `json:"licenses"`
	Browser      []string     `json:"browsers"`
}

func GetResourceTemplate() *Resource {
	return &Resource{
		ID:           0,
		Title:        "",
		ResourceType: None,
		Regions:      []string{},
		Comment:      "",
		URL:          "",
		Description:  "",
		Tags:         []string{},
		License:      []string{},
		Browser:      []string{},
	}
}

// from DB value to runtime value (int64 -> ResourceType)
func resourceTypeInt64ToEntityProperty(dbValue int64) (ResourceType, error) {
	return ResourceType(dbValue), nil
}

// from runtime value to DB value (ResourceType -> int64)
func resourceTypeInt64ToDatabaseValue(r ResourceType) (int64, error) {
	return int64(r), nil
}
