package maps

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNoSuchWord            = DictionaryErr("no such word")
	ErrWordAlreadyExists     = DictionaryErr("word already exists")
	ErrUpdateNonExistingWord = DictionaryErr("cannot update a non-existing word")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if definition, ok := d[word]; ok {
		return definition, nil
	}
	return "", ErrNoSuchWord
}

func (d Dictionary) Add(word, definition string) error {
	switch _, err := d.Search(word); err {
	case ErrNoSuchWord:
		d[word] = definition
	case nil:
		return ErrWordAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	switch _, err := d.Search(word); err {
	case ErrNoSuchWord:
		return ErrUpdateNonExistingWord
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
