package model

import (
	"bytes"
	"encoding/gob"
	"sync"
)

type Video struct {
	Name     string
	URL      string
	CoverImg string
	Source   string
}

type VideoList struct {
	mu     sync.RWMutex
	Videos []Video
}

func (v *VideoList) MarshalVideosToBin() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	return buf.Bytes(), err
}

func (v *VideoList) UnmarshalBinToVideos(data []byte) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(v)
	return err
}
