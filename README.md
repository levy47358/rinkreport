# RinkReport
- Live NHL scores in your terminal


## Install
```bash
go get github.com/levy47358/rinkreport
```

## Usage
```bash
$ rinkreport --help
Usage of rinkreport:
  -date string
    	Show Scores from a specific date
  -goals
    	Show scoring summary
  -team string
    	Show only the score where this team is playing
```

- show all past, current & future games playing today:
```bash
$ rinkreport
+--------------------+-----------------+
| AMERANT BANK ARENA | 3RD PERIOD 3:59 |
+--------------------+-----------------+
| CHI                |               3 |
| FLA                |               4 |
+--------------------+-----------------+
+--------------------+---------+
| XCEL ENERGY CENTER | 6:00 PM |
+--------------------+---------+
| DAL                |       0 |
| MIN                |       0 |
+--------------------+---------+
```

- include a scoring summary for each game:

```bash
$ rinkreport --goals
+--------------------+-------+
| AMERANT BANK ARENA | FINAL |
+--------------------+-------+
| CHI                |     3 |
| FLA                |     4 |
+--------------------+-------+
Scoring Summary:
+------+------------------+--------+-------+
| TEAM |      SCORER      | PERIOD | TIME  |
+------+------------------+--------+-------+
| FLA  | O. Ekman-Larsson |      1 | 00:39 |
| CHI  | J. Dickinson     |      1 | 07:04 |
| FLA  | S. Reinhart      |      1 | 11:54 |
| CHI  | C. Bedard        |      1 | 19:04 |
| FLA  | S. Reinhart      |      2 | 07:00 |
| CHI  | C. Bedard        |      2 | 08:18 |
| FLA  | C. Verhaeghe     |      3 | 02:44 |
+------+------------------+--------+-------+
```

- view the score of a specific team (abbreviation):
```bash
$ rinkreport --team CHI 
+--------------------+-------+
| AMERANT BANK ARENA | FINAL |
+--------------------+-------+
| CHI                |     3 |
| FLA                |     4 |
+--------------------+-------+
```

- view games on a previous date (fmt: YYY-MM-DD):
```bash
$ rinkreport --date 2023-11-11
+----------------------+-------+
| LITTLE CAESARS ARENA | FINAL |
+----------------------+-------+
| CBJ                  |     4 |
| DET                  |     5 |
+----------------------+-------+
```
