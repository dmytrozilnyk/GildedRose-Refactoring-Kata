package gildedrose

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	name            string
	sellIn          int
	quality         int
	daysApply       int
	expectedSellIn  int
	expectedQuality int
}{
	{"normalProduct", 5, 20, 3, 2, 17},
	{"normalProduct", 2, 20, 3, -1, 16},
	{"normalProduct", -1, 20, 3, -4, 14},
	{"normalProduct", 2, 0, 3, -1, 0},
	{"normalProduct", 5, 0, 3, 2, 0},
	{"Aged Brie", 5, 20, 3, 2, 23},
	{"Aged Brie", 2, 20, 3, -1, 24},
	{"Aged Brie", 0, 47, 3, -3, 50},
	{"Aged Brie", -1, 50, 3, -4, 50},
	{"Aged Brie", -1, 20, 3, -4, 26},
	{"Aged Brie", 5, 50, 3, 2, 50},
	{"Sulfuras, Hand of Ragnaros", 1, 80, 3, 1, 80},
	{"Sulfuras, Hand of Ragnaros", 0, 80, 3, 0, 80},
	{"Sulfuras, Hand of Ragnaros", -1, 80, 3, -1, 80},
	{"Backstage passes to a TAFKAL80ETC concert", 15, 20, 3, 12, 23},
	{"Backstage passes to a TAFKAL80ETC concert", 12, 20, 3, 9, 24},
	{"Backstage passes to a TAFKAL80ETC concert", 12, 47, 3, 9, 50},
	{"Backstage passes to a TAFKAL80ETC concert", 9, 20, 3, 6, 26},
	{"Backstage passes to a TAFKAL80ETC concert", 7, 20, 3, 4, 27},
	{"Backstage passes to a TAFKAL80ETC concert", 7, 47, 3, 4, 50},
	{"Backstage passes to a TAFKAL80ETC concert", 5, 20, 3, 2, 29},
	{"Backstage passes to a TAFKAL80ETC concert", 3, 20, 3, 0, 29},
	{"Backstage passes to a TAFKAL80ETC concert", 2, 20, 3, -1, 0},
	{"Backstage passes to a TAFKAL80ETC concert", 5, 47, 3, 2, 50},
	{"Backstage passes to a TAFKAL80ETC concert", 0, 47, 3, -3, 0},
	{"Conjured Mana Cake", 5, 20, 3, 2, 14},
	{"Conjured Mana Cake", 2, 20, 3, -1, 12},
	{"Conjured Mana Cake", -1, 20, 3, -4, 8},
	{"Conjured Mana Cake", 2, 0, 3, -1, 0},
	{"Conjured Mana Cake", 5, 0, 3, 2, 0},
}

func TestGildedRose(t *testing.T) {
	for _, test := range tests {
		item := &Item{test.name, test.sellIn, test.quality}
		var items []*Item
		items = append(items, item)
		for day := 0; day < test.daysApply; day++ {
			UpdateQuality(items)
		}
		assert.Equal(t, test.expectedQuality, item.Quality)
		assert.Equal(t, test.expectedSellIn, item.SellIn)
	}
}

func BenchmarkTestGildedRose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			item := &Item{test.name, test.sellIn, test.quality}
			var items []*Item
			items = append(items, item)
			for day := 0; day < test.daysApply; day++ {
				UpdateQuality(items)
			}
		}
	}
}
