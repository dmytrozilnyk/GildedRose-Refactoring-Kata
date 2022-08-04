package gildedrose

type SulfurasItem struct {
	*Item
}

func NewSulfurasItem(item *Item) UpdateI {
	return &SulfurasItem{item}
}

func (i *SulfurasItem) Update() {
	//Nothing to do, quality always 80
}
