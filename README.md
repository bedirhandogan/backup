### Scheduled Backup System
A program that makes periodic backups with Golang. You can specify the source and target folder and continue to take backups periodically within the time interval you specify.

### Setup
```bash
# install program
$ git clone https://github.com/bedirhandogan/backup.git

# to see usage
$ go run .
```

### Usage
```bash
# usage
$ go run . -source [source path] -destination [destination path] -time [backup duration]

# back up every 1 day e.g
$ go run . -source C:/Users/John/Project -destination C:/Users/John/Backup -time "d:1 h:0 m:0 s:0"
```