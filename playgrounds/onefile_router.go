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
48
49
50
51
52
53
54
55
56
57
58
59
60
61
62
63
64
65
66
67
68
69
70
71
72
73
74
75
76
77
78
79
80
81
82
83
84
85
86
87
88
89
90
91
92
93
94
95
96
97
98
99
100
101
102
103
104
105
106
107
108
109
110
111
112
113
114
115
116
117
118
119
120
121
122
123
124
125
126
127
128
129
130
131
132
133
134
135
136
137
138
139
140
141
142
143
144
145
146
147
148
149
150
151
152
153
154
155
156
157
158
package main

import (
	"fmt"
	"github.com/gyozatech/temaki"
)

func main() {	
	
	err := temaki.NewRouter().
	                 UseMiddleware(exampleMiddleware).
	                 PATCH ("/api/widgets/([^/]+)/parts/([0-9]+)", apiUpdateWidgetPart).
	                 DELETE("/api/widgets/([^/]+)/parts/([0-9]+)", apiDeleteWidgetPart).
	                 Start(8080)
	if err != nil {            
	    log.Fatal(err)
	}
}

func exampleMiddleware(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Executing middleware before request phase!")
        handler.ServeHTTP(w, r)
        fmt.Println("Executing middleware after request phase!")
    })
}

func apiUpdateWidgetPart(w http.ResponseWriter, r *http.Request) {
	slug := temaki.GetField(r, 0)
	id, _ := strconv.Atoi(temaki.GetField(r, 1))
	fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", slug, id)
}

func apiDeleteWidgetPart(w http.ResponseWriter, r *http.Request) {
	slug := temaki.GetField(r, 0)
	id, _ := strconv.Atoi(temaki.GetField(r, 1))
	fmt.Fprintf(w, "apiDeleteWidgetPart %s %d\n", slug, id)
}


// ~~~~~~~~~~



type Route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func NewRoute(method, pattern string, handler http.HandlerFunc) route {
	return Route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type Middleware func(handler http.Handler) http.Handler

type ctxKey struct{}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


type Router struct {
   Routes []Route
   Middlewares []Middleware
}

func NewRouter() *Router {
   return &Router{ []Route{}, []Middleware{} }
}

func (router *Router) UseMiddleware(middleware Middleware) *Router {
   router.Middlewares = append(router.Middlewares, middleware)
   return router
}

func (router *Router) GET(pattern string, handlerFunc http.HandlerFunc) *Router {
   router.Routes = append(router.Routes, newRoute("GET", pattern, handlerFunc))
   return router 
}

func (router *Router) POST(pattern string, handlerFunc http.HandlerFunc) *Router {
   router.Routes = append(router.Routes, newRoute("OPTIONS", pattern, handlerFunc))
   router.Routes = append(router.Routes, newRoute("POST", pattern, handlerFunc))
   return router 
}

func (router *Router) PUT(pattern string, handlerFunc http.HandlerFunc) *Router {
   router.Routes = append(router.Routes, newRoute("OPTIONS", pattern, handlerFunc))
   router.Routes = append(router.Routes, newRoute("PUT", pattern, handlerFunc))
   return router 
}

func (router *Router) PATCH(pattern string, handlerFunc http.HandlerFunc) *Router {
   router.Routes = append(router.Routes, newRoute("OPTIONS", pattern, handlerFunc))
   router.Routes = append(router.Routes, newRoute("PATCH", pattern, handlerFunc))
   return router 
}

func (router *Router) DELETE(pattern string, handlerFunc http.HandlerFunc) *Router {
   router.Routes = append(router.Routes, newRoute("OPTIONS", pattern, handlerFunc))
   router.Routes = append(router.Routes, newRoute("DELETE", pattern, handlerFunc))
   return router 
}

func (router *Router) OPTIONS(pattern string, handlerFunc http.HandlerFunc) *Router {
   router.Routes = append(router.Routes, newRoute("OPTIONS", pattern, handlerFunc))
   return router 
}

func (router *Router) HEAD(pattern string, handlerFunc http.HandlerFunc) *Router {
   router.Routes = append(router.Routes, newRoute("HEAD", pattern, handlerFunc))
   return router 
}

func (router *Router) DispatcherHandler() http.HandlerFunc {
   return func(w http.ResponseWriter, r *http.Request) {
        var allow []string
	for _, route := range router.Routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
   }
}

func (router *Router) Serve() http.Handler {
    var finalHandler http.Handler = http.HandlerFunc(router.DispatcherHandler())
    for _, middleware := router.Middlewares {
        finalHandler = middleware(finalHandler)
    }
    return finalHandler
}

func (router *Router) Start(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), router.Serve())
}

// ~~~~~~~~~~~~

func GetField(r *http.Request, index int) string {
	fields := r.Context().Value(ctxKey{}).([]string)
	return fields[index]
}


go: finding module for package github.com/gyozatech/temaki
go: downloading github.com/gyozatech/temaki v0.0.0-20220502135914-9545ea1be471
./prog.go:141:23: syntax error: cannot use _, middleware := router.Middlewares as value

Go build failed.
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
