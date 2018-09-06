File System Torture
===================

Create large directory trees with large amount of large (sparse) files.

Usage
=====
```
fst [-r path] [-l levels] [-d dirs] [-f files] [-s sparse size] [-v]

  -r string
      	root dir (default "./fstdata")
  -l int
      	levels (default 3)
  -d int
    	subdirs to create on each level (default 10)
  -f int
    	files to create on each level (default 10)
  -s int
    	sparse file size to create (default 1073741824)
  -v	Verbose
```