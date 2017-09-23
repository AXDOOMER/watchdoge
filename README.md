# watchdoge

An utility that can reveal open ports. It can be used to find computers on a network and which ports will respond to TCP requests. 

### Usage

Scan an IP for open ports: `watchdoge [IP address] [timeout]`

Scan a subnet of an IP: `watchdoge [IP address]/24 [port] [timeout]`

#### Example
```
> go build watchdoge.go 
> ./watchdoge 192.168.1.1
TCP port 21 open
TCP port 22 open
TCP port 53 open
TCP port 80 open
TCP port 2000 open
TCP port 8291 open
TCP port 8728 open
TCP port 8729 open
done
```

##### Wow security. Much port scan. 
![watch_doge](https://user-images.githubusercontent.com/6194072/30776128-1222dd54-a06f-11e7-9b26-6f60664d1669.jpg)


