# LearnGO

*Lesson 2. Package*
- All package files should be in the same directory;
Example: `$GOPATH/src/github.com/username/learngo/second_lesson/explain/packages/what`

- All files should belong to the same package (_package main_, _package x_, _package y_ etc.);

- Package clause.
You can use package clause in a file to let GO know that which package that file belongs to.
Example:
`package main` --> package clause, this should be the first code and it can only appear once.

*Executable package*
- An executable package go program should belong to package main (go build, go run);
- Executable package you can run it;
- But it's not portable (*non-importable*);
- Name should be `main`;
- Should always contain a `func main`.

*Library package*
- Created for reusability (package fmt etc.);
- Library package can not be executable (*non-executable*);
- Can be only imported from other package;
- Can have any name;
- No `func main`.

*Scope*
Scope means visibility who can see what and who can access to what and so on. There are `package`, `file`, `func` and `block`.
Every line of code can have different scope depending on their position in a GO file.

`package main`

`inport "fmt"  ---> file scoped (only visible in this file)`

`const ok = true ---> package scoped (visible to all the files belong to the package). Other packages can't see then.`
`func main () {---> `
*Declorations*
Declares a unique name bound to a scope.