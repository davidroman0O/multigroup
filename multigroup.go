package multigroup

import "errors"

var (
	ErrKeyValueNotFound = errors.New("key value not found")
)

// `KeyValue` represents a key-value pair.
type KeyValue[K comparable] struct {
	Key   string
	Value K
}

// `Group` represents a single group of items with multiple key levels.
type Group[K comparable, V any] struct {
	Keys  []KeyValue[K]
	Items []V
}

func (g *Group[K, V]) GetKeyValue(name string) (*KeyValue[K], error) {
	for i := 0; i < len(g.Keys); i++ {
		if g.Keys[i].Key == name {
			return &g.Keys[i], nil
		}
	}
	return nil, ErrKeyValueNotFound
}

// `By` groups the slice data by the provided iteratees.
func By[T any, K comparable](data []T, iteratees ...func(T) (string, K)) []Group[K, T] {
	if len(iteratees) == 0 {
		panic("no iteratees provided for grouping")
	}

	var recursiveGroupBy func([]T, []func(T) (string, K)) []Group[K, T]
	recursiveGroupBy = func(data []T, itrs []func(T) (string, K)) []Group[K, T] {
		if len(itrs) == 0 {
			return []Group[K, T]{{Items: data}} // base case, no more grouping
		}

		groups := make(map[K][]T)
		groupKeys := make(map[K]string)
		for _, item := range data {
			groupKey, keyValue := itrs[0](item)
			groups[keyValue] = append(groups[keyValue], item)
			groupKeys[keyValue] = groupKey
		}

		var result []Group[K, T]
		for keyValue, items := range groups {
			subGroups := recursiveGroupBy(items, itrs[1:]) // recurse with remaining iteratees
			if len(subGroups) == 1 && len(subGroups[0].Keys) == 0 {
				// If no more subgroups, just append the current group
				result = append(result, Group[K, T]{Keys: []KeyValue[K]{{Key: groupKeys[keyValue], Value: keyValue}}, Items: items})
			} else {
				for _, sg := range subGroups {
					newKeys := append([]KeyValue[K]{{Key: groupKeys[keyValue], Value: keyValue}}, sg.Keys...)
					result = append(result, Group[K, T]{Keys: newKeys, Items: sg.Items})
				}
			}
		}
		return result
	}

	return recursiveGroupBy(data, iteratees)
}
