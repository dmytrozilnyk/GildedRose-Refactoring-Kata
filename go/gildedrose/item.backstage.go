package gildedrose

type BackstageItem struct {
	*Item
}

func NewBackstageItem(item *Item) UpdateI {
	return &BackstageItem{item}
}

func (i *BackstageItem) Update() {
	if i.Item.Quality < MaxQuality {
		i.Item.Quality = i.Item.Quality + 1
		if i.Item.SellIn < 11 {
			if i.Item.Quality < MaxQuality {
				i.Item.Quality = i.Item.Quality + 1
			}
		}
		if i.Item.SellIn < 6 {
			if i.Item.Quality < MaxQuality {
				i.Item.Quality = i.Item.Quality + 1
			}
		}
	}

	i.Item.SellIn = i.Item.SellIn - 1

	if i.Item.SellIn < MinDays {
		i.Item.Quality = i.Item.Quality - i.Item.Quality
	}
}
