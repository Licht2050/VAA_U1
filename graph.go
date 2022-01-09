package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type Node struct {
	Name  string
	Nodes map[string]*Node
}

func NewNode(node string) *Node {
	return &Node{
		Name:  node,
		Nodes: map[string]*Node{},
	}
}

func (g *Graph) AddNode(node string) {
	//if node already exist; then return
	if _, ok := g.Nodes[node]; ok {
		return
	}

	node1 := NewNode(node)
	g.Nodes[node] = node1
}

func (g *Graph) RemoveEdge(node string) {
	g.Nodes[node].Nodes = make(map[string]*Node)
	// delete(g.Nodes, node)
}

func (g *Graph) RemoveNode(node string) {
	//first remove all edges
	g.RemoveEdge(node)
	delete(g.Nodes, node)
}

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: map[string]*Node{},
	}
}

func (g *Graph) AddEdge(nodeFrom, nodeTo string) {
	node1 := g.Nodes[nodeFrom]
	node2 := g.Nodes[nodeTo]

	if node1 == nil || node2 == nil {
		panic("Nodes are not exist")
	}

	if _, ok := node1.Nodes[node2.Name]; ok {
		return
	}

	node1.Nodes[node2.Name] = node2
	g.Nodes[node1.Name] = node1
	// g.nodes[from] = append(g.nodes[from], edge{node: to, label: label})
}

func (g *Graph) GetEdges(node string) *Node {
	return g.Nodes[node]
}

func (g *Graph) String() string {
	out := `digraph Graph {
		rankdir=UD;
		size="8,5"
		node [shape = circle];`
	out += "\n"
	for k := range g.Nodes {
		for _, v := range g.GetEdges(k).Nodes {
			out += fmt.Sprintf("\t%s -> %s\n", k, v.Name)
		}
	}
	out += "}"
	return out
}

func (g *Graph) ParseStringToDiG(data string) {
	if len(strings.TrimSpace(data)) == 0 {
		panic("Variable is empty!")
	}

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		str := regexp.MustCompile(`[\w\d]+.\->.[\w\d]+`)

		for _, element := range str.FindAllString(line, -1) {
			g.AddNode(strings.Split(element, "->")[0])
			g.AddNode(strings.Split(element, "->")[1])
			g.AddEdge(strings.Split(element, "->")[0], strings.Split(element, "->")[1])
		}

	}
}

func (g *Graph) ParseFileToGraph(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	g.ParseStringToDiG(string(data))
}

type Person struct {
	Name    string
	Vorname string
}

type Test struct {
	People []Person
}

func (t *Test) test() {

	t.People = append(t.People, Person{Name: "Yousuf", Vorname: "Yousuf"})
}

func TestStruct(value interface{}) {
	if msg, ok := value.(*Person); ok {
		fmt.Println(msg.Name)
		msg.Name = "new new"
	}

}

func main() {
	// var test *Test

	p := new(Person)
	p.Name = " old old"
	TestStruct(p)

	// people.People = append(people.People, Person{Name: "Ahmad", Vorname: "Mosa"})

	// people.test()

	// for _, p := range people.People {
	// 	fmt.Println(p.Name)
	// }

	fmt.Println(p.Name)

	// g := NewGraph()
	// path := "mygraph.dot"
	// g.ParseFileToGraph(path)
	// fmt.Println(g.String())
	// re, err := regexp.Compile(`[^\w]`)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// text := ""
	// scanner := bufio.NewScanner(strings.NewReader(b1))
	// for scanner.Scan() {
	// 	text = scanner.Text()
	// 	str := regexp.MustCompile(`[\w\d]+.\->`)
	// 	str1 := str.MatchString(text)
	// 	fmt.Println(str1)

	// 	for _, element := range str.FindAllString(text, -1) {
	// 		fmt.Println(element)
	// 	}

	// }
	// b2 := strings.Trim(b1, "_")

	// g.AddNode("node01")
	// g.AddNode("node02")
	// g.AddNode("node03")
	// g.AddNode("node04")
	// g.AddNode("node05")
	// g.AddNode("node06")
	// g.AddNode("node07")
	// g.AddNode("node08")

	// g.AddEdge("node01", "node02")
	// g.AddEdge("node01", "node03")
	// g.AddEdge("node02", "node06")
	// g.AddEdge("node02", "node08")
	// g.AddEdge("node03", "node05")
	// g.AddEdge("node03", "node02")
	// g.AddEdge("node04", "node07")
	// g.AddEdge("node04", "node06")
	// g.AddEdge("node05", "node07")
	// g.AddEdge("node06", "node08")

	// fmt.Println(str1)

}
