package gildedrose

type AgedBrieItem struct {
	*Item
}

func NewAgedBrieItem(item *Item) UpdateI {
	return &AgedBrieItem{item}
}

func (i *AgedBrieItem) Update() {
	if i.Item.Quality < MaxQuality {
		i.Item.Quality = i.Item.Quality + 1
	}

	i.Item.SellIn = i.Item.SellIn - 1

	if i.Item.SellIn < MinDays {
		if i.Item.Quality < 50 {
			i.Item.Quality = i.Item.Quality + 1
		}
	}
}
