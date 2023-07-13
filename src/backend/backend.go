package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"github.com/gorilla/mux"
	"fmt"
	// "strconv"
)

const port = ":8080"

type spaHandler struct {
	staticPath string
	indexPath  string
}

func createAdjacencyMatrix(input string) [][]bool {
	// Split the input into separate edges
	edges := strings.Split(input, "\n")

	// Find the maximum vertex label to determine the size of the matrix
	maxVertex := 0
	for _, edge := range edges {
		vertices := strings.Split(edge, " ")
		for _, vertex := range vertices {
			if int(vertex[0]-'A')+1 > maxVertex {
				maxVertex = int(vertex[0]-'A') + 1
			}
		}
	}

	// Create the adjacency matrix with all entries set to false
	matrix := make([][]bool, maxVertex)
	for i := range matrix {
		matrix[i] = make([]bool, maxVertex)
	}

	// Populate the matrix based on the edges
	for _, edge := range edges {
		vertices := strings.Split(edge, " ")
		from := int(vertices[0][0] - 'A')
		to := int(vertices[1][0] - 'A')
		matrix[from][to] = true
	}

	return matrix
}

var (
	index      int
	lowLink    []int
	onStack    []bool
	indexStack []int
	disc       []int
	sccs       [][]int
	bridges    [][]int
)

func tarjansSCC(adjMatrix [][]bool) ([][]int, [][]int) {
	n := len(adjMatrix)
	index = 0
	lowLink = make([]int, n)
	disc = make([]int, n)
	for i := 0; i < n; i++ {
		lowLink[i] = -1
		disc[i] = -1
	}
	onStack = make([]bool, n)
	indexStack = []int{}
	sccs = [][]int{}
	bridges = [][]int{}

	for v := 0; v < n; v++ {
		if lowLink[v] == -1 {
			tarjanDFS(v, -1, adjMatrix)
		}
	}

	return bridges, sccs
}

func tarjanDFS(v int, parent int, adjMatrix [][]bool) {
	index++
	lowLink[v] = index
	disc[v] = index
	onStack[v] = true
	indexStack = append(indexStack, v)
	for w := 0; w < len(adjMatrix); w++ {
		if adjMatrix[v][w] {
			if disc[w] == -1 {
				tarjanDFS(w, v, adjMatrix)
				lowLink[v] = min(lowLink[v], lowLink[w])
				if lowLink[w] > disc[v] {
					bridges = append(bridges, []int{v, w})
				}
			} else {
				lowLink[v] = min(lowLink[v], disc[w])
			}
		}
	}
	// fmt.Println("OKK " + strconv.Itoa(v) + " " + strconv.Itoa(lowLink[v]) + " " + strconv.Itoa(index))
	if lowLink[v] == disc[v] {
		scc := []int{}
		w := -1
		for w != v {
			w = indexStack[len(indexStack)-1]
			indexStack = indexStack[:len(indexStack)-1]
			onStack[w] = false
			scc = append(scc, w)
		}
		sccs = append(sccs, scc)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func usernameHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Username string `json:"username"`
	}
	user := User{os.Getenv("USERNAME")}
	p, _ := json.Marshal(user)
	w.Write(p)
}

type SCCResponse struct {
	Bridges [][]string `json:"bridges"`
	SCCs    [][]string `json:"sccs"`
}

type SCCRequest struct {
	Input string `json:"input"`
}

func sccHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into the SCCRequest struct
	var request SCCRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Failed to parse input", http.StatusBadRequest)
		return
	}

	// Extract the input value from the struct
	input := request.Input

	// Create the adjacency matrix
	matrix := createAdjacencyMatrix(input)

	// Execute Tarjan's algorithm
	bridges, sccs := tarjansSCC(matrix)

	// Create the response objects for nodes and edges
var nodes []map[string]interface{}
var edges []map[string]interface{}

// Process edges from the adjacency matrix
for from := 0; from < len(matrix); from++ {
	for to := 0; to < len(matrix[from]); to++ {
		if matrix[from][to] {
			edge := map[string]interface{}{
				"from":    string(from + 'A'),
				"to":      string(to + 'A'),
				"bridges": false,
				"color":   "red",
			}
			edges = append(edges, edge)
		}
	}
}

// Process bridges
for _, bridge := range bridges {
	fromNode := string(bridge[0] + 'A')
	toNode := string(bridge[1] + 'A')

	// Find the corresponding edge in the edges slice
	for i := range edges {
		if edges[i]["from"] == fromNode && edges[i]["to"] == toNode {
			// Set the bridge property to true
			edges[i]["bridges"] = true
			edges[i]["color"] = "blue"
			break
		}
	}
}

// Process SCCs
for i, scc := range sccs {
	component := i

	for _, vertex := range scc {
		node := map[string]interface{}{
			"id":        string(vertex + 'A'),
			"label":     fmt.Sprintf("Node %s", string(vertex+'A')),
			"component": component,
		}

		nodes = append(nodes, node)
	}
}

// Combine nodes and edges into the final response format
response := map[string]interface{}{
	"nodes": nodes,
	"edges": edges,
}

// Convert the response to JSON
responseJSON, err := json.Marshal(response)
if err != nil {
	http.Error(w, "Failed to create response", http.StatusInternalServerError)
	return
}

	// // Convert the bridges and sccs to string representations
	// bridgeStrings := make([][]string, len(bridges))
	// sccStrings := make([][]string, len(sccs))
	// for i, bridge := range bridges {
	// 	bridgeStrings[i] = []string{string(bridge[0] + 'A'), string(bridge[1] + 'A')}
	// }
	// for i, scc := range sccs {
	// 	sccString := make([]string, len(scc))
	// 	for j, vertex := range scc {
	// 		sccString[j] = string(vertex + 'A')
	// 	}
	// 	sccStrings[i] = sccString
	// }

	// // Create the response object
	// response := SCCResponse{
	// 	Bridges: bridgeStrings,
	// 	SCCs:    sccStrings,
	// }

	// // Convert the response to JSON
	// responseJSON, err := json.Marshal(response)
	// if err != nil {
	// 	http.Error(w, "Failed to create response", http.StatusInternalServerError)
	// 	return
	// }

	// Set the Content-Type header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}


func main() {
	log.Println("Starting Backend")

	r := mux.NewRouter()
	// Define API routes
	r.HandleFunc("/api/username", usernameHandler).Methods("GET")
	r.HandleFunc("/api/scc", sccHandler).Methods("POST")

	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	distDir := filepath.Join(currentDir, "dist")
	spa := spaHandler{staticPath: distDir, indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	log.Println("Http Listening")
	http.ListenAndServe(
		port, r)
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)
	log.Println(path)
	indexFile := filepath.Join(h.staticPath, h.indexPath)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, indexFile)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
