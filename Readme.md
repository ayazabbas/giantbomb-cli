## System Requirements
- Go installed - https://golang.org/doc/install#install
- GiantBomb API Key - https://www.giantbomb.com/api/

## Usage
- Clone this repo
- From the repo directory, run `go build cmd/giantbomb/giantbomb.go` to build the executable
- Set your API Key as an environment variable named `GB_API_KEY`
- Run the executable:
    - `./giantbomb` will show commands and available flags

## Example
- `./giantbomb search-games --name gta`
```
GAMES
+------------+----------------------------------+------------+
| GUID       | NAME                             | RELEASE YR |
+------------+----------------------------------+------------+
| 3030-13567 | Grand Theft Auto                 |          0 |
| 3030-36765 | Grand Theft Auto V               |       2013 |
| 3030-20457 | Grand Theft Auto IV              |       2008 |
| 3030-3724  | Grand Theft Auto III             |       2001 |
| 3030-44366 | Grand Theft Auto: iFruit         |       2014 |
| 3030-3527  | Grand Theft Auto Advance         |       2004 |
| 3030-21100 | Grand Theft Auto: Chinatown Wars |       2009 |
| 3030-10593 | Grand Theft Auto 2               |          0 |
| 3030-7120  | Grand Theft Auto: San Andreas    |       2004 |
| 3030-17036 | Grand Theft Auto: Vice City      |       2002 |
+------------+----------------------------------+------------+
```
- `./giantbomb describe-game --guid 3030-20457 --show-dlc`
```
GAMES
+------------+---------------------+------------+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| GUID       | NAME                | RELEASE YR | DECK                                                                                                                                                                                                                                              |
+------------+---------------------+------------+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| 3030-20457 | Grand Theft Auto IV |       2008 | Take on the role of Niko Bellic, a Serbian immigrant who comes to the US at his cousin Roman's request, to find a better life, search for "that special someone" and participate in lawless activities in an upgraded generation of Liberty City. |
+------------+---------------------+------------+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
DLCs
+-----------+------------------------+---------------+---------------------+---------------------------------------------------------------------------------------------------------------------------------------------------+
| GUID      | NAME                   | PLATFORM      | RELEASE DATE        | DECK                                                                                                                                              |
+-----------+------------------------+---------------+---------------------+---------------------------------------------------------------------------------------------------------------------------------------------------+
| 3020-6    | The Lost & Damned      | Xbox 360      | 2009-02-17 00:00:00 | The first downloadable add-on for Grand Theft Auto IV puts players in the role of a biker named Johnny Klebitz.                                   |
| 3020-7    | The Ballad of Gay Tony | Xbox 360      | 2009-10-29 00:00:00 | The Ballad of Gay Tony is the second downloadable episode for Grand Theft Auto IV where the player takes on the role of club manager, Luis Lopez. |
| 3020-1805 | The Lost & Damned      | PC            | 2010-04-13 00:00:00 | The first downloadable add-on for Grand Theft Auto IV puts players in the role of a biker named Johnny Klebitz.                                   |
| 3020-1688 | The Ballad of Gay Tony | PlayStation 3 | 2010-04-13 00:00:00 | The Ballad of Gay Tony is the second downloadable episode for Grand Theft Auto IV where the player takes on the role of club manager, Luis Lopez. |
| 3020-1689 | The Lost & Damned      | PlayStation 3 | 2010-04-13 00:00:00 | The first downloadable add-on for Grand Theft Auto IV puts players in the role of a biker named Johnny Klebitz.                                   |
+-----------+------------------------+---------------+---------------------+---------------------------------------------------------------------------------------------------------------------------------------------------+
```
