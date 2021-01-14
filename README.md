# postproc
Brute force post processing, supports arrays of arrays of rules, attempts to encode how Dave figures Jira logs out
## History
Developed this while working on a bot implementation.

### Logs
Using skynet it is simple to get every log item for a time period that references any Jira issue. In a 24 hour period that's well into 6 figures of log items.

I've been monitoring and supporting this bot for a while so I've learned to classify various Jira states from their logs. I was curious what states were around that I did not know how to classify? Using skynet it's simple to download the last 24 hours of all logs with Jira references to a .csv file. 
### Code Classifier
I wrote a golang importer that reads the .csv and creates a map from Jira IDs to arrays of message. Every record in the .csv is processed, making sure every Jira referenced is added to the map and every message in the logs with that Jira is included in the map's array value for that Jira key.
### Ugly Monster Conditional
Now taking the map from Jiras to message arrays using brute force "if message #1 is this and #2 is that, call this Jira 'In progress'" and looped over the map's Jira keys checking each array of messages for a match. Any that don't match get noted for further review.

Now the next most common was Not yet active, so I added another clause to filter those out as well. Rinse, rather, repeat, and after 10 or 20 "classes" all the Jira states have been classified.

The result was one of the more hideous conditionals I've ever seen. I saw a zip code conditional once that was worse I suppose, but this was pretty bad.
### Clauses and Rules and []Rules
Thinking about it, I turned the individual clauses of the conditionals into a clause structure and defined rules as having a name and an array of clauses that have to all be true. Sloppy and not great, but it worked.

Eventually I noticed order isn't guaranted, so as a bandaid I allow each rule to be a set of one or more clause arrays, so each clause array under a rule can be checked and if any one is satisfied then it satisfies the rule. 
### UI
I've been looking for an excuse to try coding some UI in golang so I siezed on this project. Now that I have rules setup I added some additional structure to track all of the rules as an accessible collective and also added a function to check the map for any Jiras that failed to match one of the rules.

Using a simple button and dropdown UI built with the fyne widgets 
```
fyne.io/fyne/app
fyne.io/fyne/widget
```
a button driven front end was added. The drop down can be used to select a report of the Jiras that match any specific rule.
There's also a button that shows any Jiras that don't have a matching rule.

### Posted on github
Added the usual permissive license and these notes and published on github.
# Compatibility
This code was built and tested on a Mac only. This library works fine for me on Linux and Mac so this code liklely qorks on Linux just fine. I suspect it can be used on Windows with the proper support as well, but I haven't tried to figure that out.
# Running It
Without the proper dat in a .csv file it doesn't do much.

After cloning the source code change into the new postproc/cmd/postproc directory and run
```go build . ...``` 
This shouldproduce the maimn executable in the root project directory - 2 directories above postproc/cmd/postproc oin the postproc directory. Simply invoke it using gthe path
```../../main```
or change directory to it and invoke
```./main```
## Fixes
The current clause bit is lame, having a location independent clause (the message present anywhere in the array counts as a match for the clause) is likely a more reliable/useful model.