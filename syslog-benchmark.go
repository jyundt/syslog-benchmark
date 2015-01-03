package main

import (
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func sendData(max *int, syslogwriter *syslog.Writer) (int, time.Time) {
	msg := 0
	//Start a channel to listen for os.Inerrupt
	sig := make(chan os.Signal, 1)
	//Start a channel to stop upon os.Inerrupt
	stop := make(chan bool, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		fmt.Println("")
		log.Println("Caught interrupt, stopping")
		stop <- true
	}()
	for {
		msg++
		syslogwriter.Write([]byte(strconv.Itoa(msg)))
		select {
		case <-stop:
			return msg, time.Now()
		default:
			if *max == -1 {
				continue
			} else if msg >= *max {
				return msg, time.Now()
			}
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
	log.Print("Starting sending messages")
	startTime := time.Now()
	lastmessage, stopTime := sendData(msgs, syslogwriter)

	//how many message did we send?
	log.Print("Total messages sent = ", lastmessage)
	log.Print("Total time = ", stopTime.Sub(startTime))
	log.Print("Throughput = ", float64(lastmessage)/stopTime.Sub(startTime).Seconds(), " message per second")
}
