package rose_test

import (
	"testing"

	"github.com/anatollupacescu/gilded-rose-kata"
	"github.com/stretchr/testify/assert"
)

func TestItem(t *testing.T) {
	item := rose.Item{Name: "onion", SellIn: 1, Quality: 30}
	i := rose.NewItem(item, 1)
	t.Run("updates quality and sell in value", func(t *testing.T) {
		i.UpdateQuality()
		assert.Equal(t, 29, i.Quality)
		assert.Equal(t, 0, i.SellIn)
	})
	t.Run("quality degrades twice as fast when sell date is past", func(t *testing.T) {
		qb := i.Quality
		i.UpdateQuality()
		assert.Equal(t, -1, i.SellIn)
		assert.Equal(t, qb-2, i.Quality)
	})
}

func TestAgingItem(t *testing.T) {
	item := rose.Item{Name: "Brie", SellIn: 1, Quality: 30}
	i := rose.NewAgeingItem(item, 1)
	t.Run("increases quality with time", func(t *testing.T) {
		i.UpdateQuality()
		assert.Equal(t, 31, i.Quality)
		assert.Equal(t, 0, i.SellIn)
	})
}

func TestQualityIsAlwaysBetween0and50(t *testing.T) {
	t.Run("quality is never below zero", func(t *testing.T) {
		i := rose.NewItem(rose.Item{Name: "test", SellIn: 3}, 1)
		i.UpdateQuality()
		assert.Equal(t, 0, i.Quality)
	})
	t.Run("quality is never more than fifty", func(t *testing.T) {
		i := rose.NewAgeingItem(rose.Item{Name: "test", SellIn: 3, Quality: 50}, 1)
		i.UpdateQuality()
		assert.Equal(t, 50, i.Quality)
	})
}

func TestSulfurasNeverHasToBeSold(t *testing.T) {
	i := rose.NewLegendaryItem("sulfuras")
	i.UpdateQuality()
	assert.Equal(t, 80, i.Quality)
	assert.Equal(t, 0, i.SellIn)
}

func TestBackstagePasses(t *testing.T) {
	t.Run("quality increases one unit if sell in more than ten", func(t *testing.T) {
		item := rose.Item{Name: "sufluras", SellIn: 11, Quality: 5}
		i := rose.NewEventAwareItem(item)
		i.UpdateQuality()
		assert.Equal(t, 6, i.Quality)
		assert.Equal(t, 10, i.SellIn)
	})
	t.Run("quality increases one unit ifSellInMoreThanTen", func(t *testing.T) {
		item := rose.Item{Name: "sufluras", SellIn: 10, Quality: 5}
		i := rose.NewEventAwareItem(item)
		i.UpdateQuality()
		assert.Equal(t, 7, i.Quality)
		assert.Equal(t, 9, i.SellIn)
	})
	t.Run("quality increases one unit if sell in more than ten", func(t *testing.T) {
		item := rose.Item{Name: "sufluras", SellIn: 5, Quality: 5}
		i := rose.NewEventAwareItem(item)
		i.UpdateQuality()
		assert.Equal(t, 8, i.Quality)
		assert.Equal(t, 4, i.SellIn)
	})
	t.Run("quality is zero after sell date", func(t *testing.T) {
		item := rose.Item{Name: "sufluras", SellIn: 0, Quality: 5}
		i := rose.NewEventAwareItem(item)
		i.UpdateQuality()
		assert.Equal(t, 0, i.Quality)
	})
	t.Run("quality is never more than fifty when increasing on", func(t *testing.T) {
		item := rose.Item{Name: "sufluras", SellIn: 7, Quality: 50}
		i := rose.NewEventAwareItem(item)
		i.UpdateQuality()
		assert.Equal(t, 50, i.Quality)
	})
}

func TestConjured(t *testing.T) {
	t.Run("sell in decreases one unit", func(t *testing.T) {
		item := rose.Item{Name: "regular item", SellIn: 7, Quality: 5}
		i := rose.NewConjuredItem(item, 1)
		i.UpdateQuality()
		assert.Equal(t, 3, i.Quality)
	})
	t.Run("quality decreases four units after sell date", func(t *testing.T) {
		item := rose.Item{Name: "regular item", SellIn: 0, Quality: 5}
		i := rose.NewConjuredItem(item, 1)
		i.UpdateQuality()
		assert.Equal(t, 1, i.Quality)
	})
	t.Run("quality is never negative", func(t *testing.T) {
		item := rose.Item{Name: "regular item", SellIn: 1, Quality: 0}
		i := rose.NewConjuredItem(item, 1)
		i.UpdateQuality()
		assert.Equal(t, 0, i.Quality)
	})
}

func TestUpdateInventory(t *testing.T) {
	onions := rose.Item{Name: "onions", SellIn: 3, Quality: 5}
	onionsItem := rose.NewItem(onions, 2)

	brie := rose.Item{Name: "brie", SellIn: 2, Quality: 6}
	brieItem := rose.NewAgeingItem(brie, 1)

	rose.UpdateInventory(&onionsItem, &brieItem)

	assert.Equal(t, 3, onionsItem.Quality)
	assert.Equal(t, 2, onionsItem.SellIn)

	assert.Equal(t, 7, brieItem.Quality)
	assert.Equal(t, 1, brieItem.SellIn)
}
