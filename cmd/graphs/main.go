package main

import "fmt"

// Graph Structure

type Graph struct {
	vertices []*Vertex
}

// Vertex Structure
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Add Vertex to Graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}

}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Add Edge to Graph
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge %v --> %v", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge %v --> %v", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// Get Vertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if k == v.key {
			return g.vertices[i]
		}
	}
	return nil
}

// Print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v", v.key)
		}
	}
}

func main() {
	test := &Graph{}

	for i := range 5 {
		test.AddVertex(i)
	}
	test.AddVertex(0)
	test.AddVertex(0)
	test.AddVertex(0)
	test.AddVertex(0)
	test.AddVertex(0)
	test.AddVertex(0)

	test.Print()

}
