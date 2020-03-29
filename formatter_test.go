package formatter


import (
    "testing"
)


type Node struct {
    Name string
    Next *Node
}

func TestFormat( _ *testing.T ) {
    var t interface{}
    t = "asdf"
    Format(t)
    t = 67
    Format(t)
    t = &Node{
        Name : "adsf",
        Next : nil, 
    }
    Format(t)
}
