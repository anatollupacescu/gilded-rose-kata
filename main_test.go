package main

import "testing"

func TestCanNotCreateItemWithBadArguments(t *testing.T) {
	i, e := NewItem("", 6, 9, 5)
	if i != nil || e == nil {
		t.Fail()
	}
	i, e = NewItem("test", 6, -1, 5)
	if i != nil || e == nil {
		t.Fail()
	}
	i, e = NewItem("test", 6, 51, 5)
	if i != nil || e == nil {
		t.Fail()
	}
}

func TestSellInDecreasesByOneDay(t *testing.T) {
	i, _ := NewItem("test", 3, 0, 1)
	i.updateQuality()
	item := i.(*item)
	if item.SellIn != 2 {
		t.Fail()
	}
}

func TestQualityDecreasesOneUnit(t *testing.T) {
	i, _ := NewItem("test", 6, 50, 1)
	i.updateQuality()
	item := i.(*item)
	if item.Quality != 49 {
		t.Fail()
	}
}

func TestQualityDecreasesTwoUnitsAfterSellDate(t *testing.T) {
	i, _ := NewItem("test", 0, 50, 1)
	i.updateQuality()
	item := i.(*item)
	if item.Quality != 48 {
		t.Fail()
	}
}

func TestQualityIsNeverNegative(t *testing.T) {
	i, _ := NewItem("test", 3, 0, 1)
	i.updateQuality()
	item := i.(*item)
	if item.Quality != 0 {
		t.Fail()
	}
}

func TestAgedBrieQualityIncreasesOneUnit(t *testing.T) {
	i, _ := NewAgeingItem("test", 3, 0, 1)
	i.updateQuality()
	item := i.(*ageingItem)
	if item.Quality != 1 {
		t.Fail()
	}
}

func TestAgedBrieQualityIsNeverMoreThanFifty(t *testing.T) {
	i, _ := NewAgeingItem("test", 3, 50, 1)
	i.updateQuality()
	item := i.(*ageingItem)
	if item.Quality > 50 {
		t.Fail()
	}
}

func TestSulfurasNeverHasToBeSold(t *testing.T) {
	i, _ := NewLegendaryItem("sulfuras")
	i.updateQuality()
	item := i.(*legendaryItem)
	if item.Quality != 0 || item.SellIn != 0 {
		t.Fail()
	}
}

func TestBackstagePassesQualityIncreasesOneUnitIfSellInMoreThanTen(t *testing.T) {
	i, _ := NewEventAwareItem("sulfuras", 11, 5)
	i.updateQuality()
	item := i.(*eventAwareItem)
	if item.Quality != 6 || item.SellIn != 10 {
		t.Fail()
	}
}

func TestBackstagePassesQualityIncreasesTwoUnitsIfSellInLessThanEleven(t *testing.T) {
	i, _ := NewEventAwareItem("sulfuras", 10, 5)
	i.updateQuality()
	item := i.(*eventAwareItem)
	if item.Quality != 7 || item.SellIn != 9 {
		t.Fail()
	}
}

func TestBackstagePassesQualityIncreasesThreeUnitsIfSellInLessThanSix(t *testing.T) {
	i, _ := NewEventAwareItem("sulfuras", 5, 5)
	i.updateQuality()
	item := i.(*eventAwareItem)
	if item.Quality != 8 || item.SellIn != 4 {
		t.Fail()
	}
}

func TestBackstagePassesQualityIsZeroAfterSellDate(t *testing.T) {
	i, _ := NewEventAwareItem("sulfuras", 0, 5)
	i.updateQuality()
	item := i.(*eventAwareItem)
	if item.Quality != 0 {
		t.Fail()
	}
}

func TestBackstagePassesQualityIsNeverMoreThanFiftyWhenIncreasingOne(t *testing.T) {
	i, _ := NewEventAwareItem("sulfuras", 7, 50)
	i.updateQuality()
	item := i.(*eventAwareItem)
	if item.Quality != 50 {
		t.Fail()
	}
}

func TestConjuredSellInDecreasesOneUnit(t *testing.T) {
	i, _ := NewShortTermItem("short", 7, 5, 1)
	i.updateQuality()
	item := i.(*item)
	if item.Quality != 3 {
		t.Fail()
	}
}

func TestConjuredQualityDecreasesFourUnitsAfterSellDate(t *testing.T) {
	i, _ := NewShortTermItem("short", 0, 5, 1)
	i.updateQuality()
	item := i.(*item)
	if item.Quality != 1 {
		t.Fail()
	}
}

func TestConjuredQualityIsNeverNegative(t *testing.T) {
	i, _ := NewShortTermItem("short", 0, 2, 1)
	i.updateQuality()
	item := i.(*item)
	if item.Quality < 0 {
		t.Fail()
	}
}
