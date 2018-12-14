// Package tree provides a function to build a tree
// based on the specifications of the tests.
package tree

import (
	"errors"
	"fmt"
	"sort"
)

// Record is the input format for the tree
type Record struct {
	ID, Parent int
}

// Node defines each node of the tree
type Node struct {
	ID       int
	Children []*Node
}

// Build checks that the records are in the proper format and constructs the tree
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// sort because it costs nlogn and cuts many loops out of what we have to do in the future
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	// NOTE: I have formatted the error checking to reflect the tests

	// check that root node exists
	if records[0].ID != 0 {
		return nil, errors.New("no root node")
	}

	// check that root node has no parent
	if records[0].Parent > 0 {
		return nil, errors.New("root node has parent")
	}

	// start the tree with the root node
	root := &Node{}
	nodes := []*Node{root}

	// check for duplicates (including duplicate root), check that we have continuous IDs,
	// check that children have lower IDs than parents, check for direct cycle
	// NOTE: although we should have to check for indirect cycle after we've built the tree,
	//       the "higher id parent of lower id" catches indirect cycles
	// after the checks, build the tree
	for i := 1; i < len(records); i++ {
		if records[i-1].ID == records[i].ID {
			return nil, fmt.Errorf("Node is a duplicate - {ID: %d, Parent: %d}", records[i].ID, records[i].Parent)
		}
		if records[i-1].ID != records[i].ID-1 {
			return nil, errors.New("non-continuous")
		}
		if records[i].ID < records[i].Parent {
			return nil, errors.New("higher id parent of lower id")
		}
		if records[i].ID == records[i].Parent {
			return nil, errors.New("direct cycle")
		}

		// add the node to the node list, and add it to its parent
		n := &Node{i, nil}
		nodes = append(nodes, n)
		nodes[records[i].Parent].Children = append(nodes[records[i].Parent].Children, n)
	}

	return root, nil
}
