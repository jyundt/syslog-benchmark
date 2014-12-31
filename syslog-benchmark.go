package main

import(
	"log"
	"os"
	"flag"
	"log/syslog"
	"fmt"
	"strconv"
	"math"
)

func main() {
	//Define our test flags
	host := flag.String("host","localhost","hostname of syslog server")
	port := flag.Int("port",514,"port of syslog server")
	proto := flag.String("proto","tcp","protocol of syslog server: tcp/udp")
	msgs := flag.Int("msgs",-1,"number of messages to send, will be unlimited if not set")
	tag := flag.String("tag","syslog-benchmark","syslog message tag")

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
	syslogwriter, e := syslog.Dial(*proto,*host + ":" + strconv.Itoa(*port),syslog.LOG_ERR,*tag)
	if e == nil {
		log.SetOutput(syslogwriter)
	} else {
		fmt.Printf("Error: could not communicate with %s://%s:%d\n",*proto,*host,*port)
		os.Exit(1)
	}

	msg := 0
	for {
		msg++
		//fmt.Println("msg:",msg)
		log.Print(msg)
		if math.Mod(float64(msg),10000) == 0 {
			fmt.Println(msg)
		}
		if *msgs == -1 { 
			continue 
		} else if msg >= *msgs {
			break
		}
	}

}
