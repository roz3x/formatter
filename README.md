# formatter 

json-like formatting for node like data structures .

## example

usage: 

```go`
type Node struct {
	Data int
	Name string
	Next *Node
}

func main() {
	fmt.Space = 1
	fmt.Format("hello world")
	start := Node{
		Data: 0,
		Name: "roz3x",
		Next: nil,
	}
	p := &start
	for i := 0; i < 3; i++ {
		n := &Node{
			Data: 10,
			Name: "asdf",
			Next: p,
		}
		p = n
	}
	fmt.Format(p)
}
```

output:

```
"hello world",
{
 "main.Node":{
          "Data":
              10,
          "Name":
              "asdf",
          "Next":
              {
               "main.Node":{
                        "Data":
                            10,
                        "Name":
                            "asdf",
                        "Next":
                            {
                             "main.Node":{
                                      "Data":
                                          10,
                                      "Name":
                                          "asdf",
                                      "Next":
                                          {
                                           "main.Node":{
                                                    "Data":
                                                        0,
                                                    "Name":
                                                        "roz3x",
                                                    "Next":
                                                        null
                                           }
                                          }
                             }
                            }
               }
              }
 }
}

```
