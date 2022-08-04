package gildedrose

import "fmt"

const (
	AgedBrie  string = "Aged Brie"
	Sulfuras  string = "Sulfuras, Hand of Ragnaros"
	Backstage string = "Backstage passes to a TAFKAL80ETC concert"
	Conjured  string = "Conjured Mana Cake"
)

const (
	MaxQuality = 50
	MinQuality = 0
	MinDays    = 0
)

// Item describes an item information
type Item struct {
	Name            string
	SellIn, Quality int
}

// UpdateI - All products implement this interface
type UpdateI interface {
	Update()
}

func (i *Item) String() string {
	return fmt.Sprintf("%s: %d days left, quality is %d", i.Name, i.SellIn, i.Quality)
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		switch item.Name {
		case AgedBrie:
			NewAgedBrieItem(item).Update()
		case Sulfuras:
			NewSulfurasItem(item).Update()
		case Backstage:
			NewBackstageItem(item).Update()
		case Conjured:
			NewConjuredItem(item).Update()
		default:
			NewNormalItem(item).Update()
		}
		fmt.Println(item.String())
	}
}
