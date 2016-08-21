+++
title = "Analytic Tool Categories"
description = "Analytic Tools come in a variety of shapes and flavors, but they all follow basic guidelines"
date = "2016-08-21"
+++


## Scripted Languages

These are tools that are extremely useful to the analyst, however they require a certain comfort level using a text editor, an ide of some description and debugging code.  Sometimes, they're the __only__ solution to thorny problems.

### Powershell


Formerly only of Windows fame, it's now [publicly available](https://github.com/powershell/powershell) for all the major platforms.  It surfaces a bunch of the useful dot net code in windows and can solve challenging issues with gathering, processing and sharing data.


### Go

[Go](https://golang.org/) (or Golang if you're searching) is a fantastic scripting language - powerful, fast, small and challenging to use.  Plus it uses a gopher for a mascot.  Can't go wrong there.

### R

The quintessential poster child for open source gone viral, R is everywhere you want to be.  Other than being memory based only, it's the right price, it's expected in academia and it's pretty darned powerful.  

### SAS

The poster child for commercial packaged analytic software, it's used in probably every fortune 500 company and probably comprises a significant portion of their licensing budget for software.  It's the lingua franca of a __lot__ of analytic teams, and well worth getting to know.  

### Python

This is the sibling academia child to R - it's used in a wide variety of circumstances and often is the "glue" that holds a process together.

## Visual Analytic Toolsets

It's hard to swing a dead Schroedinger's cat without hitting a new visual analytic tool.  That said, there's a couple of platforms that are becoming standards.

The [Gartner Magic Quadrant for BI and Analytics](https://www.gartner.com/doc/reprints?id=1-2XXET8P&ct=160204) is a good starting point for evaluating new tools.

### Tableau

Tableau is mature, extremely popular (10K attendees at the 2015 user's conference) and easy to use.  It's not free. It's worth the investment.

### [Qlik](http://www.qlik.com/)
Qlik is the primary competitor to Tableau - they're exploring a couple of different licensing models that may or may not be useful for you.


## Databases

Back in the day, the pioneers of analytics would do all of their work in SQL.  And they'd like it that way!  A lot of the work you do today, can still be done very effectively in databases with straight up SQL.  SQL __is__ the default analytic language for a lot of platforms and a good working knowledge is very helpful.  Each database vendor has its own SQL dialect, however they all generally follow the same basic pattern.  

### Postgres
The open source landscape is dominated by two platforms - Postgres and MySQL.  After Oracle bought MySQL, a large number of developers moved to the moral high ground of Postgres.  Whether or not that's technically or legally valid is a personal choice, but for analysts making the first choice, they will be in good company with Postgres.  It's available in a wide variety of settings.

### MySql
MySQL is a great database platform, and aside from structural differences between it and its sibling Postgres, you can't go wrong using it.

### Oracle
Commercial applications are often backed by one of just a few database platforms - and Oracle is quite often the first.  It may have been chosen back in the '80's, so the folks who implemented it may have retired by now.  That said, it's a venerable database still commonly used in very large corporate settings.

### DB2

IBM's DB2 still keeps on kicking - it's roots date wayyyyy back to the mainframe days, and along with Oracle is most often found in the heavy iron of an older company managing vast quantities of data.


### Microsoft SQL Server

A great medium level database system (though recent improvements have made it scale even better than ever) one of the chief advantages is inherent connectivity.  If you use Windows, and who doesn't in the corporate environment, you __already__ have the drivers installed to use it.  No IT support needed.  It. Just. Works.  

### No SQL

MongoDB, CouchDB, Redis and a host of other young pups are running around the boardwalk, kicking over trashcans, yelling, "WE DON'T NEED NO SQL!!!"  Depending on your analytic requirements, not having a structured table format may suit your needs quite nicely.

Hadoop and the corollary "big data" world is a fish (whale?) of an entirely different color.  If you have to process billions of rows of data daily, you will probably be running a cluster of Hadoop machines and analysis is a wholly different proposition.
