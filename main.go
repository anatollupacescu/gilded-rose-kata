package rose

type item struct {
	Name    string
	SellIn  int
	Quality int
}

type degradableQualityItem struct {
	item
	degradationRate int
}

func (i *degradableQualityItem) UpdateQuality() {
	if i.SellIn == 0 {
		i.degradationRate *= 2
	}
	i.SellIn--
	if i.Quality > i.degradationRate {
		i.Quality -= i.degradationRate
	}
}

type ageingItem struct {
	degradableQualityItem
}

func (i *ageingItem) UpdateQuality() {
	i.SellIn--
	if i.Quality+i.degradationRate <= 50 {
		i.Quality += i.degradationRate
	}
}

//"Sulfuras"
type legendaryItem struct {
	item
}

func (s *legendaryItem) UpdateQuality() {
	//not expected to be sold or decrease in quality
}

// "Backstage passes"
type eventAwareItem struct {
	degradableQualityItem
	rule10daysTillEventApplied bool
	rule5daysTillEventApplied  bool
	ruleEventIsPastApplied     bool
}

func (i *eventAwareItem) UpdateQuality() {

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

func NewItem(name string, sellIn, quality, rate int) degradableQualityItem {
	return degradableQualityItem{item: item{name, sellIn, quality}, degradationRate: rate}
}

func NewAgeingItem(name string, sellIn, quality, rate int) ageingItem {
	return ageingItem{degradableQualityItem{item{name, sellIn, quality}, rate}}
}

func NewLegendaryItem(name string) legendaryItem {
	return legendaryItem{item{name, 0, 80}}
}

func NewEventAwareItem(name string, sellIn, quality int) eventAwareItem {
	item := degradableQualityItem{item: item{name, sellIn, quality}, degradationRate: 1}
	return eventAwareItem{item, false, false, false}
}

func NewShortTermItem(name string, sellIn, quality, rate int) degradableQualityItem {
	return degradableQualityItem{item: item{name, sellIn, quality}, degradationRate: rate * 2}
}
