package gildedrose

type AgedBrieItem struct {
	*Item
}

func NewAgedBrieItem(item *Item) UpdateI {
	return &AgedBrieItem{item}
}

func (i *AgedBrieItem) Update() {
	if i.Quality < MaxQuality {
		i.Quality += 1
	}
	i.SellIn -= 1
	if i.SellIn < MinDays {
		if i.Quality < 50 {
			i.Quality += 1
		}
	}
}
