Analysis Challenge #1 - Flipping the data for the rows

## Problem
You've got a data set that's nice and clean, _except_ for the tiny problem of source system summary.  The data is actually a downstream report from two pre-joined tables.  The first is simple - a list of sales including products and total dollar amounts.  The second is simple - a list of the products that could be sold.  The combination is a little evil - the aggregate report only goes to the basket level.  The products included are summarized and concatenated with semicolons (;).  Just to make it fun, some sales only have one item, so no semicolons needed.

### Sales
Date | Products | Total
------- |--------------|-------
12/01/15 | Apple;Orange;Onion | $12.32
12/05/15 | Tomato;Grapes;Banana | $15.33
12/11/15 | Apple | $5.99
12/16/15 | Onion; Potato | $44.11
12/20/16 | Broccoli; Watermelon | $9.32

