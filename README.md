# Blog Aggregator

We're going to build an RSS feed aggregator in Go!  We'll call it "Gator", you know, because aggreGATOR 🐊.

The project assumes that you're already familiar with the Go programming language and SQL databases.

## RSS

The whole point of the gator program is to fetch the RSS feed of a website and store its content in a structured format in our database. 
That way we can display it nicely in our CLI.

RSS stands for **"Really Simple Syndication"** and is a way to get the latest content from a website in a structured format. 
It's fairly ubiquitous on the web: most content sites have an RSS feed.