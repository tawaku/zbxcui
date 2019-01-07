# zbxcui
Zabbix event viewer in command line.

## Usage
Configure config.json to connect to zabbix server.
```
❯ cat config.json 
{
	"host": "127.0.0.1",
	"port": "80",
	"user": "Admin",
	"pass": "zabbix"
}
```

Then execute zbxcui.
```
❯ ./zbxcui
┌─Event─────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ Severity |Status    |Info      |Time                |Age              |Ack       |Host          |Name             │
│ average  |problem   |          |2019-01-07 17:33:51 |          24m28s |yes       |Zabbix server |                 │
│---                                                                                                                │
│  Time                |User                         |Message       |User action                                    │
│  2019-01-07 17:47:31 |Admin (Zabbix Administrator) |Acknowledged. |undefined                                      │
│---                                                                                                                │
│                                                                                                                   │
└───────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
┌───────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│                                                                                                                   │
└───────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
  k: Down, j: Up, /: Search, space: Toggle ack                                                                       
```

## Installation
Execute below in zbxcui/ directory.
```
❯ go build
```
