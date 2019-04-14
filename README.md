## Installation
```sh
$ go get -u github.com/magiclea/mydevtool
```

## Timestamp Command
A timestamp conversion tool

### Usage
```
Usage:
  mydevtool timestamp

Aliases:
  timestamp, ts
```

### Examples
#### Epoch to human readable time
```sh
$ mydevtool timestamp 1546272000
Epoch to human readable time is:	2019-01-01 00:00:00 +0800 CST
```
#### Human readable time to epoch
```sh
$ mydevtool timestamp 2019 1 1 0 0 0
2019/01/01 00:00:00 Local -> 1546272000
```

## Pritunl Command
Using this command to communicate with pritunl-client.

### Usage
```
Usage:
  mydevtool pritunl [flags]

Flags:
      --disconnectAll     disconnect all connections
  -h, --help              help for pritunl
      --listOnly          list profiles only, no create new connection
      --password string   specify password
```

### Examples
#### connect to server by interactive command line
```sh
$ mydevtool pritunl
+----+-------------+--------------+
| ID |    Name     |    Status    |
+----+-------------+--------------+
|  1 |   server1   | Disconnected |
|  2 |   server2   | Disconnected |
|  3 |   server3   | Disconnected |
+----+-------------+--------------+
Enter Profile ID or Name: 1
Enter the username: user
Enter the password: ****
sent request to connect new one
```
### connect to server by command line
```sh
$ PIN=000000
$ OTP=123456
$ mydevtool pritunl --password=$PIN$OTP
+----+-------------+--------------+
| ID |    Name     |    Status    |
+----+-------------+--------------+
|  1 |   server1   | Disconnected |
|  2 |   server2   | Disconnected |
|  3 |   server3   | Disconnected |
+----+-------------+--------------+
Enter Profile ID or Name: 1
sent request to connect new one
```
if you would like to generate OTP automatically, try two-factor authentication agent, [the agent](`https://github.com/rsc/2fa`) is I'm used currently.
```sh
$ go get -u rsc.io/2fa
$ 2fa -add serv1
2fa key for gg: ****************
$ PIN=000000
$ mydevtool pritunl --password=$PIN$(2fa serv1)
+----+-------------+--------------+
| ID |    Name     |    Status    |
+----+-------------+--------------+
|  1 |   server1   | Disconnected |
|  2 |   server2   | Disconnected |
|  3 |   server3   | Disconnected |
+----+-------------+--------------+
Enter Profile ID or Name: 1
sent request to connect new one
```