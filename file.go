package worklog

import "io/ioutil"

import "gopkg.in/yaml.v2"

// FileRead is file reader
var FileRead = ioutil.ReadFile

// FileWrite is file writer
var FileWrite = ioutil.WriteFile

// ReadFile reads given file and unmarshal it to out
func ReadFile(filename string, out interface{}) error {
	b, e := FileRead(filename)
	if e != nil {
		return e
	}

	e = yaml.Unmarshal(b, out)
	return e
}

// WriteFile marshal writes given in
func WriteFile(filename string, in interface{}) error {
	b, e := yaml.Marshal(in)
	if e != nil {
		return e
	}

	e = FileWrite(filename, b, 0644)
	return e
}
