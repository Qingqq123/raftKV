package set

import (
	"github.com/hdt3213/godis/datastruct/dict"
)

type Set struct {
	dict dict.Dict
}

func Make(members ...string) *Set {
	set := &Set{
		dict: dict.MakeSimple(),
	}
	for _, member := range members {
		set.Add(member)
	}
	return set
}

func (set *Set) Add(val string) int {
	return set.dict.Put(val, nil)
}

func (set *Set) Remove(val string) int {
	return set.dict.Remove(val)
}

func (set *Set) Has(val string) bool {
	if set == nil || set.dict == nil {
		return false
	}
	_, exists := set.dict.Get(val)
	return exists
}

func (set *Set) Len() int {
	if set == nil || set.dict == nil {
		return 0
	}
	return set.dict.Len()
}

func (set *Set) ToSlice() []string {
	slice := make([]string, set.Len())
	i := 0
	set.dict.ForEach(func(key string, val interface{}) bool {
		if i < len(slice) {
			slice[i] = key
		} else {
			slice = append(slice, key)
		}
		i++
		return true
	})
	return slice
}

func (set *Set) ForEach(consumer func(member string) bool) {
	if set == nil || set.dict == nil {
		return
	}
	set.dict.ForEach(func(key string, val interface{}) bool {
		return consumer(key)
	})
}

func (set *Set) ShallowCopy() *Set {
	result := Make()
	set.ForEach(func(member string) bool {
		result.Add(member)
		return true
	})
	return result
}

func Intersect(sets ...*Set) *Set {
	result := Make()
	if len(sets) == 0 {
		return result
	}

	countMap := make(map[string]int)
	for _, set := range sets {
		set.ForEach(func(member string) bool {
			countMap[member]++
			return true
		})
	}
	for k, v := range countMap {
		if v == len(sets) {
			result.Add(k)
		}
	}
	return result
}

func Union(sets ...*Set) *Set {
	result := Make()
	for _, set := range sets {
		set.ForEach(func(member string) bool {
			result.Add(member)
			return true
		})
	}
	return result
}

func Diff(sets ...*Set) *Set {
	if len(sets) == 0 {
		return Make()
	}
	result := sets[0].ShallowCopy()
	for i := 1; i < len(sets); i++ {
		sets[i].ForEach(func(member string) bool {
			result.Remove(member)
			return true
		})
		if result.Len() == 0 {
			break
		}
	}
	return result
}

func (set *Set) RandomMembers(limit int) []string {
	if set == nil || set.dict == nil {
		return nil
	}
	return set.dict.RandomKeys(limit)
}

func (set *Set) RandomDistinctMembers(limit int) []string {
	return set.dict.RandomDistinctKeys(limit)
}
