package day8

import (
	"2023/ergo"
	"2023/mymath"
	"fmt"
	"strings"
)

type node struct {
	name  string
	left  string
	right string
}

func Solve1() (steps int) {
	scanner := ergo.GetInputScanner("solutions/day8/input.txt")

	scanner.Scan()
	instructionStr := scanner.Text()

	nodes := map[string]node{}

	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		nodeParts := strings.Split(line, " = ")
		nodeName, nodeDirsStr := nodeParts[0], nodeParts[1]
		nodeLeft, nodeRight, _ := strings.Cut(nodeDirsStr, ", ")
		nodeLeft, nodeRight = nodeLeft[1:], nodeRight[:len(nodeRight)-1]
		nodes[nodeName] = node{nodeName, nodeLeft, nodeRight}
	}

	currNode := nodes["AAA"]
	instructionStrLen := len(instructionStr)
	for i := 0; currNode.name != "ZZZ"; i++ {
		currInstruction := instructionStr[i%instructionStrLen]
		if currInstruction == 'L' {
			currNode = nodes[currNode.left]
		} else {
			currNode = nodes[currNode.right]
		}
		steps += 1
	}

	return steps
}

func (n node) advance(dir uint8, nodes map[string]node) (node, bool) {
	var newNode node
	if dir == 'L' {
		newNode = nodes[n.left]
	} else {
		newNode = nodes[n.right]
	}
	return newNode, newNode.name[len(newNode.name)-1] == 'Z'
}

func advanceAllNodes(dir uint8, nodeMap map[string]node, nodes []node) (allDone bool) {
	allDone = true
	for i := 0; i < len(nodes); i++ {
		isDone := false
		nodes[i], isDone = nodes[i].advance(dir, nodeMap)
		if !isDone {
			allDone = false
		}
	}

	return allDone
}

func Solve2() int {
	scanner := ergo.GetInputScanner("solutions/day8/input.txt")

	scanner.Scan()
	instructionStr := scanner.Text()
	nodes := map[string]node{}
	scanner.Scan()

	var startNodes []node

	for scanner.Scan() {
		line := scanner.Text()

		nodeParts := strings.Split(line, " = ")
		nodeName, nodeDirsStr := nodeParts[0], nodeParts[1]
		nodeLeft, nodeRight, _ := strings.Cut(nodeDirsStr, ", ")
		nodeLeft, nodeRight = nodeLeft[1:], nodeRight[:len(nodeRight)-1]

		newNode := node{nodeName, nodeLeft, nodeRight}
		nodes[nodeName] = newNode
		if nodeName[len(nodeName)-1] == 'A' {
			startNodes = append(startNodes, newNode)
		}
	}

	instructionStrLen := len(instructionStr)

	// some directed ai modifications
	ans := 1
	for _, startNode := range startNodes {
		steps := 0
		isDone := false
		currNode := startNode

		for !isDone {
			dir := instructionStr[steps%instructionStrLen]
			currNode, isDone = currNode.advance(dir, nodes)
			steps++
		}

		ans = mymath.Lcm(ans, steps)
		fmt.Println(startNode.name, "reached end in", steps, "steps")
	}

	return ans
}
