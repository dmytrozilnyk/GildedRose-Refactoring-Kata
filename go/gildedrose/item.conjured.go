package gildedrose

type ConjuredItem struct {
	*Item
}

func NewConjuredItem(item *Item) UpdateI {
	return &ConjuredItem{item}
}

func (i *ConjuredItem) Update() {
	if i.Quality > MinQuality {
		i.Quality = i.Quality - 2
	}

	i.SellIn = i.Item.SellIn - 1

	if i.SellIn < MinDays {
		if i.Quality > 0 {
			i.Quality = i.Quality - 2
		}
	}
}
