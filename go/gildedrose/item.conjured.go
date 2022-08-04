package gildedrose

type ConjuredItem struct {
	*Item
}

func NewConjuredItem(item *Item) UpdateI {
	return &ConjuredItem{item}
}

func (i *ConjuredItem) Update() {
	if i.Item.Quality > MinQuality {
		i.Item.Quality = i.Item.Quality - 2
	}

	i.Item.SellIn = i.Item.SellIn - 1

	if i.Item.SellIn < MinDays {
		if i.Item.Quality > 0 {
			i.Item.Quality = i.Item.Quality - 2
		}
	}
}
