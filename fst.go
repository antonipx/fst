package main

import (
    "log"
    "os"
    "strconv"
    "time"
    "flag"
)

var root string
var levels int
var dirs int
var files int
var size int64
var verbose bool

func make(lvl int, path string) {
    if lvl > levels  {
        return
    }
    
    if verbose {
        log.Printf("In %s lvl=%d", path, lvl)
    }
    
    os.Mkdir(path, 0777)
    
    for d := 0; d < dirs; d++ {
        make(lvl + 1, path + "/d" + strconv.Itoa(d))
    }
    
    var fn string
    for f := 0; f < files; f++ {
        fn = path + "/f" + strconv.Itoa(f)
        fd, _ := os.Create(fn)
        fd.Close()
        os.Truncate(fn, size)
    }
}

func main() {
    flag.StringVar(&root, "r", "./fstdata", "root dir")
    flag.IntVar(&levels, "l", 3, "levels")
    flag.IntVar(&dirs, "d", 10, "subdirs to create on each level")
    flag.IntVar(&files, "f", 10, "files to create on each level")
    flag.Int64Var(&size, "s", 1024*1024*1024, "sparse file size to create")
    flag.BoolVar(&verbose, "v", false, "Verbose")
    flag.Parse()
    log.Printf("Start in %s, Levels=%d, Dirs=%d, Files=%d Size=%d", root, levels, dirs, files, size)
    start := time.Now()
    os.MkdirAll(root, 0777)
    make(0, root)
    log.Printf("Finished, %s", time.Now().Sub(start).String())
}
