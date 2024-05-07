package kmod

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"

	kmodpb "source.monogon.dev/osbase/kmod/spec"
)

// LookupModules looks up all matching modules for a given modalias device
// identifier.
func LookupModules(meta *kmodpb.Meta, modalias string) (mods []*kmodpb.Module) {
	matches := make(map[uint32]bool)
	lookupModulesRec(meta.ModuleDeviceMatches, modalias, matches)
	for idx := range matches {
		mods = append(mods, meta.Modules[idx])
	}
	sort.Slice(mods, func(i, j int) bool { return mods[i].Name < mods[j].Name })
	return
}

func lookupModulesRec(n *kmodpb.RadixNode, needle string, matches map[uint32]bool) {
	for _, c := range n.Children {
		switch c.Type {
		case kmodpb.RadixNode_LITERAL:
			if len(needle) < len(c.Literal) {
				continue
			}
			if c.Literal == needle[:len(c.Literal)] {
				lookupModulesRec(c, needle[len(c.Literal):], matches)
			}
		case kmodpb.RadixNode_WILDCARD:
			for i := 0; i <= len(needle); i++ {
				lookupModulesRec(c, needle[i:], matches)
			}
		case kmodpb.RadixNode_SINGLE_WILDCARD:
			if len(needle) < 1 {
				continue
			}
			lookupModulesRec(c, needle[1:], matches)
		case kmodpb.RadixNode_BYTE_RANGE:
			if len(needle) < 1 {
				continue
			}
			if needle[0] >= byte(c.StartByte) && needle[0] <= byte(c.EndByte) {
				lookupModulesRec(c, needle[1:], matches)
			}
		}
	}
	if len(needle) == 0 {
		for _, mi := range n.ModuleIndex {
			matches[mi] = true
		}
	}
}

// AddPattern adds a new pattern associated with a moduleIndex to the radix tree
// rooted at root.
func AddPattern(root *kmodpb.RadixNode, pattern string, moduleIndex uint32) error {
	pp, err := parsePattern(pattern)
	if err != nil {
		return fmt.Errorf("error parsing pattern %q: %w", pattern, err)
	}
	if len(pp) > 0 {
		pp[len(pp)-1].ModuleIndex = []uint32{moduleIndex}
	} else {
		// This exists to handle empty patterns, which have little use in
		// practice (but their behavior is well-defined). It exists primarily
		// to not crash in that case as well as to appease the Fuzzer.
		root.ModuleIndex = append(root.ModuleIndex, moduleIndex)
	}
	return addPatternRec(root, pp, nil)
}

// addPatternRec recursively adds a new pattern to the radix tree.
// If currPartOverride is non-nil it is used instead of the first part in the
// parts array.
func addPatternRec(n *kmodpb.RadixNode, parts []*kmodpb.RadixNode, currPartOverride *kmodpb.RadixNode) error {
	if len(parts) == 0 {
		return nil
	}
	var currPart *kmodpb.RadixNode
	if currPartOverride != nil {
		currPart = currPartOverride
	} else {
		currPart = parts[0]
	}
	for _, c := range n.Children {
		if c.Type != currPart.Type {
			continue
		}
		switch c.Type {
		case kmodpb.RadixNode_LITERAL:
			if c.Literal[0] == currPart.Literal[0] {
				var i int
				for i < len(c.Literal) && i < len(currPart.Literal) && c.Literal[i] == currPart.Literal[i] {
					i++
				}
				if i == len(c.Literal) && i == len(currPart.Literal) {
					if len(parts) == 1 {
						c.ModuleIndex = append(c.ModuleIndex, parts[0].ModuleIndex...)
						return nil
					}
					return addPatternRec(c, parts[1:], nil)
				}
				if i == len(c.Literal) {
					return addPatternRec(c, parts, &kmodpb.RadixNode{Type: kmodpb.RadixNode_LITERAL, Literal: currPart.Literal[i:], ModuleIndex: currPart.ModuleIndex})
				}
				// Split current node
				splitOldPart := &kmodpb.RadixNode{
					Type:        kmodpb.RadixNode_LITERAL,
					Literal:     c.Literal[i:],
					Children:    c.Children,
					ModuleIndex: c.ModuleIndex,
				}
				var splitNewPart *kmodpb.RadixNode
				// Current part is a strict subset of the node being traversed
				if i == len(currPart.Literal) {
					if len(parts) < 2 {
						c.Children = []*kmodpb.RadixNode{splitOldPart}
						c.Literal = currPart.Literal
						c.ModuleIndex = currPart.ModuleIndex
						return nil
					}
					splitNewPart = parts[1]
					parts = parts[1:]
				} else {
					splitNewPart = &kmodpb.RadixNode{
						Type:        kmodpb.RadixNode_LITERAL,
						Literal:     currPart.Literal[i:],
						ModuleIndex: currPart.ModuleIndex,
					}
				}
				c.Children = []*kmodpb.RadixNode{
					splitOldPart,
					splitNewPart,
				}
				c.Literal = currPart.Literal[:i]
				c.ModuleIndex = nil
				return addPatternRec(splitNewPart, parts[1:], nil)
			}

		case kmodpb.RadixNode_BYTE_RANGE:
			if c.StartByte == currPart.StartByte && c.EndByte == currPart.EndByte {
				if len(parts) == 1 {
					c.ModuleIndex = append(c.ModuleIndex, parts[0].ModuleIndex...)
				}
				return addPatternRec(c, parts[1:], nil)
			}
		case kmodpb.RadixNode_SINGLE_WILDCARD, kmodpb.RadixNode_WILDCARD:
			if len(parts) == 1 {
				c.ModuleIndex = append(c.ModuleIndex, parts[0].ModuleIndex...)
			}
			return addPatternRec(c, parts[1:], nil)
		}
	}
	// No child or common prefix found, append node
	n.Children = append(n.Children, currPart)
	return addPatternRec(currPart, parts[1:], nil)
}

// PrintTree prints the tree from the given root node to standard out.
// The output is not stable and should only be used for debugging/diagnostics.
// It will log and exit the process if it encounters invalid nodes.
func PrintTree(r *kmodpb.RadixNode) {
	printTree(r, 0, false)
}

func printTree(r *kmodpb.RadixNode, indent int, noIndent bool) {
	if !noIndent {
		for i := 0; i < indent; i++ {
			fmt.Print("  ")
		}
	}
	if len(r.ModuleIndex) > 0 {
		fmt.Printf("%v ", r.ModuleIndex)
	}
	switch r.Type {
	case kmodpb.RadixNode_LITERAL:
		fmt.Printf("%q: ", r.Literal)
	case kmodpb.RadixNode_SINGLE_WILDCARD:
		fmt.Printf("?: ")
	case kmodpb.RadixNode_WILDCARD:
		fmt.Printf("*: ")
	case kmodpb.RadixNode_BYTE_RANGE:
		fmt.Printf("[%c-%c]: ", rune(r.StartByte), rune(r.EndByte))
	default:
		log.Fatalf("Unknown tree type %T\n", r)
	}
	if len(r.Children) == 1 {
		printTree(r.Children[0], indent, true)
		return
	}
	fmt.Println("")
	for _, c := range r.Children {
		printTree(c, indent+1, false)
	}
}

// parsePattern parses a string pattern into a non-hierarchical list of
// RadixNodes. These nodes can then be futher modified and integrated into
// a Radix tree.
func parsePattern(pattern string) ([]*kmodpb.RadixNode, error) {
	var out []*kmodpb.RadixNode
	var i int
	var currentLiteral strings.Builder
	storeCurrentLiteral := func() {
		if currentLiteral.Len() > 0 {
			out = append(out, &kmodpb.RadixNode{
				Type:    kmodpb.RadixNode_LITERAL,
				Literal: currentLiteral.String(),
			})
			currentLiteral.Reset()
		}
	}
	for i < len(pattern) {
		switch pattern[i] {
		case '*':
			storeCurrentLiteral()
			i += 1
			if len(out) > 0 && out[len(out)-1].Type == kmodpb.RadixNode_WILDCARD {
				continue
			}
			out = append(out, &kmodpb.RadixNode{
				Type: kmodpb.RadixNode_WILDCARD,
			})
		case '?':
			storeCurrentLiteral()
			out = append(out, &kmodpb.RadixNode{
				Type: kmodpb.RadixNode_SINGLE_WILDCARD,
			})
			i += 1
		case '[':
			storeCurrentLiteral()
			if len(pattern) <= i+4 {
				return nil, errors.New("illegal byte range notation, not enough characters")
			}
			if pattern[i+2] != '-' || pattern[i+4] != ']' {
				return nil, errors.New("illegal byte range notation, incorrect dash or closing character")
			}
			nn := &kmodpb.RadixNode{
				Type:      kmodpb.RadixNode_BYTE_RANGE,
				StartByte: uint32(pattern[i+1]),
				EndByte:   uint32(pattern[i+3]),
			}
			if nn.StartByte > nn.EndByte {
				return nil, errors.New("byte range start byte larger than end byte")
			}
			out = append(out, nn)
			i += 5
		case '\\':
			if len(pattern) <= i+1 {
				return nil, errors.New("illegal escape character at the end of the string")
			}
			currentLiteral.WriteByte(pattern[i+1])
			i += 2
		default:
			currentLiteral.WriteByte(pattern[i])
			i += 1
		}
	}
	storeCurrentLiteral()
	return out, nil
}
