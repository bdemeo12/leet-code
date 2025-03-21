package main

// You have information about n different recipes. You are given a string array recipes and a 2D string array ingredients. The ith recipe has the name recipes[i], and you can create it if you have all the needed ingredients from ingredients[i]. A recipe can also be an ingredient for other recipes, i.e., ingredients[i] may contain a string that is in recipes.

// You are also given a string array supplies containing all the ingredients that you initially have, and you have an infinite supply of all of them.

// Return a list of all the recipes that you can create. You may return the answer in any order.

// Note that two recipes may contain each other in their ingredients.

import "fmt"

func main() {
	// recipes := []string{"bread"}
	// ingredients := [][]string{{"yeast", "flour"}}
	// supplies := []string{"yeast"}

	recipes := []string{"xevvq", "izcad", "p", "we", "bxgnm", "vpio", "i", "hjvu", "igi", "anp", "tokfq", "z", "kwdmb", "g", "qb", "q", "b", "hthy"}
	ingredients := [][]string{
		{"wbjr"}, {"otr", "fzr", "g"}, {"fzr", "wi", "otr", "xgp", "wbjr", "igi", "b"},
		{"fzr", "xgp", "wi", "otr", "tokfq", "izcad", "igi", "xevvq", "i", "anp"},
		{"wi", "xgp", "wbjr"}, {"wbjr", "bxgnm", "i", "b", "hjvu", "izcad", "igi", "z", "g"},
		{"xgp", "otr", "wbjr"}, {"wbjr", "otr"}, {"wbjr", "otr", "fzr", "wi", "xgp", "hjvu", "tokfq", "z", "kwdmb"},
		{"xgp", "wi", "wbjr", "bxgnm", "izcad", "p", "xevvq"}, {"bxgnm"}, {"wi", "fzr", "otr", "wbjr"},
		{"wbjr", "wi", "fzr", "xgp", "otr", "g", "b", "p"}, {"otr", "fzr", "xgp", "wbjr"},
		{"xgp", "wbjr", "q", "vpio", "tokfq", "we"}, {"wbjr", "wi", "xgp", "we"},
		{"wbjr"}, {"wi"},
	}
	supplies := []string{"wi", "otr", "wbjr", "fzr", "xgp"}
	fmt.Println(findAllRecipes(recipes, ingredients, supplies))
}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	possibleRecipies := []string{}

	// put all supplies in a map so we can keep count
	suppliesCount := make(map[string]bool)
	processed := make(map[string]bool)
	for _, s := range supplies {
		suppliesCount[s] = true
	}

	changed := true
	for changed {
		changed = false

		for i, recipe := range recipes {

			if suppliesCount[recipe] || processed[recipe] {
				continue
			}

			canMake := false
			for j := 0; j < len(ingredients[i]); j++ {

				ingredient := ingredients[i][j]
				if _, ok := suppliesCount[ingredient]; !ok {
					canMake = false
					break
				} else {
					canMake = true
				}
			}

			if !canMake {
				continue
			}

			suppliesCount[recipe] = true
			processed[recipe] = true
			possibleRecipies = append(possibleRecipies, recipe)
			changed = true
		}

	}

	return possibleRecipies
}
