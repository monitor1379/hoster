package hoster

/*
 * @Date: 2020-11-29 14:07:02
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-11-29 14:49:49
 */

import (
	"os"
)

func isPathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func reduceEmptyItem(items []string) []string {
	newItems := make([]string, 0)
	for _, item := range items {
		if item != "" {
			newItems = append(newItems, item)
		}
	}
	return newItems
}

// diffset calculates set(left) - set(right).
func diffset(left []string, right []string) ([]string, []string) {
	rightSet := make(map[string]struct{})
	for _, r := range right {
		rightSet[r] = struct{}{}
	}

	newLeft := make([]string, 0)
	bothPresence := make([]string, 0)

	for _, l := range left {
		if _, ok := rightSet[l]; ok {
			bothPresence = append(bothPresence, l)
		} else {
			newLeft = append(newLeft, l)
		}
	}
	return newLeft, bothPresence
}
