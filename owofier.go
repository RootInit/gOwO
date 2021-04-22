//CLI owofier.
//Copyright 2021, I_Mod_Things

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	s := string(in)

	s = strings.ToLower(s)

	replacements := map[string]string{
		"(ing|er|ai|ed|ion|ore|est|ol|it|ure|ese|et|ad|en|[ou]l)": "w${1}",
		"ac":                  "awc",
		"ant":                 "awnt",
		"ar":                  "awr",
		"ay":                  "ai",
		"ble":                 "bwul",
		"bod":                 "bwod",
		"bout":                "bwout",
		"cal":                 "cwal",
		"com":                 "cum",
		"coun":                "cown",
		"dif":                 "dwif",
		"dis":                 "dwis",
		"do":                  "dwo",
		"eir":                 "ewr",
		"ere":                 "ewr",
		"fi":                  "fy",
		"ful":                 "fwul",
		"fur":                 "fwur",
		"gan":                 "gwan",
		"ght":                 "wgt",
		"go":                  "gow",
		"hap":                 "hawp",
		"ith":                 "iff",
		"know":                "kno",
		"lol":                 "lawl",
		"ly":                  "lwy",
		"me":                  "mi",
		"mem":                 "mwem",
		"ment":                "went",
		"not":                 "nawt",
		"ood":                 "wod",
		"oo":                  "wo",
		"ost":                 "owst",
		"our":                 "owr",
		"par":                 "pawr",
		"peo":                 "pwe",
		"pic":                 "pwic",
		"ple":                 "pwl",
		"pro":                 "pwo",
		"ree":                 "wee",
		"sen":                 "swen",
		"so":                  "sew",
		"stud":                "stwud",
		"tence":               "twence",
		"the":                 "te",
		"tle":                 "twle",
		"to":                  "two",
		"ture":                "twur",
		"ty":                  "twy",
		"ute":                 "woot",
		"want":                "wnt",
		"ward":                "wawd",
		"you":                 "yu",
		`v|([^\saeou])r`:      "${1}w",
		`[a-z]([^aeiouUO])\.`: "${1} UwU.",
		`[a-z]([aeiou])\.`:    "${1} OwO.",
		`[a-z]([^aeiouUO])\,`: "${1} uwu,",
		`[a-z]([aeiou])\,`:    "${1} owo,",
	}

	for pattern, replacement := range replacements {
		regexstr, _ := regexp.Compile(pattern)
		s = regexstr.ReplaceAllString(s, replacement)
	}

	fmt.Println(s)
}
