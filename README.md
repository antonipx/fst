File System Torture
===================

Benchmark creation, listing and removal of large file system directory trees.

Usage
=====
```
fst [-o op] [-r path] [-l levels] [-d dirs] [-f files] [-s sparse size] [-v]

  -o : operand c:create l:list d:delete a:all (default "a")
  -r : root dir (default "./fstdata")
  -l : levels (default 3)
  -d : subdirs to create on each level (default 10)
  -f : files to create on each level (default 10)
  -s : sparse file size to create (default 1GB)
  -p : display pgrogress every second
  -v : verbose (print all file/folder names)
```
