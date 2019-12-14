# RosterKit

This is the beginning of a scraper framework for league rosters. Currently the focus is on supporting the Guardbook and Chattahooligan Hymnal apps for the Northern Guard Supporters and Chattahooligans, respectively. For now, this only grabs data on opponents for the Foes feature in the rosters and outputs it to a JSON file that's compatible with the 1.4.0+ version of the apps.

Long-term, the goal is to also create post-processors that can automatically reference scraped players on Wikipedia, Transfermarkt, etc. and output them to multiple formats, including CSV, JSON, or directly to a database. There is no timeline for this, however.