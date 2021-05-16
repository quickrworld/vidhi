package vidhi

import "fmt"

type Arg struct {
	Name  string
	Value interface{}
}

func (a Arg) String() string {
	return fmt.Sprintf("vidhi.Arg: { Name: %s, Value: %v }", a.Name, a.Value)
}

type Function struct {
	Name string
	Args []Arg
}

func (f Function) String() string {
	return fmt.Sprintf("Function: { Name: %s, Args: %s", f.Name, f.Args)
}

type RuleSet struct {
	Conjunction string
	Rules       []interface{}
}

func (r *RuleSet) String() string {
	return fmt.Sprintf("RuleSet: { Conjunction: %s, Rules: %v }", r.Conjunction, r.Rules)
}