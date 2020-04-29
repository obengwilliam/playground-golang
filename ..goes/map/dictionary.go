package main

const ErrWordDoesNotExist = DictionaryErr("word does not exist")
const ErrWordExist = DictionaryErr("word does exist")

type DictionaryErr string

func (err DictionaryErr) Error() string {
	return string(err)
}

// Create maps
// Search for items in maps
// Add new items to maps
// Update items in maps
// Delete items from a map
// Learned more about errors
// How to create errors that are constants
// Writing error wrappers

type Dictionary map[string]string

func (d Dictionary) Update(w string, def string) (string, error) {
	if _, err := d.Search(w); err != nil {
		return def, err
	}
	d[w] = def
	return def, nil
}

// Search a Dictionary
func (d Dictionary) Search(w string) (string, error) {
	definiationForW, ok := d[w]
	if ok {
		return definiationForW, nil
	}
	return d[w], ErrWordDoesNotExist
}

func (d Dictionary) Add(w string, def string) (string, error) {
	_, ok := d[w]

	if ok {
		return w, ErrWordExist
	}
	d[w] = def
	return w, nil
}
