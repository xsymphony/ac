package ac

type node struct {
    next map[rune]*node // 子节点
    fail *node          // 失效函数
    isWord bool         // 是否是单词
}

func newNode() *node {
    return &node{
        next: make(map[rune]*node),
    }
}

type Automaton struct {
    root *node
}

func NewAutomaton() *Automaton {
    return &Automaton{
        root:newNode(),
    }
}

// Add word to trie tree
func (m *Automaton) Add(word string) {
    node := m.root
    runes := []rune(word)
    for _, r := range runes {
        if _, ok := node.next[r]; !ok {
            node.next[r] = newNode()
        }
        node = node.next[r]
    }
    node.isWord = true
}

// Build AC failure function
func (m *Automaton) Build() {
    var queue []*node
    queue = append(queue, m.root)
    for len(queue) != 0 {
        parent := queue[0]
        queue = queue[1:]

        for char, child := range parent.next {
            if parent == m.root {
                child.fail = m.root
            } else {
                failAcNode := parent.fail
                for failAcNode != nil {
                    if _, ok := failAcNode.next[char]; ok {
                        child.fail = parent.fail.next[char]
                        break
                    }
                    failAcNode = failAcNode.fail
                }
                if failAcNode == nil {
                    child.fail = m.root
                }
            }
            queue = append(queue, child)
        }
    }
}

// Find returns matched words with index
func (m *Automaton) Find(sentence string) (words []string, index []int) {
    var start, end int
    runes := []rune(sentence)
    node := m.root
    for i, r := range runes {
        if _, ok := node.next[r]; !ok && node != m.root {
            node = node.fail
        }

        if _, ok := node.next[r]; ok {
            if node == m.root {
                start = i
            }

            node = node.next[r]
            if node.isWord {
                end = i
                words = append(words, string(runes[start:end+1]))
                index = append(index, start, end)
            }
        }
    }

    return
}

