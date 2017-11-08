# Small parser to compare leaderboard times from specific tracks and users

## Goal

1. Find own username and time for track
2. Find friends username and time for track
    1. Compare to own time
3. Find next ahead of own username
    1. Show difference to get the next better place
4. Find first for track and show difference to own position

## ToDo
1. Beautify results - can always be better
2. ~~Create parser to check if there are new tracks that are not tracked~~
3. ~~Parallel parsing~~
4. ~~Order by track name~~

### Sequencial execution of parsing all leaderboards
Total time for parsing:  1m39.563299323s

### Parallel execution of parsing all leaderboards
Total time for parsing:  56.648224155s

## Usage

    go run Velociparser.go -filter=<trackfilter> -user=<additionalUser> -orderBy=<orderValue>

If you don't use commandline arguments, all tracks are parsed with the user
that is configured in config.yaml. If the first argument is set, only some tracks are parsed.
If the second argument is set, an additional user is parsed to compare your results to this user.

Order value can be rank or track, but track is the default

## Validate if all tracks are in the config.yml

    go run Velociparser.go -validate=true

## Configuration

Users: Array of users that should be compared
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
