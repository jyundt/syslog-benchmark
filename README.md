#syslog-benchmark

Syslog benchmark utility written in Go

##Background
At $DAYJOB, I was tasked with building an über rsyslog cluster to collect all logs from everything.
In order to make sure a single node in this cluster could handle a massive syslog load, I needed to build a tool that could simulate sufficient syslog traffic.  

##Usage
Specify as many (or as few) flags as you'd like.  Hit ctrl+c (or whatever your OS interrupt is) to stop sending messages.

Upon receiving an interrupt, `syslog-benchmark` will stop sending messages and perform some basic post processing.

###Flags

    -host="localhost": hostname of syslog server
    -msgs=100: number of messages to send, -1 = unlimited
    -port=514: port of syslog server
    -proto="tcp": protocol of syslog server: tcp/udp
    -tag="syslog-benchmark": syslog message tag

