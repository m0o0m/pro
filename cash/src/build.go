// +build ignore
package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	goarch      string
	goos        string
	gocc        string
	gocxx       string
	cgo         string
	version     string
	race        bool
	environment string
)

/*
go run build.go govendor
go run build.go build
*/
//或者 bin
func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	flag.StringVar(&goarch, "goarch", runtime.GOARCH, "GOARCH")
	flag.StringVar(&goos, "goos", runtime.GOOS, "GOOS")
	flag.StringVar(&gocc, "cc", "", "CC")
	flag.StringVar(&gocxx, "cxx", "", "CXX")
	flag.StringVar(&cgo, "cgo-enabled", "", "CGO_ENABLED")
	flag.BoolVar(&race, "race", race, "Use race detector")
	flag.StringVar(&environment, "environment", "development", "Use environment testing production development")
	flag.StringVar(&version, "v", "1.0000", "Use version x.xxxx")

	flag.Parse()
	//设置gopath
	ensureGoPath()

	if flag.NArg() == 0 {
		log.Println("Usage: go run build.go build")
		return
	}
	cmds := flag.Args()
	//workingDir, _ = os.Getwd()
	//对应的命令
	args := []string{}
	for k, cmd := range cmds {
		if k == 0 {
			continue
		}
		args = append(args, cmd)
	}
	fmt.Println(cmds[0], args)
	switch cmds[0] {
	case "govendor":
		fmt.Println(cmds[0], args)
		runPrint("govendor", args...)
		//runPrint("govendor", "fetch","+out")
	case "build-run":
		clean()
		if len(args) == 0 {
			log.Println("Usage: go run build.go build  [bin name]")
			return
		}
		build(args[0], "./cmd/"+args[0], []string{})
		if goos == "windows" {
			goos += ".exe"
		}
		runPrint("../bin/" + args[0] + "_" + goos)
	case "run-front-nocopy":
		// TODO 避免开发阶段编译时间过长
		if len(args) == 0 {
			log.Println("Usage: go run build.go build  [bin name]")
			return
		}
		build(args[0], "./cmd/"+args[0], []string{})
		if goos == "windows" {
			goos += ".exe"
		}
		runPrint("../bin/" + args[0] + "_" + goos)
	case "build":
		build_clean()
		if len(args) == 0 {
			log.Println("Usage: go run build.go build  [bin name]")
			return
		}
		build(args[0], "./cmd/"+args[0], []string{})
	case "builds":
		clean()
		if len(args) == 0 {
			log.Println("Usage: go run build.go build  [bin names]")
			return
		}
		for _, binary := range args {
			//goos = "windows"
			//build(binary, "./cmd/"+binary, []string{})
			goos = "linux"
			build(binary, "./cmd/"+binary, []string{})
			//goos = "darwin"
			//build(binary, "./cmd/"+binary, []string{})
		}
	case "test":
		if len(args) == 0 {
			test("./..")
		} else {
			test(args[0])
		}
	case "template-build":
		clean()
		if len(args) == 0 {
			log.Println("Usage: go run build.go build  [bin names]")
			return
		}
		for _, binary := range args {
			goos = "windows"
			build(binary, "./cmd/"+binary, []string{})
			goos = "linux"
			build(binary, "./cmd/"+binary, []string{})
			goos = "darwin"
			build(binary, "./cmd/"+binary, []string{})
		}
		// cleanTemplate()
		// if len(args) == 0 {
		// 	log.Println("Usage: go run build.go build  [bin name]")
		// }
		// goos = "windows"
		// build(args[0], "./cmd/"+args[0], []string{}, "template/bin/")
		// goos = "linux"
		// build(args[0], "./cmd/"+args[0], []string{}, "template/bin/")
		// goos = "darwin"
		// build(args[0], "./cmd/"+args[0], []string{}, "template/bin/")
		// if goos == "windows" {
		// 	goos += ".exe"
		// }
		//runPrint("template/bin/" + args[0] + "_" + goos)
	case "sha-dist":
		shaFilesInDist()

	case "clean":
		clean()
	default:
		log.Fatalf("Unknown command %q", cmds[0])
	}

}

//配置 gopath
func ensureGoPath() {
	//govendor add +external

	g_gopath := os.Getenv("GOPATH")

	if g_gopath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		gopath := filepath.Clean(filepath.Join(cwd, "/../")) //+ ":" + filepath.Clean(cwd)
		log.Println("GOPATH is", gopath)
		os.Setenv("GOPATH", gopath)
	} else {
		//设置gopath
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		gopath := filepath.Clean(filepath.Join(cwd, "/../")) + string(os.PathListSeparator) + g_gopath
		log.Println("GOPATH is", gopath)
		os.Setenv("GOPATH", gopath)
	}
}

func ChangeWorkingDir(dir string) {
	os.Chdir(dir)
}

//依赖管理
func setup() {
	vendor := ""
	//配置依赖
	str_arr := strings.Split(vendor, "\n")
	for _, v := range str_arr {
		if len(v) > 0 {
			runPrint("go", "get", "-v", v)
		}
	}
	//runPrint("go", "get", "-v", "github.com/kardianos/govendor")
	//runPrint("go", "install", "-v", "./pkg/cmd/grafana-server")
}

func test(pkg string) {
	setBuildEnv()
	runPrint("go", "test", "-short", "-timeout", "60s", pkg)
}

//

//编译
func build(binaryName, pkg string, tags []string, paths ...string) {
	var binary string
	switch len(paths) {
	case 1:
		binary = paths[0] + binaryName
	case 0:
		binary = "../bin/" + binaryName
	default:
		log.Fatal("paths count err,there can only be one")
	}
	switch goos {
	case "windows":
		binary += "_windows.exe"
	case "linux":
		binary += "_linux"
	case "darwin":
		binary += "_darwin"
	}

	rmr(binary, binary+".md5")
	args := []string{"build", "-ldflags", ldflags()}
	if len(tags) > 0 {
		args = append(args, "-tags", strings.Join(tags, ","))
	}
	if race {
		args = append(args, "-race")
	}
	args = append(args, "-o", binary)
	args = append(args, pkg)
	setBuildEnv()

	runPrint("go", "version")
	runPrint("go", args...)

	// Create an md5 checksum of the binary, to be included in the archive for
	// automatic upgrades.
	err := md5File(binary)
	if err != nil {
		log.Fatal(err)
	}
}

func ldflags() string {
	var b bytes.Buffer
	b.WriteString("-w")
	b.WriteString(fmt.Sprintf(" -X global.VERSION=%s", version))
	b.WriteString(fmt.Sprintf(" -X main.ENVIRONMENT=%s", environment))
	//b.WriteString(fmt.Sprintf(" -X main.commit=%s", getGitSha()))
	//b.WriteString(fmt.Sprintf(" -X main.buildstamp=%d", buildStamp()))
	return b.String()
}

func rmr(paths ...string) {
	for _, path := range paths {
		log.Println("rm -r", path)
		os.RemoveAll(path)
	}
}

//build_clean

func build_clean() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	rmr(filepath.Join(cwd, "../bin"))
}

//配置文件
func clean() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	rmr(filepath.Join(cwd, "../bin/log"))
	//rmr(filepath.Join(cwd, "../bin/template"))
	rmr(filepath.Join(cwd, "../bin/_etc"))
	createDir("../bin/log/server") //创建log目录
	createDir("../bin/log/front")  //创建log目录
	createDir("../bin/log/admin")  //创建log目录
	createDir("../bin/log/wap")    //创建log目录
	//createDir("../bin/template")   //创建template目录
	createDir("../bin/_etc") //创建etc目录
	createDir("../bin/log")  //创建log目录
	copyFiles(getFileNames("etc", "../bin/_etc"))
	//copyFiles(getFileNames("template", "../bin/template"))
}

//配置文件-template目录
func cleanTemplate() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	rmr(filepath.Join(cwd, "template/bin/_etc"))
	createDir("template/bin")            //创建bin目录
	createDir("template/bin/log")        //创建log目录
	createDir("template/bin/_etc/fonts") //创建fonts目录
	syncCopyFiles(getFileNames("etc/fonts", "template/bin/_etc/fonts"))
	copyFile("etc/front.yaml", "template/bin/_etc/front.yaml")
}

func setBuildEnv() {
	os.Setenv("GOOS", goos)
	if strings.HasPrefix(goarch, "armv") {
		os.Setenv("GOARCH", "arm")
		os.Setenv("GOARM", goarch[4:])
	} else {
		os.Setenv("GOARCH", goarch)
	}
	if goarch == "386" {
		os.Setenv("GO386", "387")
	}
	if cgo != "" {
		os.Setenv("CGO_ENABLED", cgo)
	}
	if gocc != "" {
		os.Setenv("CC", gocc)
	}
	if gocxx != "" {
		os.Setenv("CXX", gocxx)
	}
}

func getGitSha() string {
	v, err := runError("git", "rev-parse", "--short", "HEAD")
	if err != nil {
		return "unknown-dev"
	}
	return string(v)
}

func buildStamp() int64 {
	bs, err := runError("git", "show", "-s", "--format=%ct")
	if err != nil {
		return time.Now().Unix()
	}
	s, _ := strconv.ParseInt(string(bs), 10, 64)
	return s
}

func buildArch() string {
	os := goos
	if os == "darwin" {
		os = "macosx"
	}
	return fmt.Sprintf("%s-%s", os, goarch)
}

func run(cmd string, args ...string) []byte {
	bs, err := runError(cmd, args...)
	if err != nil {
		log.Println(cmd, strings.Join(args, " "))
		log.Println(string(bs))
		log.Fatal(err)
	}
	return bytes.TrimSpace(bs)
}

func runError(cmd string, args ...string) ([]byte, error) {
	ecmd := exec.Command(cmd, args...)
	bs, err := ecmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return bytes.TrimSpace(bs), nil
}

func createDir(dirName string) {
	err := os.MkdirAll(dirName, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
func getFileNames(src string, dst string) (srcFileNames []string, dstFileNames []string) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			if strings.Count(dst, "/") >= 20 {
				log.Println("文件夹层级过深")
				return
			}
			createDir(dst + "/" + file.Name())
			childSrcFileNames, childDstFileNames := getFileNames(src+"/"+file.Name(), dst+"/"+file.Name())
			srcFileNames = append(srcFileNames, childSrcFileNames...)
			dstFileNames = append(dstFileNames, childDstFileNames...)
		} else {
			srcFileNames = append(srcFileNames, src+"/"+file.Name())
			dstFileNames = append(dstFileNames, dst+"/"+file.Name())
		}
	}
	return
}

func syncCopyFiles(srcNames, dstNames []string) {
	time1 := time.Now().UnixNano()
	var wg sync.WaitGroup
	//for j := 0; j < 1000; j++ {
	for i := 0; i < len(srcNames); i++ {
		wg.Add(1)
		go copyFile(srcNames[i], dstNames[i], &wg)
	}
	//}
	wg.Wait()
	log.Println("耗时2", time.Now().UnixNano()-time1)
}
func copyFiles(srcNames, dstNames []string) {
	time1 := time.Now().UnixNano()
	//for j := 0; j < 1000; j++ {
	for i := 0; i < len(srcNames); i++ {
		copyFile(srcNames[i], dstNames[i])
	}
	//}
	log.Println("耗时", time.Now().UnixNano()-time1)
}

func copyFile(srcName, dstName string, wg ...*sync.WaitGroup) {
	if len(wg) > 0 {
		defer wg[0].Done()
	}
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err, 153)
		return
	}
	defer src.Close()
	_, err = os.Open(dstName)
	if err == nil {
		//文件存在不修改
		return
	}
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err, 358)
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatal(err, 364)
	}
}

func runPrint(cmd string, args ...string) {
	log.Println(cmd, strings.Join(args, " "))
	ecmd := exec.Command(cmd, args...)
	ecmd.Stdout = os.Stdout
	ecmd.Stderr = os.Stderr
	err := ecmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func md5File(file string) error {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	h := md5.New()
	_, err = io.Copy(h, fd)
	if err != nil {
		return err
	}

	out, err := os.Create(file + ".md5")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "%x\n", h.Sum(nil))
	if err != nil {
		return err
	}

	return out.Close()
}

func shaFilesInDist() {
	filepath.Walk("./dist", func(path string, f os.FileInfo, err error) error {
		if path == "./dist" {
			return nil
		}

		if strings.Contains(path, ".sha256") == false {
			err := shaFile(path)
			if err != nil {
				log.Printf("Failed to create sha file. error: %v\n", err)
			}
		}
		return nil
	})
}

func shaFile(file string) error {
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	h := sha256.New()
	_, err = io.Copy(h, fd)
	if err != nil {
		return err
	}

	out, err := os.Create(file + ".sha256")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(out, "%x\n", h.Sum(nil))
	if err != nil {
		return err
	}

	return out.Close()
}
