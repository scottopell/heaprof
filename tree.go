package main

import (
	"fmt"

	"github.com/remeh/diago/pprof"
	profile "github.com/remeh/diago/profile"
)

type NestedTreeNode struct {
	Children []NestedTreeNode `json:"children"`
	Function profile.Function `json:"function"`
	Self     int64            `json:"self"`
	Value    int64            `json:"value"`
	Percent  float64          `json:"percent"`
}

type TraceTree struct {
	name string
	Root NestedTreeNode `json:"root"`
}

func toNestedNode(node *profile.TreeNode) NestedTreeNode {
	children := make([]NestedTreeNode, 0, len(node.Children))
	treeNode := NestedTreeNode{
		Function: node.Function,
		Self:     node.Self,
		Value:    node.Value,
		Percent:  node.Percent,
	}
	for _, child := range node.Children {
		children = append(children, toNestedNode(child))
	}
	treeNode.Children = children

	return treeNode
}

func toNestedTree(prof *profile.FunctionsTree) TraceTree {
	ret := TraceTree{name: prof.Name, Root: toNestedNode(prof.Root)}
	return ret
}

func printPointerTree(node *profile.TreeNode) {
	fmt.Println(node)
	for _, child := range node.Children {
		printPointerTree(child)
	}
}

func loadTreeProf(prof *pprof.Profile) TraceTree {
	p, err := profile.NewProfile(prof, profile.ModeHeapAlloc)
	if err != nil {
		panic(err)
	}

	fmt.Println("p: ", p)
	pointerTree := p.BuildTree("test profile", false, "")

	fmt.Println("pointerTree: ")
	printPointerTree(pointerTree.Root)
	nestedTree := toNestedTree(pointerTree)

	fmt.Println("nestedTree: ", nestedTree)

	return nestedTree
}
