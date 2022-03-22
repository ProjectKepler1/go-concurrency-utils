# du example

This program simulates the unix `du` program. It takes a directory
and calculates the total size of its files.

It uses the Golang standard library trace package to instrument the Go routines. See [main.go](main.go)

Run:

```shell
time go run main.go ./images <option> 2>/tmp/trace.out
```

valid options are:

```text
--- fast-IO ---
f1: serial for-each
f2: serial map-reduce
f3: parallel map (ordered) then reduce
f4: parallel map (unordered) then reduce
f5: parallel map-reduce

--- slow-IO ---
s1: serial for-each
s2: serial map-reduce
s3: parallel map (ordered) then reduce
s4: parallel map (unordered) then reduce
s5: parallel map-reduce
```

(e.g. `time go run main.go ./images s4 2>/tmp/trace.out`)

Then, launch the trace explorer:

```shell
go tool trace -http=:41053 /tmp/trace.out
```

Open `http://127.0.0.1:41053/trace` in your host browser.
