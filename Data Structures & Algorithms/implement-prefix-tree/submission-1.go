type PrefixTree struct {
    children map[rune]*PrefixTree
    isEnd bool
}

func Constructor() *PrefixTree {
    return &PrefixTree{children: make(map[rune]*PrefixTree), isEnd: false}
}

func (this *PrefixTree) Insert(word string) {
   node := this
    for _, v := range word {
        nextNode, exists := node.children[v]
        if !exists {
            node.children[v] = Constructor()
            nextNode = node.children[v]
        }
        node = nextNode  
    }
            node.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
    node := this
    for _, v := range word {
        nextNode, exists := node.children[v]
        if !exists {
            return false
        }
        node = nextNode 
    }
    if !node.isEnd {
        return false
    }
    return true
}

func (this *PrefixTree) StartsWith(prefix string) bool {
    node := this
    for _, v := range prefix {
        nextNode, exists := node.children[v]
        if !exists {
            return false
        }
        node = nextNode 
    }
    return true
}
