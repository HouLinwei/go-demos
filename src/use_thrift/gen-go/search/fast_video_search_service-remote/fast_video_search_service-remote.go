// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"search"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "   suggest(SuggestRequest request)")
	fmt.Fprintln(os.Stderr, "  VideoResult search(VideoRequest request)")
	fmt.Fprintln(os.Stderr, "  string getName()")
	fmt.Fprintln(os.Stderr, "   getCounters()")
	fmt.Fprintln(os.Stderr, "   getCountersByCategory(string prefix)")
	fmt.Fprintln(os.Stderr, "   getCounterNames()")
	fmt.Fprintln(os.Stderr, "  i64 getCounter(string key)")
	fmt.Fprintln(os.Stderr, "  i64 aliveSince()")
	fmt.Fprintln(os.Stderr, "  void shutdown()")
	fmt.Fprintln(os.Stderr, "  void setLogLevel(string name, string level)")
	fmt.Fprintln(os.Stderr, "   getPerfCounters()")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := search.NewFastVideoSearchServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "suggest":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Suggest requires 1 args")
			flag.Usage()
		}
		arg34 := flag.Arg(1)
		mbTrans35 := thrift.NewTMemoryBufferLen(len(arg34))
		defer mbTrans35.Close()
		_, err36 := mbTrans35.WriteString(arg34)
		if err36 != nil {
			Usage()
			return
		}
		factory37 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt38 := factory37.GetProtocol(mbTrans35)
		argvalue0 := search.NewSuggestRequest()
		err39 := argvalue0.Read(jsProt38)
		if err39 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Suggest(value0))
		fmt.Print("\n")
		break
	case "search":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Search requires 1 args")
			flag.Usage()
		}
		arg40 := flag.Arg(1)
		mbTrans41 := thrift.NewTMemoryBufferLen(len(arg40))
		defer mbTrans41.Close()
		_, err42 := mbTrans41.WriteString(arg40)
		if err42 != nil {
			Usage()
			return
		}
		factory43 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt44 := factory43.GetProtocol(mbTrans41)
		argvalue0 := search.NewVideoRequest()
		err45 := argvalue0.Read(jsProt44)
		if err45 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Search(value0))
		fmt.Print("\n")
		break
	case "getName":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetName requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetName())
		fmt.Print("\n")
		break
	case "getCounters":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetCounters requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetCounters())
		fmt.Print("\n")
		break
	case "getCountersByCategory":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetCountersByCategory requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetCountersByCategory(value0))
		fmt.Print("\n")
		break
	case "getCounterNames":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetCounterNames requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetCounterNames())
		fmt.Print("\n")
		break
	case "getCounter":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetCounter requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetCounter(value0))
		fmt.Print("\n")
		break
	case "aliveSince":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "AliveSince requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.AliveSince())
		fmt.Print("\n")
		break
	case "shutdown":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Shutdown requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Shutdown())
		fmt.Print("\n")
		break
	case "setLogLevel":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SetLogLevel requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.SetLogLevel(value0, value1))
		fmt.Print("\n")
		break
	case "getPerfCounters":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetPerfCounters requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetPerfCounters())
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
