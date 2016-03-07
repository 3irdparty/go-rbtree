go-rbtree
=================

## About ##
A generic rbtree implemented in Go (Golang).

## Usage ##
Import the package from GitHub first.

```
import "github.com/lsword/go-rbtree/rbtree" rbtree
```
Then you can start having fun.

```
//Create a rbtree
rbtree := rbtree.NewRBTree()

//Insert node into a rbtree
rbtree.Insert(int64(123), 10)

//Delete node from a rbtree
rbtree.Delete(int64(123))

//Print a rbtree
rbtree.PrintTree()

//Get min node
rbtree.Min()

//Get max node
rbtree.Max()

//Search a node by key
rbtree.Search(int64(123))

//Get key of a node 
rbtree.Search(int64(123)).Key()

//Get Value of a node
rbtree.Search(int64(123)).Value()

//Get Next node
rbtree.NextOf(rbtree.Search(int64(123)))

//Get Prev node
rbtree.PrevOf(rbtree.Search(int64(123)))
```
## License ##

MIT, check the `LICENSE` file.
