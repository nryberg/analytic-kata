Analysis Challenge #1 - Flipping the data for the rows

## Problem
You've got a data set that's nice and clean, _except_ for the tiny problem of source system summary.  The data is actually a downstream report from two pre-joined tables.  The first is simple - a list of sales including products and total dollar amounts.  The second is simple - a list of the products that could be sold.  The combination is a little evil - the aggregate report only goes to the basket level.  The products included are summarized and concatenated with semicolons (;).

### Sales

