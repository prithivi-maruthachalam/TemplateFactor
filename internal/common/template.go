package common

import (
	"fmt"
	"os"
	"strings"

	tf_io "github.com/prithivi-maruthachalam/TemplateFactory/templatefactory/internal/io"
)

// Represents a single node in a template. This can be a directory or a file with or without content
type TemplateNode struct {
	NodePath  string
	IsFile    bool
	IsContent bool
	Content   string
}

func (node *TemplateNode) String() string {
	nodeTypeStr := "dir"
	if node.IsFile {
		nodeTypeStr = "file"
	}
	return fmt.Sprintf("  - %s %s content:%d\n", node.NodePath, nodeTypeStr, len(node.Content))
}

// Template factory template object
type Template struct {
	TemplateName string
	Nodes        []TemplateNode
}

// Add a file/directory node to the template
func (template *Template) AddNode(node TemplateNode) {
	template.Nodes = append(template.Nodes, node)
}

func (template *Template) String() string {
	str := fmt.Sprintf("Name: %s\n", template.TemplateName)

	for _, node := range template.Nodes {
		str += fmt.Sprintf("%v", &node)
	}

	return str
}

func (template *Template) Describe(printContent bool) string {
	str := tf_io.Title(fmt.Sprintf("Template Name : %s\n", template.TemplateName))

	// returns the depth from root for a given path
	getDepthFromPath := func(path string) int {
		parts := strings.Split(path, string(os.PathSeparator))
		return len(parts)
	}

	// Returns the colored name for a node
	getFormattedName := func(node *TemplateNode) string {
		parts := strings.Split(node.NodePath, string(os.PathSeparator))
		name := parts[len(parts)-1]
		if node.IsFile {
			return tf_io.FileName(name)
		} else {
			return tf_io.DirName(name)
		}
	}

	// Get the depth of the next node
	getNextNodeDepth := func(i int) int {
		if i < len(template.Nodes)-1 {
			return getDepthFromPath(template.Nodes[i+1].NodePath)
		} else {
			return -1
		}
	}

	for i, currentNode := range template.Nodes {
		// Depth of the current node
		currentNodeDepth := getDepthFromPath(currentNode.NodePath)

		// Depth of the next node
		nextNodeDepth := getNextNodeDepth(i)

		if currentNodeDepth > 1 {
			str += strings.Repeat("│   ", currentNodeDepth-1)
		}

		if currentNodeDepth > nextNodeDepth {
			str += "└── "
		} else {
			str += "├── "
		}

		str += getFormattedName(&currentNode)
		if currentNode.IsFile && currentNode.IsContent {
			str += tf_io.SubtleText(fmt.Sprintf(" %dB", len(currentNode.Content)))
		}

		str += "\n"
	}

	if printContent {
		str += "\n"
		for _, node := range template.Nodes {
			if node.IsFile && node.IsContent {
				str += tf_io.FileName(fmt.Sprintf("file: %s\n", node.NodePath))
				str += tf_io.FileContent(node.Content)
				str += "\n"
			}
		}
	}

	return str
}
