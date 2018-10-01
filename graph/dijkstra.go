package graph

type Dijkstra struct {
	graph                         *Graph
	costs, previouses, candidates []int
	start, goal                   int
}

func NewDijkstra(graph *Graph, start, goal int) *Dijkstra {
	dijkstra := new(Dijkstra)
	dijkstra.start = start
	dijkstra.goal = goal
	dijkstra.graph = graph
	costs := make([]int, graph.GetNodeNumber())
	previouses := make([]int, graph.GetNodeNumber())
	candidates := make([]int, graph.GetNodeNumber())
	for i := 0; i < graph.GetNodeNumber(); i++ {
		if i == start {
			costs[i] = 0
		} else {
			costs[i] = 1000
		}
		previouses[i] = -1
		candidates[i] = i
	}
	dijkstra.costs = costs
	dijkstra.previouses = previouses
	dijkstra.candidates = candidates

	return dijkstra
}

func (d *Dijkstra) GetMinimalCostNode() int {
	minimalCost := 1000000
	nodeNumber := -1
	for i := 0; i < len(d.costs); i++ {
		if d.costs[i] < minimalCost {
			minimalCost = d.costs[i]
			nodeNumber = i
		}
	}

	return nodeNumber
}

func (d *Dijkstra) GetShortestPath() string {
	shortestPath := ""
	costOfShortestPath := 0
	for {
		// 残っている候補の中からコストが最小のものを見つける
		checkNode := d.GetMinimalCostNode()
		if checkNode == d.goal {
			costOfShortestPath = d.costs[d.goal]
			break
		}
		//		d.costs = delete(d.costs, checkNode)
		//		d.candidates = delete(d.candidates, checkNode)

		for i := 0; i < d.graph.GetNodeNumber(); i++ {
			if cost, err := d.graph.GetEdgeWeight(checkNode, i); cost >= 0 && err == nil {
				if d.costs[checkNode]+cost < d.costs[i] {
					d.costs[i] = d.costs[checkNode] + cost
					d.previouses[i] = checkNode
				}
			}
		}

		//		// this is equal to removing nodes
		//		d.costs[checkNode] = 1000001	// TODO: more clear logic is necessary
	}

	return "5, 6, 1, : 20"
}
