▲
😐
Go
Why Go
Get Started
Docs
Packages
Play
Blog
Go.
Why Go
Get Started
Docs
Packages
Play
Blog
The Go Playground

Go 1.18
Run
Format Share

Hello, World!
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
35
36
37
38
39
40
41
42
43
44
45
46
47
package main

import (
	"fmt"
)

// The ways to create a pointer on the fly
func main() {

	// create a pointer with new()
	var p *int = new(int)
	*p = 5
	fmt.Println(*p)

	// create a pointer with anonymous function
	var h *int = func(i int) *int { return &i }(5)
	fmt.Println(*h)

	// create a pointer with a slice
	var k *int = &[]int{5}[0]
	fmt.Println(*k)

	// create a pointer with an helper function
	var s *int = pointerOfInt(5)
	fmt.Println(*s)

	// create a pointer from a variable
	var t int = 5
	var j *int = &t
	fmt.Println(*j)
}

func pointerOfInt(i int) *int {
	return &i
}

5
5
5
5
5

Program exited.
About the Playground

The Go Playground is a web service that runs on go.dev's servers. The service receives a Go program, vets, compiles, links, and runs the program inside a sandbox, then returns the output.

If the program contains tests or examples and no main function, the service runs the tests. Benchmarks will likely not be supported since the program runs in a sandboxed environment with limited resources.

There are limitations to the programs that can be run in the playground:

The playground can use most of the standard library, with some exceptions. The only communication a playground program has to the outside world is by writing to standard output and standard error.
In the playground the time begins at 2009-11-10 23:00:00 UTC (determining the significance of this date is an exercise for the reader). This makes it easier to cache programs by giving them deterministic output.
There are also limits on execution time and on CPU and memory usage.
The article "Inside the Go Playground" describes how the playground is implemented. The source code is available at https://go.googlesource.com/playground.

The playground uses the latest stable release of Go.

The playground service is used by more than just the official Go project (Go by Example is one other instance) and we are happy for you to use it on your own site. All we ask is that you contact us first (note this is a public mailing list), that you use a unique user agent in your requests (so we can identify you), and that your service is of benefit to the Go community.

Any requests for content removal should be directed to security@golang.org. Please include the URL and the reason for the request.

Why Go
Use Cases
Case Studies
Get Started
Playground
Tour
Stack Overflow
Help
Packages
Standard Library
About
Download
Blog
Issue Tracker
Release Notes
Brand Guidelines
Code of Conduct
Connect
Twitter
GitHub
Slack
r/golang
Meetup
Golang Weekly
The Go Gopher
Copyright
Terms of Service
Privacy Policy
Report an Issue
Light theme
Google logo
