package utils

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func UniqueStrings(strSlice []string) []string {
    keys := make(map[string]bool)
    list := make([]string, 0)
    for _, entry := range strSlice {
        if _, found := keys[entry]; !found {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

func FindConfigDir() (string, error) {
    osintdb_dir := os.Getenv("OSINTDB_DIR")
    if osintdb_dir == "" {
        home_dir := os.Getenv("HOME")
        if _, err := os.Stat(home_dir); err != nil {
            return "", err
        }
        osintdb_dir = filepath.Join(home_dir, ".local", "osintdb")
    }
    return osintdb_dir, nil
}

func FindInitData() (string, error) {
    dir, err := FindConfigDir()
    if err != nil {
        return "", err
    }
    initDataPath := filepath.Join(dir, "tools.min.json")
    if _, err := os.Stat(initDataPath); err != nil {
        return "", err
    }
    return initDataPath, nil
}

func CopyFile(src, dst string) error {
    // modified version of Method 1 in this article: https://opensource.com/article/18/6/copying-files-go
    sourceFileStat, err := os.Stat(src)
    if err != nil {
        return err
    }

    if !sourceFileStat.Mode().IsRegular() {
        return fmt.Errorf("%s is not a regular file", src)
    }

    source, err := os.Open(src)
    if err != nil {
        return err
    }
    defer source.Close()

    destination, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destination.Close()
    _, err = io.Copy(destination, source)
    return err
}