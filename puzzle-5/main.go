package main

// https://adventofcode.com/2024/day/5

import (
	"aoc/helper"
	"fmt"
	"regexp"
	"sort"
)

func main() {
	lines := helper.ReadLines("input.txt")
	rules, updates := parseRulesAndUpdates(lines)

	correctlyOrderedUpdates, incorrectlyOrderedUpdates := separateUpdatesByOrdering(rules, updates)

	solution1 := sumMidPageNumbers(correctlyOrderedUpdates)
	fmt.Println("-> part 1:", solution1)

	fixedUpdates := fixOrderingOfUpdates(rules, incorrectlyOrderedUpdates)
	solution2 := sumMidPageNumbers(fixedUpdates)
	fmt.Println("-> part 2:", solution2)
}

type Rule struct {
	A, B int
}

type Update []int

func parseRulesAndUpdates(lines []string) ([]Rule, []Update) {
	patternRule := regexp.MustCompile(`^(\d+)\|(\d+)$`)
	patternUpdate := regexp.MustCompile(`^\d[\d,]+$`)
	rules := make([]Rule, 0)
	updates := make([]Update, 0)
	for _, line := range lines {
		if m := patternRule.FindStringSubmatch(line); len(m) == 3 {
			rules = append(rules, Rule{
				A: helper.ParseInt[int](m[1]),
				B: helper.ParseInt[int](m[2]),
			})
		}
		if m := patternUpdate.FindStringSubmatch(line); len(m) == 1 {
			updates = append(updates, Update(helper.ParseInts(m[0])))
		}
	}
	return rules, updates
}

func sumMidPageNumbers(updates []Update) int {
	var sum int
	for _, u := range updates {
		sum += u[len(u)/2]
	}
	return sum
}

func separateUpdatesByOrdering(rules []Rule, updates []Update) ([]Update, []Update) {
	ruleMap := make(map[int]map[int]Rule)
	for _, r := range rules {
		if _, ok := ruleMap[r.A]; !ok {
			ruleMap[r.A] = make(map[int]Rule, 0)
		}
		ruleMap[r.A][r.B] = r
	}

	correctlyOrderedUpdates := make([]Update, 0)
	incorrectlyOrderedUpdates := make([]Update, 0)
	for _, u := range updates {
		if isWellOrdered(ruleMap, u) {
			correctlyOrderedUpdates = append(correctlyOrderedUpdates, u)
		} else {
			incorrectlyOrderedUpdates = append(incorrectlyOrderedUpdates, u)
		}
	}
	return correctlyOrderedUpdates, incorrectlyOrderedUpdates
}

func isWellOrdered(ruleMap map[int]map[int]Rule, update Update) bool {
	pageIndices := make(map[int]int)
	for i := range update {
		pageIndices[update[i]] = i
	}

	matchingRules := make([]Rule, 0)
	for _, p := range update {
		if rc, ok := ruleMap[p]; ok {
			for _, r := range rc {
				// first value is contained
				if _, ok := pageIndices[r.B]; ok {
					matchingRules = append(matchingRules, r)
				}
			}
		}
	}

	for _, r := range matchingRules {
		if pageIndices[r.A] > pageIndices[r.B] {
			return false
		}
	}
	return true
}

func fixOrderingOfUpdates(rules []Rule, incorrectlyOrderedUpdates []Update) []Update {
	ruleMap := make(map[int]map[int]Rule)
	for _, r := range rules {
		if _, ok := ruleMap[r.A]; !ok {
			ruleMap[r.A] = make(map[int]Rule, 0)
		}
		ruleMap[r.A][r.B] = r
	}

	fixedUpdates := make([]Update, 0, len(incorrectlyOrderedUpdates))
	for _, u := range incorrectlyOrderedUpdates {
		fixedUpdates = append(fixedUpdates, fixOrdering(ruleMap, u))
	}
	return fixedUpdates
}

func fixOrdering(ruleMap map[int]map[int]Rule, update Update) Update {
	pageIndices := make(map[int]int)
	for i := range update {
		pageIndices[update[i]] = i
	}

	matchingRules := make([]Rule, 0)
	for _, p := range update {
		if rc, ok := ruleMap[p]; ok {
			for _, r := range rc {
				// first value is contained
				if _, ok := pageIndices[r.B]; ok {
					matchingRules = append(matchingRules, r)
				}
			}
		}
	}

	fixed := make(Update, len(update))
	copy(fixed, update)
	sort.Slice(fixed, func(i, j int) bool {
		for _, r := range matchingRules {
			if (fixed[i] == r.A || fixed[i] == r.B) && (fixed[j] == r.A || fixed[j] == r.B) {
				return fixed[i] == r.A
			}
		}
		return false
	})
	return fixed
}
