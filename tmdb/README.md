# tmdb

A cli to search movies and rename files based on **TheMovieDB** database.

## Usage

### Start

```bash
make build
build/tmdb
```

### Search movie

`bashtmbd search Avengers`

### Rename movie based in TheMovieDB database

`tmbd rename Avengers.Endgame.2019.BlueRay.mkv`

### Rename movie and move file to another directory

`tmdb rename /path/Avengers.Endgame.2019.BlueRay/Avengers.Endgame.2019.BlueRay.mkv --move /path/destination/`
