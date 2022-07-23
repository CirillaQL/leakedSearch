package model

import (
	"bytes"
	"encoding/gob"
)

type Video struct {
	Name     string
	URL      string
	CoverImg string
	Source   string
}

type Videos struct {
	Videos []Video
}

func (v *Videos) MarshalVideosToBin() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	return buf.Bytes(), err
}

func (v *Videos) UnmarshalBinToVideos(data []byte) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(v)
	return err
}
