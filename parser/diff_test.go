package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
)

func ExampleDiffParser() {
	const sample = `diff --git a/gofmt.go b/gofmt.go
--- a/gofmt.go	2020-07-26 08:01:09.260800318 +0000
+++ b/gofmt.go	2020-07-26 08:01:09.260800318 +0000
@@ -1,6 +1,6 @@
 package testdata
 
-func    fmt     () {
+func fmt() {
 	// test
 	// test line
 	// test line
@@ -10,11 +10,11 @@
 	// test line
 	// test line
 
-println(
-		"hello, gofmt test"    )
-//comment
+	println(
+		"hello, gofmt test")
+	//comment
 }
 
+type s struct{ A int }
 
-type s struct { A int }
 func (s s) String() { return "s" }
`
	const strip = 1
	p := NewDiffParser(strip)
	diagnostics, err := p.Parse(strings.NewReader(sample))
	if err != nil {
		panic(err)
	}
	for _, d := range diagnostics {
		rdjson, _ := protojson.MarshalOptions{Indent: "  "}.Marshal(d)
		var out bytes.Buffer
		json.Indent(&out, rdjson, "", "  ")
		fmt.Println(out.String())
	}
	// Output:
	// {
	//   "location": {
	//     "path": "gofmt.go",
	//     "range": {
	//       "start": {
	//         "line": 3
	//       },
	//       "end": {
	//         "line": 3
	//       }
	//     }
	//   },
	//   "suggestions": [
	//     {
	//       "range": {
	//         "start": {
	//           "line": 3
	//         },
	//         "end": {
	//           "line": 3
	//         }
	//       },
	//       "text": "func fmt() {"
	//     }
	//   ],
	//   "originalOutput": "gofmt.go:3:-func    fmt     () {\ngofmt.go:3:+func fmt() {"
	// }
	// {
	//   "location": {
	//     "path": "gofmt.go",
	//     "range": {
	//       "start": {
	//         "line": 13
	//       },
	//       "end": {
	//         "line": 15
	//       }
	//     }
	//   },
	//   "suggestions": [
	//     {
	//       "range": {
	//         "start": {
	//           "line": 13
	//         },
	//         "end": {
	//           "line": 15
	//         }
	//       },
	//       "text": "\tprintln(\n\t\t\"hello, gofmt test\")\n\t//comment"
	//     }
	//   ],
	//   "originalOutput": "gofmt.go:13:-println(\ngofmt.go:14:-\t\t\"hello, gofmt test\"    )\ngofmt.go:15:-//comment\ngofmt.go:13:+\tprintln(\ngofmt.go:14:+\t\t\"hello, gofmt test\")\ngofmt.go:15:+\t//comment"
	// }
	// {
	//   "location": {
	//     "path": "gofmt.go",
	//     "range": {
	//       "start": {
	//         "line": 18,
	//         "column": 1
	//       },
	//       "end": {
	//         "line": 18,
	//         "column": 1
	//       }
	//     }
	//   },
	//   "suggestions": [
	//     {
	//       "range": {
	//         "start": {
	//           "line": 18,
	//           "column": 1
	//         },
	//         "end": {
	//           "line": 18,
	//           "column": 1
	//         }
	//       },
	//       "text": "type s struct{ A int }\n"
	//     }
	//   ],
	//   "originalOutput": "gofmt.go:18:+type s struct{ A int }"
	// }
	// {
	//   "location": {
	//     "path": "gofmt.go",
	//     "range": {
	//       "start": {
	//         "line": 19
	//       },
	//       "end": {
	//         "line": 19
	//       }
	//     }
	//   },
	//   "suggestions": [
	//     {
	//       "range": {
	//         "start": {
	//           "line": 19
	//         },
	//         "end": {
	//           "line": 19
	//         }
	//       }
	//     }
	//   ],
	//   "originalOutput": "gofmt.go:19:-type s struct { A int }"
	// }
}

func ExampleDiffParser_addNewLine() {
	const sample = `diff --git a/newline.txt b/newline.txt
--- a/newline.txt	2024-10-10 20:15:37.618432000 +0900
+++ b/newline.txt	2024-10-10 20:15:02.110606546 +0900
@@ -1,2 +1,2 @@
 No newline at end of the old file only
-a
\ No newline at end of file
+a
`
	const strip = 1
	p := NewDiffParser(strip)
	diagnostics, err := p.Parse(strings.NewReader(sample))
	if err != nil {
		panic(err)
	}
	for _, d := range diagnostics {
		rdjson, _ := protojson.MarshalOptions{Indent: "  "}.Marshal(d)
		var out bytes.Buffer
		json.Indent(&out, rdjson, "", "  ")
		fmt.Println(out.String())
	}
	// Output:
	// {
	//   "location": {
	//     "path": "newline.txt",
	//     "range": {
	//       "start": {
	//         "line": 2
	//       },
	//       "end": {
	//         "line": 2
	//       }
	//     }
	//   },
	//   "suggestions": [
	//     {
	//       "range": {
	//         "start": {
	//           "line": 2
	//         },
	//         "end": {
	//           "line": 2
	//         }
	//       },
	//       "text": "a\n"
	//     }
	//   ],
	//   "originalOutput": "newline.txt:2:-a\nnewline.txt:2:+a"
	// }
}

func ExampleDiffParser_removeNewLine() {
	const sample = `diff --git a/newline.txt b/newline.txt
--- a/newline.txt	2024-10-10 20:15:37.618432000 +0900
+++ b/newline.txt	2024-10-10 20:15:02.110606546 +0900
@@ -1,2 +1,2 @@
 No newline at end of the new file only
-a
+a
\ No newline at end of file
`
	const strip = 1
	p := NewDiffParser(strip)
	diagnostics, err := p.Parse(strings.NewReader(sample))
	if err != nil {
		panic(err)
	}
	for _, d := range diagnostics {
		rdjson, _ := protojson.MarshalOptions{Indent: "  "}.Marshal(d)
		var out bytes.Buffer
		json.Indent(&out, rdjson, "", "  ")
		fmt.Println(out.String())
	}
	// Output:
	// {
	//   "location": {
	//     "path": "newline.txt",
	//     "range": {
	//       "start": {
	//         "line": 2
	//       },
	//       "end": {
	//         "line": 2
	//       }
	//     }
	//   },
	//   "suggestions": [
	//     {
	//       "range": {
	//         "start": {
	//           "line": 2
	//         },
	//         "end": {
	//           "line": 2
	//         }
	//       },
	//       "text": "a"
	//     }
	//   ],
	//   "originalOutput": "newline.txt:2:-a\nnewline.txt:2:+a"
	// }
}

func ExampleDiffParser_keepNoNewLine() {
	const sample = `diff --git a/newline.txt b/newline.txt
--- a/newline.txt	2024-10-10 20:15:37.618432000 +0900
+++ b/newline.txt	2024-10-10 20:15:02.110606546 +0900
@@ -1,2 +1,2 @@
 No newline at end of both the old and new file
-a
\ No newline at end of file
+b
\ No newline at end of file
`
	const strip = 1
	p := NewDiffParser(strip)
	diagnostics, err := p.Parse(strings.NewReader(sample))
	if err != nil {
		panic(err)
	}
	for _, d := range diagnostics {
		rdjson, _ := protojson.MarshalOptions{Indent: "  "}.Marshal(d)
		var out bytes.Buffer
		json.Indent(&out, rdjson, "", "  ")
		fmt.Println(out.String())
	}
	// Output:
	// {
	//   "location": {
	//     "path": "newline.txt",
	//     "range": {
	//       "start": {
	//         "line": 2
	//       },
	//       "end": {
	//         "line": 2
	//       }
	//     }
	//   },
	//   "suggestions": [
	//     {
	//       "range": {
	//         "start": {
	//           "line": 2
	//         },
	//         "end": {
	//           "line": 2
	//         }
	//       },
	//       "text": "b"
	//     }
	//   ],
	//   "originalOutput": "newline.txt:2:-a\nnewline.txt:2:+b"
	// }
}

func ExampleDiffParser_noChangeLastLine() {
	const sample = `diff --git a/newline.txt b/newline.txt
--- a/newline.txt	2024-10-10 20:15:37.618432000 +0900
+++ b/newline.txt	2024-10-10 20:15:02.110606546 +0900
@@ -1,3 +1,3 @@
 No newline at end of both the old and new file
-a
+b
 Last line
\ No newline at end of file
`
	const strip = 1
	p := NewDiffParser(strip)
	diagnostics, err := p.Parse(strings.NewReader(sample))
	if err != nil {
		panic(err)
	}
	for _, d := range diagnostics {
		rdjson, _ := protojson.MarshalOptions{Indent: "  "}.Marshal(d)
		var out bytes.Buffer
		json.Indent(&out, rdjson, "", "  ")
		fmt.Println(out.String())
	}
	// Output:
	// {
	//   "location": {
	//     "path": "newline.txt",
	//     "range": {
	//       "start": {
	//         "line": 2
	//       },
	//       "end": {
	//         "line": 2
	//       }
	//     }
	//   },
	//   "suggestions": [
	//     {
	//       "range": {
	//         "start": {
	//           "line": 2
	//         },
	//         "end": {
	//           "line": 2
	//         }
	//       },
	//       "text": "b"
	//     }
	//   ],
	//   "originalOutput": "newline.txt:2:-a\nnewline.txt:2:+b"
	// }
}
