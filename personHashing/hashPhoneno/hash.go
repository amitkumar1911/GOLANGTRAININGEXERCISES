package hashPhoneno

import "hash/fnv"

func Hashing(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
