# Mtgrestservice
This README is still in draft

This tool is a small set of rest api, to be used to fetch some information regarding the game Magic The Gathering.
The information that are currently available are:

1. All expansion names in English, Spanish and Italian (Spanish and Italian where available).
2. Card composition of every expansion (Number of lands, commons, uncommons, etc)
3. If the set is preconstructed: number of decks in the set, and type (duel, planechase, etc.)
3. If it is an edition available only online or it is a "From the vault" set

Please note, that having a list of all the cards is beyond the purpose of this project (even because it has already done in the mtgjson), so for now i prefer to focus on other information that are not accessible in an organized way (for developers :) )

My future plans are to add list of all magic abilities, with their translations, and what is their effect.

# Prerequisites

No particular prerequisites

# Compile
In order to compile this tool you must have go language installed. And you need the following go libraries:

* gorilla/mux

To download this package just type in your go workspace:
```bash
go get -u github.com/gorilla/mux
```

Then to compile just move into project directory and run:

```bash
go install
```

# Launch
If it is the first time that you run this tool, make sure to copy the asset folder into the same folder of the executable file (in the future the path will be provided using a configuration file)

Cd into the folder where the executable is installed and run (if you have compiled it from sources it should in your workspace bin folder):

```bash
mtgrestservice
```

Once the service is started just try to access:
* http://&lt;hostname&gt;:4040/listeditions for a list of all editions in json format
* http://&lt;hostname&gt;:4040/getedition/{key}/{format} where key is the edition code, and format is xml or json, will return the information available for the edition specified in key in the selected format.

The port used by default is 4040, if you change it in config.properties, change the port accordingly in the above examples. 

# Using Docker
If you are using docker you can build a container containing the service. 
First you need to build the container with: 

```bash
docker build --tag=mtgrestservicedocker .
```

Then to run it:
```bash
docker run -p 4040:4040 --name=mtgrestservice -d mtgrestservicedocker
```

And that's it, now you are running mtgrestservice in a docker container!
In this case you don't need to install any go dependency.
# Help

I will be very happy if you want to help me with this project, you can do it in three ways:

1. Reporting bugs (using the project issue tracker)
2. Keeping the xml information up to date, and add more languages.
3. Help me in the development
4. Using it :)

# Known issues/TODO

There are several known issue:
* ~~When the information is not available the node will be inserted even if is empty.~~
* ~~Vault and Online information are empty nodes, so if the node is present for an edition, this mean that this edition has that attribute. Now when the tool generate the JSON output they are both present with the value null, in the case the edition doesn't have these attributes~~
* Fix bugs
* The english of this README should be fixed :)
