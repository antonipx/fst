package main

import (
    "log"
    "os"
    "strconv"
    "time"
    "flag"
    "math"
//    "syscall"
)

var root string
var levels int
var dirs int
var files int
var size int64
var verbose bool
var total float64
var items float64

func progress() {
    var prev float64
    for {
        log.Printf("Created %.f/%.f %.1f%% %.f/s ", items, total, 100*(items/total), items-prev)
        prev=items
        time.Sleep(time.Second)
    }
}

func make(lvl int, path string) {
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
        make(lvl + 1, path + "/d" + strconv.Itoa(d))
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

func main() {
    flag.StringVar(&root, "r", "/px", "root dir")
    flag.IntVar(&levels, "l", 3, "levels")
    flag.IntVar(&dirs, "d", 10, "subdirs to create on each level")
    flag.IntVar(&files, "f", 10, "files to create on each level")
    flag.Int64Var(&size, "s", 1024*1024*1024, "sparse file size to create")
    flag.BoolVar(&verbose, "v", false, "Verbose")
    flag.Parse()

    for i:=1; i<=levels; i++ {
        total+=math.Pow(float64(dirs), float64(i))
    }
    total=total*float64(files+1)
    log.Printf("Start in %s, Levels=%d, Dirs=%d, Files=%d, Total~%.f", root, levels, dirs, files, total)

    start := time.Now()
    err := os.MkdirAll(root, 0777)
    if err != nil && !os.IsExist(err) {
        log.Fatal(err)
    }
    go progress()
    make(0, root)
    log.Printf("Finished, %s", time.Now().Sub(start).String())
}
