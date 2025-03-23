package main

import (
	"fmt"
	"sort"
	"strings"
)

// Given a list of folders folder, return the folders after removing all sub-folders in those folders.
//  You may return the answer in any order.

// If a folder[i] is located within another folder[j], it is called a sub-folder of it.
// A sub-folder of folder[j] must start with folder[j], followed by a "/". For example, "/a/b"
// is a sub-folder of "/a", but "/b" is not a sub-folder of "/a/b/c".

// The format of a path is one or more concatenated strings of the form:
// '/' followed by one or more lowercase English letters.

// For example, "/leetcode" and "/leetcode/problems" are valid paths while an empty string and "/" are not.

func removeSubfolders(folder []string) []string {
	if len(folder) < 2 {
		return folder
	}

	// sort folders, so that the smaller folders are at the front. This guarentees that these folders which could be parents are checked first
	sort.Slice(folder, func(a, b int) bool {
		return folder[a] < folder[b]
	})

	results := []string{}

	for _, f := range folder { // add / to mark dir
		if len(results) == 0 || !strings.HasPrefix(f, results[len(results)-1]+"/") { // the last item in results must be a parent
			results = append(results, f)
		}
	}

	return results
}

func main() {
	//folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"} //["/a","/c/d","/c/f"]
	folder := []string{"/a/b/c", "/a/b/ca", "/a/b/d"}
	fmt.Println(removeSubfolders(folder))
}
