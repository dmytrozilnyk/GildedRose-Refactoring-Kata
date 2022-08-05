package gildedrose

type sulfurasItem struct {
	*Item
}

func NewSulfurasItem(item *Item) UpdateI {
	return &sulfurasItem{item}
}

func (i *sulfurasItem) Update() {
	//Nothing to do, quality always 80
}
