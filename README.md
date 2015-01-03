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

##Examples
All defaults (100 messages, port 514, tcp, 'syslog-benchmark' tag)

    bash-4.2# ./syslog-benchmark -host 172.17.0.18                              
    2015/01/03 08:56:55 Starting sending messages
    2015/01/03 08:56:55 Total messages sent = 100
    2015/01/03 08:56:55 Total time = 2.76832ms
    2015/01/03 08:56:55 Throughput = 36122.991561669165 message per second
    bash-4.2# 

Sending 1,000,000 messages

    bash-4.2# ./syslog-benchmark -host 172.17.0.18 -msgs 1000000
    2015/01/03 08:58:32 Starting sending messages
    2015/01/03 08:59:02 Total messages sent = 1000000
    2015/01/03 08:59:02 Total time = 29.82881548s
    2015/01/03 08:59:02 Throughput = 33524.629922716595 message per second
    bash-4.2# 

Sending unlimited messages, stopping with ctrl+c

    bash-4.2# ./syslog-benchmark -host 172.17.0.18 -msgs -1 -proto udp -tag udp_test
    2015/01/03 09:00:31 Starting sending messages
    ^C
    2015/01/03 09:01:52 Caught interrupt, stopping
    2015/01/03 09:01:52 Total messages sent = 2439423
    2015/01/03 09:01:52 Total time = 1m20.709714431s
    2015/01/03 09:01:52 Throughput = 30224.651607279582 message per second
    bash-4.2# 

