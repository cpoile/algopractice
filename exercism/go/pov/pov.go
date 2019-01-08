package pov

type Graph struct {
	root string
	adj  map[string][]string
}

func New() *Graph {
	return &Graph{adj: make(map[string][]string)}
}

func (g *Graph) AddNode(nodeLabel string) {}

// This solution uses the fact that a tree is just a undirected graph
// with no cycles.
// So, build an undirected graph, and when outputting the tree
// we just ignore the child->parent links.
func (g *Graph) AddArc(from, to string) {
	g.adj[from] = append(g.adj[from], to)
	g.adj[to] = append(g.adj[to], from)
	g.root = from
}

func (g *Graph) ArcList() []string {
	return g.dfs(g.root, g.root)
}

func (g *Graph) dfs(node, parent string) (ret []string) {
	for _, c := range g.adj[node] {
		// ignore reciprocal links
		if c != parent {
			ret = append(ret, node+" -> "+c)
			ret = append(ret, g.dfs(c, node)...)
		}
	}
	return ret
}

// This is the easy part:
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	g.root = newRoot
	return g
}
