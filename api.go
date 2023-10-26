package fdatabase

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var FolderName string = "data"

// The file in which it saves the register depends on FolderName and the StructFolder (output of the function GetFolderName
// for it's type) : FolderName / StructFolder / ID.json
func Save(r register) error {

	if r.GetId() == 0 {
		return errors.New("ID can't be 0")
	}

	path := fmt.Sprintf("%s/%s/%d.json", FolderName, r.GetFolderName(), r.GetId())

	err := handleMissingDir(fmt.Sprintf("%s", FolderName))
	if err != nil {
		return err
	}

	err = handleMissingDir(fmt.Sprintf("%s/%s", FolderName, r.GetFolderName()))
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)

	encoder.Encode(r)

	return nil
}

// The type between [] determines the type of the output, loads from FolderName / StructFolder / ID.json
// More about that on the Save() doc
func Load[T register](ID uint) (T, error) {
	var dummy T // Only purpose is call GetFolderName
	path := fmt.Sprintf("%s/%s/%d.json", FolderName, dummy.GetFolderName(), ID)

	file, err := os.Open(path)
	if err != nil {
		return dummy, err
	}

	decoder := json.NewDecoder(file)

	var data T
	err = decoder.Decode(&data)
	if err != nil {
		return dummy, err
	}

	return data, nil
}
