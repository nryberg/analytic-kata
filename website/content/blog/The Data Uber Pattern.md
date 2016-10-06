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

Three Dimensions to build the context - one of which should be outside one of the other two.
One Fact Table to join them all

Dimensions should chain, and should be more or less static.  Though it could have small variations??
Facts should be relatable (hard?).  Facts are generated randomly, and should vary moderately.


Stores Sample
Dimension_1 : Fruit Attributes
Dimension_D2 : Store Attributes
Dimension_D3 : Location Demographics
Fact : Retail Fruit Sold

Data should be randomly generated every week, packaged and compressed into single zip file.

## D1 Fruit Facts
<table class='table'>
 <tr>
  <th>Name</th>  <th> Weight_lbs</th>  <th> Wholesale_Price</th>  <th> Crate_Size_Lbs</th> </tr>
 <tr>
  <td>Apple</td>  <td> .33</td>  <td> 31.00</td>  <td> 33</td> </tr>
 <tr>
  <td>Banana</td>  <td> .40</td>  <td> 17.00</td>  <td> 40</td> </tr>
 <tr>
  <td>Orange</td>  <td> .375</td>  <td> 20.00 </td>  <td> 33</td> </tr>
 <tr>
 </tr>
</table>

## D2 Store Facts

<table class='table'>
 <tr>
  <th>Category</th>  <th> Volume</th>  <th> Footprint</th> </tr>
 <tr>
  <td>Mom-n-pop</td>  <td> 10</td>  <td> 2500</td> </tr>
 <tr>
  <td>Gas-n-go</td>  <td> 30</td>  <td> 4000</td> </tr>
 <tr>
  <td>Urban</td>  <td> 120</td>  <td> 10000</td> </tr>
 <tr>
  <td>Metro</td>  <td> 400</td>  <td> 12000</td> </tr>
 <tr>
  <td>Suburban</td>  <td> 1000</td>  <td> 125000</td> </tr>
 <tr>
 </tr>
</table>

Other Examples
===

Bike Rentals
D1 : Bicycle model
D2 : Manufacturer
D3 : Location
Fact : Rental, Time, Bike, Location

Basketball Games
D1 : team one
D2 : team two
D3 : school?
Fact: Games, Scores, Date, ?

Banks and Bucks
D1 : Customer
D2 : Branch
D3 : Product
Fact : Accounts, Balances, Last Use Date, Customer

Education
D1 : Student
D2 : Class
D3 : Grade
Fact : Class, Grade, Completed Date, Student ??
