+++
title = "The Uber Data Pattern"
description = "Sample data should have these attributes"
date = "2016-08-15"
categories = ["blog", "data",]
+++

# The Data Uber Pattern

Three Dimensions

- Location - geography, location
- Unit Description - size, name, price, type
-  Date - doesn't need a name


- Category
- Value
- Time

- Bananas, Apples, Grapes - gross size and cost
- Location - grocrey store, fruit stand, gas station
- Transactions - price, type, location, date, weight??

Two Dimensions
Two Facts

Dimensions should chain, and should be more or less static.  Though it could have small variations??
Facts should be relatable (hard?).  Facts are generated randomly, and should vary moderately.

Stores Sample
D1 : Fruit Attributes
D2 : Store Demographics
F1 : Bulk Fruit Purchased
F2 : Retail Fruit Sold

Data should be randomly generated every week, packaged and compressed into single zip file.

## D1 Fruit attributes
Name, Weight_lbs, Wholesale_Price, Crate_Size_Lbs
Apple, .33, 31.00, 33
Banana, .40, 17.00, 40
Orange, .375, 20.00 , 33

## D2 Store Demographics
Category, Volume, Footprint
Mom-n-pop, 10, 2500
Gas-n-go, 30, 4000
Urban, 120, 10000
Metro, 400, 12000
Suburban, 1000, 125000

First Header  | Second Header
  ------------- | -------------
  Content Cell  | Content Cell
  Content Cell  | Content Cell
