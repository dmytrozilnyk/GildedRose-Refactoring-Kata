package gildedrose

type BackstageItem struct {
	*Item
}

func NewBackstageItem(item *Item) UpdateI {
	return &BackstageItem{item}
}

func (i *BackstageItem) Update() {
	if i.Quality < MaxQuality {
		i.Quality += 1
		if i.SellIn < 11 {
			if i.Quality < MaxQuality {
				i.Quality += 1
			}
		}
		if i.SellIn < 6 {
			if i.Quality < MaxQuality {
				i.Quality += 1
			}
		}
	}

	i.Item.SellIn -= 1

	if i.SellIn < MinDays {
		i.Quality = MinQuality
	}
}
