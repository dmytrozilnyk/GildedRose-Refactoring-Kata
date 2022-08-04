package gildedrose

type NormalItem struct {
	*Item
}

func NewNormalItem(item *Item) UpdateI {
	return &NormalItem{item}
}

func (i *NormalItem) Update() {
	if i.Item.Quality > MinQuality {
		i.Item.Quality = i.Item.Quality - 1
	}

	i.Item.SellIn = i.Item.SellIn - 1

	if i.Item.SellIn < MinDays {
		if i.Item.Quality > 0 {
			i.Item.Quality = i.Item.Quality - 1
		}
	}
}
