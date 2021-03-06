# Gilded Rose Refactoring Kata

Hi and welcome to team Gilded Rose.
As you know, we are a small inn with a prime location in a prominent city ran by a friendly innkeeper named Allison.
We also buy and sell only the finest goods.
Unfortunately, our goods are constantly degrading in quality as they approach their sell by date.
We have a system in place that updates our inventory for us. 
It was developed by a no-nonsense type named Leeroy, who has moved on to new adventures.
Your task is to add the new feature to our system so that we can begin selling a new category of items.
First an introduction to our system:

- All items have a SellIn value which denotes the number of days we have to sell the item
- All items have a Quality value which denotes how valuable the item is
- At the end of each day our system lowers both values for every item

Also:

- Once the sell by date has passed, Quality degrades twice as fast
- The Quality of an item is never negative or more than 50
- "Aged Brie" actually increases in Quality the older it gets
- "Sulfuras", being a legendary item, never has to be sold or decreases in Quality
  - additionally, its Quality is 80 and it never alters.
- "Backstage passes", like aged brie, increases in Quality as it's SellIn value approaches; 
  - when there are 10 days or less Quality increases by 2
  - when there are 5 days or less Quality increases by 3
  - when the concert is over the Quality drops to 0
- "Conjured" items degrade in Quality twice as fast as normal items
