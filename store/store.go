package store

type KeyValStore struct {
	userData map[string]string
}

func (kvs KeyValStore) UserNameForID(userID string) (string, bool) {
	name, ok := kvs.userData[userID]
	return name, ok
}

func CreateKeyValStore() KeyValStore {
	return KeyValStore{
		userData: map[string]string{
			"1": "Alice",
			"2": "Bob",
			"3": "Charlie",
		},
	}
}
