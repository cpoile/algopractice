pub mod graph {
    pub mod graph_items {
        pub mod edge {
            pub struct Edge {
                from: String,
                to: String,
            }

            impl Edge {
                pub fn new(f: &str, t: &str) -> Self {
                    Edge {
                        from: f.to_string(),
                        to: t.to_string(),
                    }
                }
            }
        }

        pub mod node {
            #[derive(Debug, Eq, PartialEq, Clone)]
            pub struct Node {
                name: String,
            }

            impl Node {
                pub fn new(n: &str) -> Self {
                    Node {
                        name: n.to_string(),
                    }
                }
            }
        }

        pub mod attrs {
            pub struct Attrs {
                attrs: Vec<String>,
            }

            impl Attrs {
                pub fn new(att: Vec<&str>) -> Self {
                    let mut attr: Vec<_> = Vec::new();
                    for a in att {
                        attr.push(a.to_string());
                    }
                    Attrs { attrs: attr }
                }
            }
        }
    }

    use graph_items::attrs::Attrs;
    use graph_items::edge::Edge;
    use graph_items::node::Node;

    pub struct Graph {
        pub nodes: Vec<Node>,
        pub edges: Vec<Edge>,
        pub attrs: Vec<Attrs>,
    }

    impl Graph {
        pub fn new() -> Self {
            Graph {
                nodes: vec![],
                edges: vec![],
                attrs: vec![],
            }
        }

        pub fn with_nodes<'a>(&'a mut self, nodes: &[Node]) -> &'a mut Self {
            let mut orig = nodes.to_vec();
            self.nodes.append(&mut orig);
            self
        }
    }
}
