package gildedrose

import "fmt"

const (
	AgedBrie  string = "Aged Brie"
	Sulfuras  string = "Sulfuras, Hand of Ragnaros"
	Backstage string = "Backstage passes to a TAFKAL80ETC concert"
	Conjured  string = "Conjured Mana Cake"
)

// Item describes an item information
type Item struct {
	Name            string
	SellIn, Quality int
}

// ItemI depending on the product a different function is executed
type ItemI interface {
	UpdateQuantityNormalProduct(*Item)
	UpdateQuantityAgedBrie(*Item)
	UpdateQuantitySulfuras(*Item)
	UpdateQuantityBackstage(*Item)
	UpdateQuantityConjured(*Item)
}

func (i *Item) String() string {
	return fmt.Sprintf("%s: %d days left, quality is %d", i.Name, i.SellIn, i.Quality)
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		switch item.Name {
		case AgedBrie:
			item.UpdateQuantityAgedBrie(item)
		case Sulfuras:
			item.UpdateQuantitySulfuras(item)
		case Backstage:
			item.UpdateQuantityBackstage(item)
		case Conjured:
			item.UpdateQuantityConjured(item)
		default:
			item.UpdateQuantityNormalProduct(item)
		}
		fmt.Println(item.String())
	}
}

func (i Item) UpdateQuantityNormalProduct(item *Item) {
	item.SellIn = item.SellIn - 1
	if item.Quality > 0 {
		item.Quality = item.Quality - 1
	}
	if item.SellIn < 0 {
		if item.Quality > 0 {
			item.Quality = item.Quality - 1
		}
	}
}

func (i Item) UpdateQuantityAgedBrie(item *Item) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
	}
	item.SellIn = item.SellIn - 1
	if item.SellIn < 0 {
		if item.Quality < 50 {
			item.Quality = item.Quality + 1
		}
	}
}

func (i Item) UpdateQuantitySulfuras(item *Item) {
	//Nothing to do, quality always 80
}

func (i Item) UpdateQuantityBackstage(item *Item) {
	if item.Quality < 50 {
		item.Quality = item.Quality + 1
		if item.SellIn < 11 {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
		if item.SellIn < 6 {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
			}
		}
	}
	item.SellIn = item.SellIn - 1
	if item.SellIn < 0 {
		item.Quality = item.Quality - item.Quality
	}
}

func (i Item) UpdateQuantityConjured(item *Item) {
	if item.Quality > 0 {
		item.Quality = item.Quality - 2
	}
	item.SellIn = item.SellIn - 1
	if item.SellIn < 0 {
		if item.Quality > 0 {
			item.Quality = item.Quality - 2
		}
	}
}
