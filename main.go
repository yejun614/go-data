package main

import (
	"os"

	"github.com/vmihailenco/msgpack/v5"
)

type Data[T any] struct {
	Filepath string `json:"filepath" xml:"filepath" form:"filepath"`
	Data     T      `json:"data" xml:"data" form:"data"`
}

func New[T any](filepath string, sample T) *Data[T] {
	// create data
	data := new(Data[T])
	// set
	data.Filepath = filepath
	// open data
	err := data.Open()
	if err != nil {
		data.Data = sample
		data.Save() // create new file
	}
	// return data
	return data
}

func (data *Data[T]) Open() error {
	f, err := os.ReadFile(data.Filepath)
	if err != nil {
		return err
	}

	err = msgpack.Unmarshal(f, &data.Data)
	if err != nil {
		return err
	}

	return nil
}

func (data *Data[T]) Save() error {
	f, err := msgpack.Marshal(data.Data)
	if err != nil {
		return err
	}

	err = os.WriteFile(data.Filepath, f, 0644)
	if err != nil {
		return err
	}

	return nil
}
