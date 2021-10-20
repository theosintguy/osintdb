package stores

type Stores struct {
    Table *Resources
}

func New() Stores {
    return Stores{
        Table: NewResources(),
    }
}