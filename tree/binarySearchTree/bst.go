package binarySearchTree

type treeNode struct {
	value     int
	leftNode  *treeNode
	rightNode *treeNode
}

type BST struct {
	root *treeNode
	size int
}

func NewBST() *BST {
	return &BST{
		size: 0,
	}
}

func (b *BST) AddIterative(value int) bool {

	newNode := &treeNode{
		value:     value,
		leftNode:  nil,
		rightNode: nil,
	}

	if b.IsEmpty() {
		b.root = newNode
		b.size++
		return true
	}

	temp := b.root
	for temp != nil {
		if value <= temp.value {
			if temp.leftNode == nil {
				temp.leftNode = newNode
				b.size++
				return true
			}
			temp = temp.leftNode
		} else {
			if temp.rightNode == nil {
				temp.rightNode = newNode
				b.size++
				return true
			}
			temp = temp.rightNode
		}
	}
	return false
}

func (b *BST) AddRecursive(value int) bool {
	b.root = b.addRecursive(b.root, value)
	b.size++
	return true
}

func (b *BST) addRecursive(current *treeNode, value int) *treeNode {
	if current == nil {
		current = &treeNode{
			value:     value,
			leftNode:  nil,
			rightNode: nil,
		}
	} else if value <= current.value {
		current.leftNode = b.addRecursive(current.leftNode, value)
	} else {
		current.rightNode = b.addRecursive(current.rightNode, value)
	}
	return current
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

func (b *BST) Size() int {
	return b.size
}

func (b *BST) SearchIterative(value int) bool {
	if b.IsEmpty() {
		return false
	}

	temp := b.root
	for temp != nil {
		if value == temp.value {
			return true
		}

		if value <= temp.value {
			temp = temp.leftNode
		} else {
			temp = temp.rightNode
		}
	}

	return false
}

func (b *BST) SearchRecursive(value int) bool {
	if b.IsEmpty() {
		return false
	}

	return b.searchRecursive(b.root, value)
}

func (b *BST) searchRecursive(node *treeNode, value int) bool {
	if node == nil {
		return false
	}
	if value == node.value {
		return true
	} else if value <= node.value {
		return b.searchRecursive(node.leftNode, value)
	} else {
		return b.searchRecursive(node.rightNode, value)
	}
}

func (b *BST) Delete(value int) bool {
	if b.IsEmpty() {
		return false
	}

	var parent *treeNode
	current := b.root
	for current != nil && value != current.value {
		parent = current
		if value <= current.value {
			current = current.leftNode
		} else {
			current = current.rightNode
		}
	}

	if current == nil {
		return false
	}

	if current.leftNode == nil && current.rightNode == nil {

		if current == b.root {
			b.root = nil
		} else if value <= parent.value {
			parent.leftNode = nil
		} else {
			parent.rightNode = nil
		}

		b.size--
		return true
	} else if current.leftNode != nil && current.rightNode != nil {

		mostLeft := b.mostLeft(current.rightNode)
		temp := mostLeft.value
		b.Delete(temp)
		current.value = temp

		return true
	} else {

		if value <= parent.value {
			if current.leftNode != nil {
				parent.leftNode = current.leftNode
			} else {
				parent.leftNode = current.rightNode
			}
		} else {
			if current.leftNode != nil {
				parent.rightNode = current.leftNode
			} else {
				parent.rightNode = current.rightNode
			}
		}

		b.size--
		return true
	}
}

func (b *BST) mostLeft(node *treeNode) treeNode {

	temp := node
	for temp.leftNode != nil {
		temp = temp.leftNode
	}

	return *temp
}

func (b *BST) PreOrderTraversal() *[]int {
	if b.IsEmpty() {
		return nil
	}

	data := new([]int)
	b.preOrderTraversal(b.root, data)
	return data
}

func (b *BST) preOrderTraversal(node *treeNode, data *[]int) {
	*data = append(*data, node.value)
	if node.leftNode != nil {
		b.preOrderTraversal(node.leftNode, data)
	}
	if node.rightNode != nil {
		b.preOrderTraversal(node.rightNode, data)
	}
}

func (b *BST) InOrderTraversal() *[]int {
	if b.IsEmpty() {
		return nil
	}

	data := new([]int)
	b.inOrderTraversal(b.root, data)
	return data
}

func (b *BST) inOrderTraversal(node *treeNode, data *[]int) {

	if node.leftNode != nil {
		b.inOrderTraversal(node.leftNode, data)
	}

	*data = append(*data, node.value)

	if node.rightNode != nil {
		b.inOrderTraversal(node.rightNode, data)
	}
}

func (b *BST) PostOrderTraversal() *[]int {
	if b.IsEmpty() {
		return nil
	}

	data := new([]int)
	b.postOrderTraversal(b.root, data)
	return data
}

func (b *BST) postOrderTraversal(node *treeNode, data *[]int) {

	if node.leftNode != nil {
		b.postOrderTraversal(node.leftNode, data)
	}

	if node.rightNode != nil {
		b.postOrderTraversal(node.rightNode, data)
	}

	*data = append(*data, node.value)
}
