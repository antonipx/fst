File System Torture
===================

Create large directory trees with large amount of large (sparse) files.

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
  -v : verbose
```
