package main

import "errors"

func main() {
}

type item struct {
	Name            string
	SellIn          int
	Quality         int
	degradationRate int
}

type ageingItem struct {
	item
}

type legendaryItem struct {
	item
}

type eventAwareItem struct {
	item
	rule10daysTillEventApplied bool
	rule5daysTillEventApplied  bool
	ruleEventIsPastApplied     bool
}

type Degradable interface {
	updateQuality()
}

func (i *item) updateQuality() {
	if i.SellIn == 0 {
		i.degradationRate *= 2
	}
	i.SellIn -= 1
	if i.Quality > i.degradationRate {
		i.Quality -= i.degradationRate
	}
}

func (i *ageingItem) updateQuality() {
	if i.SellIn == 0 {
		i.degradationRate *= 2
	}
	i.SellIn -= 1
	if i.Quality+i.degradationRate <= 50 {
		i.Quality += i.degradationRate
	}
}

func (s *legendaryItem) updateQuality() {
	//not expected to be sold or decrease in quality
}

// "Backstage passes"
func (i *eventAwareItem) updateQuality() {

	applied := false

	if i.SellIn == 0 && !i.ruleEventIsPastApplied {
		i.Quality = 0
		i.rule5daysTillEventApplied = true
		applied = true
	} else if i.SellIn <= 5 && !i.rule5daysTillEventApplied {
		if i.Quality+i.degradationRate <= 50 {
			i.Quality += 3
		}
		applied = true
		i.rule5daysTillEventApplied = true
	} else if i.SellIn <= 10 && !i.rule10daysTillEventApplied {
		if i.Quality+i.degradationRate <= 50 {
			i.Quality += 2
		}
		applied = true
		i.rule10daysTillEventApplied = true
	}

	if !applied && i.Quality+i.degradationRate <= 50 {
		i.Quality += i.degradationRate
	}

	i.SellIn -= 1
}

// The Quality of an item is never negative and is never more than 50
func NewItem(name string, sellIn, quality, rate int) (Degradable, error) {
	if len(name) < 1 {
		return nil, errors.New("bad Name")
	}
	if quality < 0 || quality > 50 {
		return nil, errors.New("quality outside limits")
	}
	return &item{name, sellIn, quality, rate}, nil
}

func NewAgeingItem(name string, sellIn, quality, rate int) (Degradable, error) {
	newItem, e := NewItem(name, sellIn, quality, rate)
	if e != nil {
		return nil, e
	}
	i := newItem.(*item)
	return &ageingItem{*i}, nil
}

func NewLegendaryItem(name string) (Degradable, error) {
	if len(name) < 1 {
		return nil, errors.New("bad Name")
	}
	return &legendaryItem{item{name, 0, 0, 0}}, nil
}

func NewEventAwareItem(name string, sellIn, quality int) (Degradable, error) {
	if len(name) < 1 {
		return nil, errors.New("bad Name")
	}
	item := item{name, sellIn, quality, 1}
	return &eventAwareItem{item, false, false, false}, nil
}

func NewShortTermItem(name string, sellIn, quality, rate int) (Degradable, error) {
	if len(name) < 1 {
		return nil, errors.New("bad Name")
	}
	return &item{name, sellIn, quality, rate * 2}, nil
}
