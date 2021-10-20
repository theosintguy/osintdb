package data

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "osintdb/internal/model"
)

func Import(filename string) ([]*model.Resource, error) {
    var data []*model.Resource
    if _, err := os.Stat(filename); err != nil {
        return nil, err
    }
    jsonFile, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer jsonFile.Close()
    b, _ := ioutil.ReadAll(jsonFile)
    err = json.Unmarshal(b, &data)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func Export(data []*model.Resource, filename string) error {
    // note: does not check if file exists before possibly overwriting as this should be done before
    jsonString, err := json.Marshal(&data)
    if err != nil {
        return err
    }
    err = ioutil.WriteFile(filename, jsonString, 0644)
    if err != nil {
        return err
    }
    return nil
}
