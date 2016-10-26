+++
title = "You need the nulls"
description = "Explore the edges of your data and find the world "
date = "2016-10-25"
categories = ["blog", "mix",]
+++
"Screws fall out all the time, the world is an imperfect place."—Bender, The Breakfast Club

In our perfectly squared off world of tables that map the world into rows and columns, related by keys and sped up by indexes, we have to realize that the world _is_ an imperfect place.  Our sales models sometimes don't account for the product we thought we sold.  Humans are terrible at remembering details, and machines don't always communicate the entire picture - screws fall out.

Most queries are inner joins - this table of sales maps perfectly to that catalog of products sold.  But stuff happens, and that's where left and right joins can be very interesting.  Try forcing the picture manually to see what comes into focus.  Maybe you're selling product you didn't enter into the catalog correctly?  Maybe you have product that's not being sold?  Maybe the cash register hiccuped at the wrong moment and only noted the fact that a sale happened and didn't capture the quantity or product.  That's okay - but it's one of the biggest reasons two different analysts came come up with slightly or even dramatically different sets of numbers that _appear_ to cover the same events.

So be willing to explore the edges - do a quality check and see what pops.  It might just fix some assumptions that you thought were rock solid.  