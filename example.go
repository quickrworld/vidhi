package main

var example = `
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

