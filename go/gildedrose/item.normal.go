package gildedrose

type NormalItem struct {
	*Item
}

func NewNormalItem(item *Item) UpdateI {
	return &NormalItem{item}
}

func (i *NormalItem) Update() {
	if i.Quality > MinQuality {
		i.Quality = i.Quality - 1
	}

	i.SellIn = i.SellIn - 1

	if i.SellIn < MinDays {
		if i.Quality > 0 {
			i.Quality = i.Quality - 1
		}
	}
}
