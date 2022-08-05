package gildedrose

type agedBrieItem struct {
	*Item
}

func NewAgedBrieItem(item *Item) UpdateI {
	return &agedBrieItem{item}
}

func (i *agedBrieItem) Update() {
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
