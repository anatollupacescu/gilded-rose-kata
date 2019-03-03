package rose_test

import (
	"testing"

	rose "github.com/anatollupacescu/gilded-rose-kata"
	"github.com/stretchr/testify/assert"
)

func TestItem(t *testing.T) {
	i := rose.NewItem("test", 1, 30, 1)
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
	i := rose.NewAgeingItem("Brie", 1, 30, 1)
	t.Run("increases quality with time", func(t *testing.T) {
		i.UpdateQuality()
		assert.Equal(t, 31, i.Quality)
		assert.Equal(t, 0, i.SellIn)
	})
}

func TestQualityIsAlwaysBetween0and50(t *testing.T) {
	t.Run("quality is never below zero", func(t *testing.T) {
		i := rose.NewItem("test", 3, 0, 1)
		i.UpdateQuality()
		assert.Equal(t, 0, i.Quality)
	})
	t.Run("quality is never more than fifty", func(t *testing.T) {
		i := rose.NewAgeingItem("test", 3, 50, 1)
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
		i := rose.NewEventAwareItem("sulfuras", 11, 5)
		i.UpdateQuality()
		assert.Equal(t, 6, i.Quality)
		assert.Equal(t, 10, i.SellIn)
	})
	t.Run("quality increases one unit ifSellInMoreThanTen", func(t *testing.T) {
		i := rose.NewEventAwareItem("sulfuras", 10, 5)
		i.UpdateQuality()
		assert.Equal(t, 7, i.Quality)
		assert.Equal(t, 9, i.SellIn)
	})
	t.Run("quality increases one unit if sell in more than ten", func(t *testing.T) {
		i := rose.NewEventAwareItem("sulfuras", 5, 5)
		i.UpdateQuality()
		assert.Equal(t, 8, i.Quality)
		assert.Equal(t, 4, i.SellIn)
	})
	t.Run("quality is zero after sell date", func(t *testing.T) {
		i := rose.NewEventAwareItem("sulfuras", 0, 5)
		i.UpdateQuality()
		assert.Equal(t, 0, i.Quality)
	})
	t.Run("quality is never more than fifty when increasing on", func(t *testing.T) {
		i := rose.NewEventAwareItem("sulfuras", 7, 50)
		i.UpdateQuality()
		assert.Equal(t, 50, i.Quality)
	})
}

func TestConjured(t *testing.T) {
	t.Run("sell in decreases one unit", func(t *testing.T) {
		i := rose.NewShortTermItem("short", 7, 5, 1)
		i.UpdateQuality()
		assert.Equal(t, 3, i.Quality)
	})
	t.Run("quality decreases four units after sell date", func(t *testing.T) {
		i := rose.NewShortTermItem("short", 0, 5, 1)
		i.UpdateQuality()
		assert.Equal(t, 1, i.Quality)
	})
	t.Run("quality is never negative", func(t *testing.T) {
		i := rose.NewShortTermItem("short", 1, 0, 1)
		i.UpdateQuality()
		assert.Equal(t, 0, i.Quality)
	})
}
