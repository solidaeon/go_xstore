package storeitem

type Item struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

type Inventory struct {
	Items []Item
}

func New() *Inventory {
	return &Inventory{
		Items: []Item{},
	}
}

func (i *Inventory) Add(item Item) {
	i.Items = append(i.Items, item)
}

func (i *Inventory) GetItems() []Item {
	return i.Items
}