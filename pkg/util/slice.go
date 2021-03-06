package util

import (
	"math/rand"

	"golang.org/x/exp/slices"
)

func DiffSliceFunc[T any](old, new []T, cmp func(a, b *T) int) (added, removed, kept []T) {
	added = []T{}
	removed = []T{}
	kept = []T{}

	less := func(a, b T) bool {
		return cmp(&a, &b) < 0
	}

	slices.SortFunc(new, less)
	slices.SortFunc(old, less)

	i, j := 0, 0
	for i < len(old) && j < len(new) {
		c := cmp(&old[i], &new[j])
		switch {
		case c < 0: // removed
			removed = append(removed, old[i])
			i++

		case c > 0: // added
			added = append(added, new[j])
			j++

		default: // kept
			kept = append(kept, new[j])
			i++
			j++
		}
	}

	// Add rest

	for ; i < len(old); i++ {
		removed = append(removed, old[i])
	}

	for ; j < len(new); j++ {
		added = append(added, new[j])
	}

	return
}

func IndexSlice[T any, I comparable](s []T, index func(*T) I) map[I]*T {
	m := map[I]*T{}

	for i := range s {
		e := &s[i]
		m[index(e)] = e
	}

	return m
}

func ShuffleSlice[T any](s []T) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

func FilterSlice[T any](s []T, cmp func(T) bool) []T {
	t := []T{}

	for _, i := range s {
		if cmp(i) {
			t = append(t, i)
		}
	}

	return t
}
