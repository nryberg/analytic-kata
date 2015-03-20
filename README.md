Analytic Kata
=============

The Analtyic Kata are a series of exercises built on the martial arts concept of
a system of short routines that can be coupled together into a program that a
student can learn from and demonstrate mastery of the art.  In this case, we're
kicking bits and bytes, and chopping rows of data, but the idea of continually
improving is the same. 

While this is hosted in Github for convenience, it's a system that can be ported
to almost any learning and analytic environment.  Individual dialects (R,
Python, Tableau, Alteryx, SAS, etc...) will be constructed for each set of kata.

Mastery of Technique
---

The physical repetition of a kata slightly varies each time.  Hopefully over
time it comes closer and closer to perfection.  Senior masters will provide
correction, and the student will build new muscle memory around the techniques. 

Ultimately, the kata will generate a new set of challenges based on the same
techniques.  Data will be gathered, joined, condensed and shared.  However since
the practice of the kata will be (typically) solo, this system will provide a
mechanism to add some variety to the actual data.  This will force the student
to get the technique correct and know that they're on the mark without simply
repeating a pattern thoughtlessly. 

Learning by Teaching 
---

Every black belt, whether in Judo, Aikido or R, has ingrained the technique to a
point where they literally cannot express in words what they're doing.  It's
vitally critical to the continued development of the master and the field of
potential new masters that they train the lower classes on a regular basis.
Besides ensuring the survival of the technique for future generations, it keeps
surfacing the content in new and more complete ways in the master's brain.

These kata are designed to provide a foundation for masters to quickly pass on
new techniques to novices and take advantage of the virtuous cycle of mastery
through teaching.

Your Moves
---

Most martial arts disciplines have a variety of basic moves using parts of the body in
combinations.  Analytics has a similar vocabulary of moves that are built
together into combinations.

Every set of kata will have at least three levels of difficulty.  The simple
should be a nice starting point, and the expert levels ideally fiendish.

The major categories are:

* Collect - gather the data and prep it for analysis
* Connect - join datasets into a more comprehensive picture
* Change - apply formulas and clean up data
* Choose - select the columns you need and filter the rows that are important
* Compare - group similar things together, understand the deeper visual and statistic patterns

- - - 

### [Collect](./collect/README.md)

Getting data from point A to B is a critical and basic to every analytic
process. 

#### Text Files 

CSV, the universal language of datasets

#### Databases

Basic connections and SQL to get you off the ground.

- - - 

### Connect

One of these things is not like another.  Or is it?

#### Matched Joins

The simplest of concepts - a one to one match.  Things get a little more
complicated with many-to-one and one-to-many, but this is basic relationships.

#### Unmatched Joins

You had a great dinner but made a little too much and now it's sitting in the
fridge.  These are leftovers and can be just as delicious the next day for
lunch. 

Not every dataset matches every other dataset, and the leftovers can be just as
interesting or more so than the original.

Think left/right outer joins.  

#### Cartesian Joins

Rarely but significantly it makes sense to join every possible match of one
dataset to every possible match of another dataset.  While this can generate
massive amounts of data, it can also open up insights otherwise impossible to
uncover.

#### Gather

This is no different than "wipe on, wipe off".  The youngest novice to the
craggiest 6 degree black belt has to acquire data from somewhere. 

- - - 

### Change

Parsing, modification and testing - we have a variety of ways of fixing,
improving and changing data on a row by row basis and prepping it for the
analytics to come.  

#### Parsing

Splitting hairs or data is often similarly frustrating.  The hard part is just
finding the delimiter.

#### Formulas

Some tools are fantastic with formulas - think Excel, others not so much.  Every
toolset needs to be able to modify data.

#### Testing

The often neglected science of knowing whether or not you're on the right path
must be in every analyst's toolbelt.  

- - - 

### Choose 

#### Selection 

Choose your friends and columns wisely, for they will help determine if you will
successfully finish your analytic on time, under budget and with great results. 

While every dataset can be enhanced with joined data, almost every dataset has
too much to say. Pare and prune the data to just the columns you need and be
confident you can always add back what was removed.

#### Filtering

Inch by inch, row by row, the complement to selection is choosing what to focus
on.  One billion rows can encompass the world of your analytics, but at its
heart what really matters is just a few thousand rows of critical information.

#### Sampling

Judgement, random or stratified, different tools can give you widely different
results for creating small test populations and statistical projections.

- - - 

### Compare

Your data is a given, but underneath the covers are patterns of groups and
statistical insights.  

#### Groups

We as humans pigeonhole.  It's our superpower.  We group this batch of items
together with that attribute and it's how we can even make sense of the larger
patterns of life.  These groups are typically text or date, but can include
numeric labels.

#### Statistics

Numbers themselves tell a story.  Hook them up to an overall understanding of
your data, and descriptive stats can unearth some amazing information, and most
analytic toolsets deliver at least rudimentary measures.

#### Visual Patterns

Simple bar charts and line graphs can bring to life stories that capture the
imagination of your viewers and change the world.  Or at the very least, avoid
the dreaded "Sea of Numbers" problem.

- - - 

Choice of Analytic Tools
===

Within the world of every major martial art form is hundreds of variations.  Two
Aikido dojos miles apart in the same major metropolitan area will teach differently and
value different schools of thought.  Both are valid and will provide both
physical benefit and safety. (Both will also be slightly dismissive of the
other, but that's human nature)

How you exercise your kata is dependent on what tools you have at hand.  Perhaps
you're a corporate SAS or Tableau user, maybe you're a data scientist exercising
R and Python, and maybe you're an Excel spreadsheet black belt.  Each tool is
valid and will provide analytic benefit, though they vary dramatically in their
cost and speed to analysis results.

With a little adaptation, you should be able to perform almost all of the kata
within one tool environment.  Remember, limitations will set you free.

