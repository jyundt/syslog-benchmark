package main

import (
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"math"
	"os"
	"strconv"
	"os/signal"
)

func sendData(max *int, syslogwriter *syslog.Writer) {
	msg := 0
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig 
		os.Exit(123)
	}()
	for {
		msg++
		//fmt.Println("msg:",msg)
		syslogwriter.Write([]byte(strconv.Itoa(msg)))
		if math.Mod(float64(msg), 10000) == 0 {
			fmt.Println(msg)
		}
		if *max == -1 {
			continue
		} else if msg >= *max {
			break
		}
	}

}

func main() {
	//Define our test flags
	host := flag.String("host", "localhost", "hostname of syslog server")
	port := flag.Int("port", 514, "port of syslog server")
	proto := flag.String("proto", "tcp", "protocol of syslog server: tcp/udp")
	msgs := flag.Int("msgs", 100, "number of messages to send, -1 = unlimited")
	tag := flag.String("tag", "syslog-benchmark", "syslog message tag")

	//Quick sanity check on flags
	flag.Parse()
	if *msgs <= 0 && *msgs != -1 {
		fmt.Println("Error: must send at least one message!")
		os.Exit(1)
	}

	if !((*proto == "tcp") || (*proto == "udp")) {
		fmt.Println("Error: must specify tcp or udp!")
		os.Exit(1)
	}

	//Open a connection to our syslog server
	syslogwriter, e := syslog.Dial(*proto, *host+":"+strconv.Itoa(*port), syslog.LOG_ERR, *tag)
	if e != nil {
		log.Fatal(e)
	} else {
		defer syslogwriter.Close()
	}

	//Let's start sending data
	sendData(msgs, syslogwriter)
}
