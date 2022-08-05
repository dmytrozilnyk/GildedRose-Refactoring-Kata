package gildedrose

type normalItem struct {
	*Item
}

func NewNormalItem(item *Item) UpdateI {
	return &normalItem{item}
}

func (i *normalItem) Update() {
	if i.Quality > MinQuality {
		i.Quality -= 1
	}

	i.SellIn = i.SellIn - 1

	if i.SellIn < MinDays {
		if i.Quality > 0 {
			i.Quality -= 1
		}
	}
}
