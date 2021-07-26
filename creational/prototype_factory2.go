package creational

import (
	"bytes"
	"encoding/gob"
)

func (c *Candidate) DeepCopy() *Candidate {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(c)

	d := gob.NewDecoder(&b)
	n := Candidate{}
	_ = d.Decode(&n)
	return &n
}

var labourCandidate = Candidate{
	"", Labour, Constituency{"", "NSW"},
}

var liberalCandidate = Candidate{
	"", Liberal, Constituency{"", "NSW"},
}

var greensCandidate = Candidate{
	"", Greens, Constituency{"", "QLD"},
}

func NewCandidate(proto *Candidate, name string, constituency string) *Candidate {
	result := proto.DeepCopy()
	result.Name = name
	result.Constituency.Name = constituency
	return result
}

func NewLabourCandiadate(name string, constituency string) *Candidate {
	return NewCandidate(&labourCandidate, name, constituency)
}

func NewLiberalCandiadate(name string, constituency string) *Candidate {
	return NewCandidate(&liberalCandidate, name, constituency)
}

func NewGreensCandiadate(name string, constituency string) *Candidate {
	return NewCandidate(&greensCandidate, name, constituency)
}
