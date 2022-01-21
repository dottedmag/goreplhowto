# Go REPL HOWTO

How to run your Go code in REPL:

- Produce symbols file using `yaegi extract`
- Load this file into REPL

See `cmd/repl/main.go` for REPL toplevel.
See `repl.sh` for driver.

Example:

    $ ./repl.sh
    REPL ready
    > goreplhowto.Do()
    Hello
    : <nil>
    > goreplhowto.Var
    7
    > ^D
    REPL ready
    > goreplhowto.Var
    8
    >
