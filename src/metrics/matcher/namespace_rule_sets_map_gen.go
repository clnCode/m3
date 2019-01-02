// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package matcher

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// namespaceRuleSetsMapHash is the hash for a given map entry, this is public to support
// iterating over the map using a native Go for loop.
type namespaceRuleSetsMapHash uint64

// namespaceRuleSetsMapHashFn is the hash function to execute when hashing a key.
type namespaceRuleSetsMapHashFn func([]byte) namespaceRuleSetsMapHash

// namespaceRuleSetsMapEqualsFn is the equals key function to execute when detecting equality of a key.
type namespaceRuleSetsMapEqualsFn func([]byte, []byte) bool

// namespaceRuleSetsMapCopyFn is the copy key function to execute when copying the key.
type namespaceRuleSetsMapCopyFn func([]byte) []byte

// namespaceRuleSetsMapFinalizeFn is the finalize key function to execute when finished with a key.
type namespaceRuleSetsMapFinalizeFn func([]byte)

// namespaceRuleSetsMap uses the genny package to provide a generic hash map that can be specialized
// by running the following command from this root of the repository:
// ```
// make hashmap-gen pkg=outpkg key_type=Type value_type=Type out_dir=/tmp
// ```
// Or if you would like to use bytes or ident.ID as keys you can use the
// partially specialized maps to generate your own maps as well:
// ```
// make byteshashmap-gen pkg=outpkg value_type=Type out_dir=/tmp
// make idhashmap-gen pkg=outpkg value_type=Type out_dir=/tmp
// ```
// This will output to stdout the generated source file to use for your map.
// It uses linear probing by incrementing the number of the hash created when
// hashing the identifier if there is a collision.
// namespaceRuleSetsMap is a value type and not an interface to allow for less painful
// upgrades when adding/removing methods, it is not likely to need mocking so
// an interface would not be super useful either.
type namespaceRuleSetsMap struct {
	_namespaceRuleSetsMapOptions

	// lookup uses hash of the identifier for the key and the MapEntry value
	// wraps the value type and the key (used to ensure lookup is correct
	// when dealing with collisions), we use uint64 for the hash partially
	// because lookups of maps with uint64 keys has a fast path for Go.
	lookup map[namespaceRuleSetsMapHash]namespaceRuleSetsMapEntry
}

// _namespaceRuleSetsMapOptions is a set of options used when creating an identifier map, it is kept
// private so that implementers of the generated map can specify their own options
// that partially fulfill these options.
type _namespaceRuleSetsMapOptions struct {
	// hash is the hash function to execute when hashing a key.
	hash namespaceRuleSetsMapHashFn
	// equals is the equals key function to execute when detecting equality.
	equals namespaceRuleSetsMapEqualsFn
	// copy is the copy key function to execute when copying the key.
	copy namespaceRuleSetsMapCopyFn
	// finalize is the finalize key function to execute when finished with a
	// key, this is optional to specify.
	finalize namespaceRuleSetsMapFinalizeFn
	// initialSize is the initial size for the map, use zero to use Go's std map
	// initial size and consequently is optional to specify.
	initialSize int
}

// namespaceRuleSetsMapEntry is an entry in the map, this is public to support iterating
// over the map using a native Go for loop.
type namespaceRuleSetsMapEntry struct {
	// key is used to check equality on lookups to resolve collisions
	key _namespaceRuleSetsMapKey
	// value type stored
	value RuleSet
}

type _namespaceRuleSetsMapKey struct {
	key      []byte
	finalize bool
}

// Key returns the map entry key.
func (e namespaceRuleSetsMapEntry) Key() []byte {
	return e.key.key
}

// Value returns the map entry value.
func (e namespaceRuleSetsMapEntry) Value() RuleSet {
	return e.value
}

// _namespaceRuleSetsMapAlloc is a non-exported function so that when generating the source code
// for the map you can supply a public constructor that sets the correct
// hash, equals, copy, finalize options without users of the map needing to
// implement them themselves.
func _namespaceRuleSetsMapAlloc(opts _namespaceRuleSetsMapOptions) *namespaceRuleSetsMap {
	m := &namespaceRuleSetsMap{_namespaceRuleSetsMapOptions: opts}
	m.Reallocate()
	return m
}

func (m *namespaceRuleSetsMap) newMapKey(k []byte, opts _namespaceRuleSetsMapKeyOptions) _namespaceRuleSetsMapKey {
	key := _namespaceRuleSetsMapKey{key: k, finalize: opts.finalizeKey}
	if !opts.copyKey {
		return key
	}

	key.key = m.copy(k)
	return key
}

func (m *namespaceRuleSetsMap) removeMapKey(hash namespaceRuleSetsMapHash, key _namespaceRuleSetsMapKey) {
	delete(m.lookup, hash)
	if key.finalize {
		m.finalize(key.key)
	}
}

// Get returns a value in the map for an identifier if found.
func (m *namespaceRuleSetsMap) Get(k []byte) (RuleSet, bool) {
	hash := m.hash(k)
	for entry, ok := m.lookup[hash]; ok; entry, ok = m.lookup[hash] {
		if m.equals(entry.key.key, k) {
			return entry.value, true
		}
		// Linear probe to "next" to this entry (really a rehash)
		hash++
	}
	var empty RuleSet
	return empty, false
}

// Set will set the value for an identifier.
func (m *namespaceRuleSetsMap) Set(k []byte, v RuleSet) {
	m.set(k, v, _namespaceRuleSetsMapKeyOptions{
		copyKey:     true,
		finalizeKey: m.finalize != nil,
	})
}

// namespaceRuleSetsMapSetUnsafeOptions is a set of options to use when setting a value with
// the SetUnsafe method.
type namespaceRuleSetsMapSetUnsafeOptions struct {
	NoCopyKey     bool
	NoFinalizeKey bool
}

// SetUnsafe will set the value for an identifier with unsafe options for how
// the map treats the key.
func (m *namespaceRuleSetsMap) SetUnsafe(k []byte, v RuleSet, opts namespaceRuleSetsMapSetUnsafeOptions) {
	m.set(k, v, _namespaceRuleSetsMapKeyOptions{
		copyKey:     !opts.NoCopyKey,
		finalizeKey: !opts.NoFinalizeKey,
	})
}

type _namespaceRuleSetsMapKeyOptions struct {
	copyKey     bool
	finalizeKey bool
}

func (m *namespaceRuleSetsMap) set(k []byte, v RuleSet, opts _namespaceRuleSetsMapKeyOptions) {
	hash := m.hash(k)
	for entry, ok := m.lookup[hash]; ok; entry, ok = m.lookup[hash] {
		if m.equals(entry.key.key, k) {
			m.lookup[hash] = namespaceRuleSetsMapEntry{
				key:   entry.key,
				value: v,
			}
			return
		}
		// Linear probe to "next" to this entry (really a rehash)
		hash++
	}

	m.lookup[hash] = namespaceRuleSetsMapEntry{
		key:   m.newMapKey(k, opts),
		value: v,
	}
}

// Iter provides the underlying map to allow for using a native Go for loop
// to iterate the map, however callers should only ever read and not write
// the map.
func (m *namespaceRuleSetsMap) Iter() map[namespaceRuleSetsMapHash]namespaceRuleSetsMapEntry {
	return m.lookup
}

// Len returns the number of map entries in the map.
func (m *namespaceRuleSetsMap) Len() int {
	return len(m.lookup)
}

// Contains returns true if value exists for key, false otherwise, it is
// shorthand for a call to Get that doesn't return the value.
func (m *namespaceRuleSetsMap) Contains(k []byte) bool {
	_, ok := m.Get(k)
	return ok
}

// Delete will remove a value set in the map for the specified key.
func (m *namespaceRuleSetsMap) Delete(k []byte) {
	hash := m.hash(k)
	for entry, ok := m.lookup[hash]; ok; entry, ok = m.lookup[hash] {
		if m.equals(entry.key.key, k) {
			m.removeMapKey(hash, entry.key)
			return
		}
		// Linear probe to "next" to this entry (really a rehash)
		hash++
	}
}

// Reset will reset the map by simply deleting all keys to avoid
// allocating a new map.
func (m *namespaceRuleSetsMap) Reset() {
	for hash, entry := range m.lookup {
		m.removeMapKey(hash, entry.key)
	}
}

// Reallocate will avoid deleting all keys and reallocate a new
// map, this is useful if you believe you have a large map and
// will not need to grow back to a similar size.
func (m *namespaceRuleSetsMap) Reallocate() {
	if m.initialSize > 0 {
		m.lookup = make(map[namespaceRuleSetsMapHash]namespaceRuleSetsMapEntry, m.initialSize)
	} else {
		m.lookup = make(map[namespaceRuleSetsMapHash]namespaceRuleSetsMapEntry)
	}
}
