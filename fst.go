package main

import (
    "log"
    "os"
    "strconv"
    "time"
    "flag"
    "math"
    "io/ioutil"
)

var root string
var levels int
var dirs int
var files int
var size int64
var verbose bool
var progress bool
var total float64
var items float64

func pp(what string, ch chan int) {
    var prev float64
    for {
        select {
            case <-ch:
                return
            default:
                log.Printf("%s Progress %.f/%.f %.1f%% %.f/s ", what, items, total, 100*(items/total), items-prev)
                prev=items
                time.Sleep(time.Second)
        }
    }
}

func mk(lvl int, path string) {
    if lvl > levels  {
        return
    }
    
    if verbose {
        log.Printf("In %s lvl=%d", path, lvl)
    }
    
    err := os.Mkdir(path, 0777)
    if err != nil && !os.IsExist(err) {
        log.Fatal(err)
    }

    items++
    
    for d := 0; d < dirs; d++ {
        mk(lvl + 1, path + "/d" + strconv.Itoa(d))
    }
    
    var fn string
    for f := 0; f < files; f++ {
        fn = path + "/f" + strconv.Itoa(f)
        fd, err := os.Create(fn)
        if err != nil && !os.IsExist(err) {
            log.Fatal(err)
        }
        fd.Close()
        os.Truncate(fn, size)
        items++
    }
}

func ls(path string) {
    lst, err := ioutil.ReadDir(path)
    if err != nil {
        return
    }

    for _, e := range lst {
        if e.Name()=="." || e.Name()==".." {
            continue
        }
        items++
        if e.IsDir() {
            ls(path + "/" + e.Name())
        }
        if verbose {
            log.Printf("In %s Entry %s", path, e.Name())
        }
    }
}

func rm(path string) {
    lst, err := ioutil.ReadDir(path)
    if err != nil {
        return
    }

    for _, e := range lst {
        if e.Name()=="." || e.Name()==".." {
            continue
        }
        items++
        if e.IsDir() {
            rm(path + "/" + e.Name())
        } else {
            os.Remove(path + "/" + e.Name())
        }
        if verbose {
            log.Printf("In %s Entry %s", path, e.Name())
        }
    }
    os.Remove(path)
}


func main() {
    var op string
    var start time.Time
    var ch chan int

    flag.StringVar(&root, "r", "./fstdata", "root dir")
    flag.IntVar(&levels, "l", 3, "levels")
    flag.IntVar(&dirs, "d", 10, "subdirs to create on each level")
    flag.IntVar(&files, "f", 10, "files to create on each level")
    flag.Int64Var(&size, "s", 1024*1024*1024, "sparse file size to create")
    flag.BoolVar(&verbose, "v", false, "Verbose")
    flag.BoolVar(&progress, "p", false, "Display Progress")
    flag.StringVar(&op, "o", "a", "Operand c:create l:list d:delete a:all")
    flag.Parse()

    for i:=1; i<=levels; i++ {
        total+=math.Pow(float64(dirs), float64(i))
    }
    total=total*float64(files+1)
    log.Printf("Root=%s, Levels=%d, Dirs=%d, Files=%d, Total~%.f", root, levels, dirs, files, total)

    if op=="c" || op=="a" {
        log.Printf("Start Creating in %s", root)
        start := time.Now()
        items = 0
        err := os.MkdirAll(root, 0777)
        if err != nil && !os.IsExist(err) {
            log.Fatal(err)
        }
        if progress {
            ch = make(chan int)
            go pp("Creating", ch)
        }
        mk(0, root)
        if progress {
            close(ch)
        }
        log.Printf("Finished Creating, %s, Approx %.f/s", time.Now().Sub(start).String(), items/float64(time.Now().Sub(start).Seconds()))
    }

    if op=="l" || op=="a" {
        ioutil.WriteFile("/proc/sys/vm/drop_caches", []byte("3\n"), 0644)
        log.Printf("Start Listing in %s", root)
        start = time.Now()
        items = 0
        if progress {
            ch = make(chan int)
            go pp("Listing", ch)
        }
        ls(root)
        if progress {
            close(ch)
        }
        log.Printf("Finished Listing, %s, Approx %.f/s", time.Now().Sub(start).String(), items/float64(time.Now().Sub(start).Seconds()))
    }

    if op=="d" || op=="a" {
        log.Printf("Start Deleting in %s", root)
        start = time.Now()
        items = 0
        if progress {
            ch = make(chan int)
            go pp("Deleting", ch)
        }
        rm(root)
        if progress {
            close(ch)
        }
        log.Printf("Finished Deleting, %s, Approx %.f/s", time.Now().Sub(start).String(), items/float64(time.Now().Sub(start).Seconds()))
    }
}
