package main

import(
	"log"
	"os"
	"flag"
	"log/syslog"
	"fmt"
)

func main() {
	//Define our test flags
	host := flag.String("host","localhost","hostname of syslog server")
	proto := flag.String("proto","tcp","protocol of syslog server: tcp/udp")
	msgs := flag.Int("msgs",-1,"number of messages to send, will be unlimited if not set")
	tag := flag.String("tag ","","syslog message tag")

	//Quick sanity check on flags
	flag.Parse()
	if (*msgs <= 0 && *msgs != -1){
		fmt.Println("Error: must send at least one message!")
		os.Exit(1)
	}
	
	if !((*proto == "tcp")|| (*proto == "udp")){
		fmt.Println("Error: must specify tcp or udp!")
		os.Exit(1)
	}

	//Open a connection to our syslog server
	syslogwriter, e := syslog.Dial(*proto,*host,syslog.LOG_ERR,*tag)
	if e == nil {
		log.SetOutput(syslogwriter)
	}


	msg := 0
	for {
		msg++
		fmt.Println("msg:",msg)
		if *msgs == -1 { 
			continue 
		} else if msg >= *msgs {
			break
		}
	}

}