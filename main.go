package main

import (
	"encoding/json"
	"fmt"
	"plugin"
	"vidhi/vidhi"
)

var r = `
{
	"conjunction": "all",
	"rules": [
		{
			"function": "Contains",
			"Args": [
				{
					"Name": "s",
					"Value": "1abc2"
				},
				{
					"Name": "substr",
					"Value": "abc"
				}
			]
		},
		{
			"function": "HasPrefix",
			"Args": [
				{
					"Name": "s",
					"Value": "abc1"
				},
				{
					"Name": "substr",
					"Value": "abc"
				}
			]
		},
		{
			"conjunction": "any",
			"rules": [
				{
					"function": "HasSuffix",
					"Args": [
						{
							"Name": "s",
							"Value": "1abc"
						},
						{
							"Name": "substr",
							"Value": "abc"
						}
					]
				},
				{
					"function": "HasPrefix",
					"Args": [
						{
							"Name": "s",
							"Value": "abc1"
						},
						{
							"Name": "target",
							"Value": "abc"
						}
					]
				}
			]
		}
	]
}
`

func main() {
	funcs, err := makeFuncs()
	content := r
	m, err := makeMap(content)
	if err != nil {
		panic(err)
	}
	ruleSet := makeRuleSet(m)
	execRules(ruleSet, funcs)
}

func makeFunction(rule map[string]interface{}) vidhi.Function {
	function := vidhi.Function{}
	function.Name = rule["function"].(string)
	args := rule["Args"].([]interface{})
	for i := 0; i < len(args); i++ {
		switch args[i].(type) {
		case map[string]interface{}:
			ja := args[i].(map[string]interface{})
			arg := vidhi.Arg{}
			arg.Name = ja["Name"].(string)
			arg.Value = ja["Value"].(string)
			function.Args = append(function.Args, arg)
		default:
			fmt.Println("arg.(type) must be map[string]interface{}")
		}
	}
	return function
}


func execRules(ruleSet *vidhi.RuleSet, funcs map[string]plugin.Symbol) {
	for i := 0; i < len(ruleSet.Rules); i++ {
		switch ruleSet.Rules[i].(type) {
		case vidhi.Function:
			fn := ruleSet.Rules[i].(vidhi.Function)
			b, err := funcs[fn.Name].(func([]vidhi.Arg) (bool, error))(fn.Args)
			if err != nil {
				panic(err)
			}
			fmt.Println(b)
		case *vidhi.RuleSet:
			execRules(ruleSet.Rules[i].(*vidhi.RuleSet), funcs)
		default:
			fmt.Printf("unknown type: %s", ruleSet.Rules[i])
		}
	}
}

func makeFuncs() (map[string]plugin.Symbol, error) {
	// setup plugin functions
	p, err := plugin.Open("vidhi.so")
	if err != nil {
		panic(err)
	}

	funcs := map[string]plugin.Symbol{}

	// Contains
	ContainsFunc, err := p.Lookup("Contains")
	if err != nil {
		panic(err)
	}
	funcs["Contains"] = ContainsFunc

	// HasPrefix
	HasPrefixFunc, err := p.Lookup("HasPrefix")
	if err != nil {
		panic(err)
	}
	funcs["HasPrefix"] = HasPrefixFunc

	// HasSuffix
	HasSuffixFunc, err := p.Lookup("HasSuffix")
	if err != nil {
		panic(err)
	}
	funcs["HasSuffix"] = HasSuffixFunc
	return funcs, err
}

func makeMap(content string) (map[string]interface{}, error) {
	var f interface{}
	err := json.Unmarshal([]byte(content), &f)
	if err != nil {
		fmt.Printf("makeMap: Error unmarshalling: %s\n", err)
	}
	var m map[string]interface{}
	switch f.(type) {
	case map[string]interface{}:
		m = f.(map[string]interface{})
	default:
		return m, fmt.Errorf("makeMap: content type not handled %v\n", f)
	}
	return m, nil
}

func makeRuleSet(m map[string]interface{}) *vidhi.RuleSet {
	ruleSet := &vidhi.RuleSet{}
	conjunction, ok := m["conjunction"]
	if !ok {
		fmt.Println("no conjunction")
	}
	ruleSet.Conjunction = conjunction.(string)
	switch m["rules"].(type) {
	case []interface{}:
		// ok
	default:
		fmt.Println("rules.(type) must be []interface{}")
	}
	rules, ok := m["rules"].([]interface{})
	if !ok {
		fmt.Println("no rules")
		return ruleSet
	}
	processRules(rules, ruleSet)
	return ruleSet
}

func processRules(rules []interface{}, ruleSet *vidhi.RuleSet) {
	for i := 0; i < len(rules); i++ {
		rule := rules[i]
		switch rule.(type) {
		case map[string]interface{}:
			if isFunction(rule.(map[string]interface{})) {
				ruleSet.Rules = append(ruleSet.Rules, makeFunction(rule.(map[string]interface{})))
			} else if isConjunction(rule.(map[string]interface{})) {
				ruleSet.Rules = append(ruleSet.Rules, makeRuleSet(rule.(map[string]interface{})))
			}
		default:
			fmt.Println("rule.(type) must be map[string]interface{}")
		}
	}
}

func isFunction(rule map[string]interface{}) bool {
	_, ok := rule["function"]
	return ok
}

func isConjunction(rule map[string]interface{}) bool {
	_, ok := rule["conjunction"]
	return ok
}


