package godaemon

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
)

var goDaemon = flag.Bool("d", false, "run app as a daemon with -d=true or -d true.")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *goDaemon {
		log.Println("start go daemon")
		var params []string
		if len(flag.Args()) > 0 {
			params = flag.Args()[1:]
		}
		for {
			cmd, err := run(os.Args[0], params...)
			if err != nil {
				log.Printf("%s start error %v\n", os.Args[0], err)
				break
			}
			log.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
			ps, err := cmd.Process.Wait()
			if err != nil {
				log.Printf("%v Wait error %v\n", ps, err)
				break
			}
			if ps.Exited() {
				if strings.Contains(ps.String(), "status 2") {
					log.Printf("%s %t %#v %#v", ps.String(), ps.Success(), ps, err)
					continue
				}
				if strings.Contains(ps.String(), "status 0") {
					log.Printf("%s %t %#v %#v", ps.String(), ps.Success(), ps, err)
					break
				}
				//signal: killed
				if strings.Contains(ps.String(), "signal: killed") {
					continue
				}
				log.Printf("%s %t %#v %#v", ps.String(), ps.Success(), ps, err)
				break
			} else {
				log.Println("", ps, err)
			}
		}
		*goDaemon = false
		os.Exit(0)
	}
}

func run(name string, arg ...string) (*exec.Cmd, error) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	return cmd, err
}

//func serer() {
//	var map1 map[int]int
//	map1[1] = 121212
//}

//func main() {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/index", func(rw http.ResponseWriter, req *http.Request) {
//		serer()
//		rw.Write([]byte("hello, golang!\n"))
//	})
//	mux.HandleFunc("/index1", func(rw http.ResponseWriter, req *http.Request) {
//		rw.Write([]byte("hello1, golang!\n"))
//	})
//	time.AfterFunc(10*time.Second, func() {
//		serer()
//	})
//	log.Fatalln(http.ListenAndServe(":7071", mux))
//}
