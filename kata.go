package rose

type (
	Item struct {
		Name    string
		SellIn  int
		Quality int
	}
	//quality decreases with time
	variableQualityItem struct {
		Item
		rate int
	}
	//quality increases with time
	ageingItem struct {
		variableQualityItem
	}
	//"Sulfuras"
	legendaryItem struct {
		Item
	}
	// "Backstage passes"
	backstageItem struct {
		variableQualityItem
		rule10daysTillEventApplied bool
		rule5daysTillEventApplied  bool
		ruleEventIsPastApplied     bool
	}
)

func (i *variableQualityItem) UpdateQuality() {
	if i.SellIn == 0 {
		i.rate *= 2
	}
	i.SellIn--
	if i.Quality > i.rate {
		i.Quality -= i.rate
	}
}

func (i *ageingItem) UpdateQuality() {
	i.SellIn--
	if i.Quality+i.rate <= 50 {
		i.Quality += i.rate
	}
}

func (s *legendaryItem) UpdateQuality() {
	//not expected to be sold or decrease in quality
}

func (i *backstageItem) UpdateQuality() {

	applied := false

	if i.SellIn == 0 && !i.ruleEventIsPastApplied {
		i.Quality = 0
		i.rule5daysTillEventApplied = true
		applied = true
	} else if i.SellIn <= 5 && !i.rule5daysTillEventApplied {
		if i.Quality+i.rate <= 50 {
			i.Quality += 3
		}
		applied = true
		i.rule5daysTillEventApplied = true
	} else if i.SellIn <= 10 && !i.rule10daysTillEventApplied {
		if i.Quality+i.rate <= 50 {
			i.Quality += 2
		}
		applied = true
		i.rule10daysTillEventApplied = true
	}

	if !applied && i.Quality+i.rate <= 50 {
		i.Quality += i.rate
	}

	i.SellIn -= 1
}

type QualityVarying interface {
	UpdateQuality()
}

func UpdateInventory(head QualityVarying, tail ...QualityVarying) {
	head.UpdateQuality()
	for _, item := range tail {
		item.UpdateQuality()
	}
}

func NewItem(item Item, rate int) variableQualityItem {
	return variableQualityItem{Item: item, rate: rate}
}

func NewConjuredItem(item Item, rate int) variableQualityItem {
	return variableQualityItem{Item: item, rate: rate * 2}
}

func NewAgeingItem(item Item, rate int) ageingItem {
	return ageingItem{variableQualityItem{Item: item, rate: rate}}
}

func NewLegendaryItem(name string) legendaryItem {
	return legendaryItem{Item{name, 0, 80}}
}

func NewEventAwareItem(item Item) backstageItem {
	return backstageItem{variableQualityItem: variableQualityItem{Item: item, rate: 1}}
}
