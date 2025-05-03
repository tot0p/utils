# Utils

A collection of utility packages in Go.

## Binary Tree

The `tree` package provides a simple binary tree implementation with the following features:
- Create new nodes with values
- Insert values maintaining binary tree properties
- Search for values in the tree

### Usage

```go
import "github.com/tot0p/utils/tree"

// Create a new tree
root := tree.NewNode(10)

// Insert values
root.Insert(5)
root.Insert(15)

// Search for values
exists := root.Search(5) // returns true
missing := root.Search(20) // returns false
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.