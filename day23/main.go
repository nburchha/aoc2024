package main

import (
	"fmt"
	"os"
	"strings"
	"sort"
)

func main() {
	file, _ := os.ReadFile("input/day23.txt")
	connections := strings.Split(string(file), "\n")
	graph := buildGraph(connections)
	validTriads := findTriads(graph)
	fmt.Println("Number of valid triads:", len(validTriads))

	maxClique := bronKerboschPivot(graph)
	fmt.Println("Password:", generatePassword(maxClique))
}

func buildGraph(connections []string) map[string]map[string]bool {
	graph := make(map[string]map[string]bool)

	for _, conn := range connections {
		parts := strings.Split(conn, "-")
		if len(parts) != 2 {
			continue
		}
		a, b := parts[0], parts[1]

		if graph[a] == nil {
			graph[a] = make(map[string]bool)
		}
		if graph[b] == nil {
			graph[b] = make(map[string]bool)
		}

		graph[a][b] = true
		graph[b][a] = true
	}

	return graph
}

func findTriads(graph map[string]map[string]bool) [][]string {
	var triads [][]string
	for a := range graph {
		for b := range graph[a] {
			if a >= b {
				continue
			}
			for c := range graph[b] {
				if b >= c || !graph[a][c] {
					continue
				}
				triad := []string{a, b, c}
				if isValidTriad(triad) {
					triads = append(triads, triad)
				}
			}
		}
	}

	return triads
}

func isValidTriad(triad []string) bool {
	for _, node := range triad {
		if strings.HasPrefix(node, "t") {
			return true
		}
	}
	return false
}

func bronKerboschPivot(graph map[string]map[string]bool) []string {
	var maxClique []string

	allNodes := make([]string, 0, len(graph))
	for node := range graph {
		allNodes = append(allNodes, node)
	}

	var r, p, x []string
	p = append(p, allNodes...) // Start with all nodes in P
	bronKerbosch(graph, r, p, x, &maxClique)
	return maxClique
}

func bronKerbosch(graph map[string]map[string]bool, r, p, x []string, maxClique *[]string) {
	if len(p) == 0 && len(x) == 0 {
		// Found a maximal clique
		if len(r) > len(*maxClique) {
			*maxClique = append([]string{}, r...)
		}
		return
	}
	pivot := ""
	if len(p) > 0 {
		pivot = p[0]
	}

	// Consider each node not in the neighborhood of the pivot
	for _, v := range difference(p, neighbors(graph, pivot)) {
		bronKerbosch(graph, append(r, v), intersect(p, neighbors(graph, v)), intersect(x, neighbors(graph, v)), maxClique)
		p = remove(p, v)
		x = append(x, v)
	}
}

func neighbors(graph map[string]map[string]bool, node string) []string {
	var result []string
	for neighbor := range graph[node] {
		result = append(result, neighbor)
	}
	return result
}

func intersect(a, b []string) []string {
	set := make(map[string]bool)
	for _, v := range b {
		set[v] = true
	}

	var result []string
	for _, v := range a {
		if set[v] {
			result = append(result, v)
		}
	}
	return result
}

func difference(a, b []string) []string {
	set := make(map[string]bool)
	for _, v := range b {
		set[v] = true
	}

	var result []string
	for _, v := range a {
		if !set[v] {
			result = append(result, v)
		}
	}
	return result
}

func remove(slice []string, item string) []string {
	var result []string
	for _, v := range slice {
		if v != item {
			result = append(result, v)
		}
	}
	return result
}

func generatePassword(clique []string) string {
	sort.Strings(clique)
	return strings.Join(clique, ",")
}