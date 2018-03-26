# Small parser to compare leaderboard times from specific tracks and users

## Why

Playing https://www.velocidrone.com/ since quite some time now, to keep improving my flying skills during the winter month
or to just learn some new tricks. To also keep practicing on the tracks where I am not as good, I created a small parser
to parse all the tracks and show me where I am good and where I am bad. So I can keep the focus on the tracks where I am
not that good.

## Usage

You need the go runtime installed, currently there is no release with an executable available. After that, you can set your username
in the config.yaml file and run. If you want to add an additional user to compare against, just use the commandline parameter.

Default the tool is scanning all tracks that are defined in the config.yaml file. If you want to make sure that all tracks are there,
you can run it with the validate commandline argument to see if all tracks are configured.

Default ordering of the results is by scenario. If you want all good tracks on top and the bad at bottom, use the commandline parameter
orderBy like described below.

Add the end of the parsing, you will get some statistics with your ranking. Keep in mind, statistics currently only work for a single user
check. MultiUser statistics don't work until now. If you scan form ore than one user, statistics might be corrupt.

If your rank is 999 on some track, this means you are not inside the Top100 and so it is not possible to find a ranking.

### Parsing leaderboard

    go run Velociparser.go -filter=<trackfilter> -user=<additionalUser> -orderBy=<orderValue> -cache=false

If you don't use commandline arguments, all tracks are parsed with the user
that is configured in config.yaml. If the first argument is set, only some tracks are parsed.
If the second argument is set, an additional user is parsed to compare your results to this user.

The parameter cache is set to true for default. It will write some files on your disk. It will refresh the files after 10 minutes.
If the parameter is set to false, everything is scanned again.

Order value can be rank or track, but rank is the default

Another example, scannig the tracks for all VRL Tracks, oder by track

    go run VelociParser.go -filter=VRL -orderBy=track

### Validate if all tracks are in the config.yml

    go run Velociparser.go -validate=true

## Configuration

Users: Array of users that should be compared

Rank: The rank you want to be compared against

Scene: Array of Scenes and Tracks to be compared


## Available Tracks to scan

- River2
- Sportshall
- Subway
- Hangar
- Industrial Wasteland
- Football Stadium
- Countryside
- Night Factory
- Karting Track
- Blank Canvas Day
- Blank Canvas Night
- Birmingham NEC
- Warehouse
- Underground Carpark
- Coastel
- City
- Red Bull Ring
