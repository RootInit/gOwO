package owo

import (
	"bufio"
	"io"
)

type Owofier struct {
	Replacements map[string][]byte
}

func DefaultOwofier() Owofier {
	return Owofier{map[string][]byte{
		"ac":  []byte("awc"),
		"ad":  []byte("awd"),
		"age": []byte("awge"),
		"ai":  []byte("awi"),
		"ant": []byte("awnt"),
		"ar":  []byte("awr"),
		"ard": []byte("awd"),
		"are": []byte("awr"),
		"ay":  []byte("ai"),
		"ble": []byte("bwl"),
		"bod": []byte("bwod"),
		"bou": []byte("bwou"),
		"bui": []byte("bwi"),
		"cal": []byte("cwal"),
		"com": []byte("cum"),
		"con": []byte("cawn"),
		"dif": []byte("dwif"),
		"dis": []byte("dwis"),
		"do":  []byte("dwo"),
		"ed":  []byte("wed"),
		"eir": []byte("ewr"),
		"en":  []byte("ewn"),
		"er":  []byte("ewr"),
		"ere": []byte("ewr"),
		"est": []byte("ewst"),
		"et":  []byte("ewt"),
		"fi":  []byte("fy"),
		"ful": []byte("fwul"),
		"fur": []byte("fwur"),
		"gan": []byte("gwan"),
		"ght": []byte("wgt"),
		"go":  []byte("gow"),
		"hap": []byte("hawp"),
		"ime": []byte("iwm"),
		"ing": []byte("wng"),
		"ion": []byte("iwn"),
		"it":  []byte("iwt"),
		"ith": []byte("wif"),
		"lol": []byte("lawl"),
		"ly":  []byte("lwy"),
		"mem": []byte("mwem"),
		"men": []byte("mwen"),
		"ne":  []byte("nwe"),
		"not": []byte("nawt"),
		"now": []byte("naw"),
		"oin": []byte("owin"),
		"ol":  []byte("owl"),
		"oll": []byte("ow"),
		"om":  []byte("owm"),
		"oo":  []byte("wo"),
		"ood": []byte("wod"),
		"ord": []byte("owrd"),
		"ore": []byte("owre"),
		"ort": []byte("owrt"),
		"ory": []byte("owry"),
		"ost": []byte("owst"),
		"oun": []byte("own"),
		"our": []byte("owr"),
		"par": []byte("pawr"),
		"peo": []byte("pwe"),
		"pic": []byte("pwic"),
		"ple": []byte("pwl"),
		"pro": []byte("pwo"),
		"ra":  []byte("rwa"),
		"ree": []byte("wee"),
		"ris": []byte("wis"),
		"sen": []byte("swen"),
		"so":  []byte("sew"),
		"tem": []byte("twm"),
		"ter": []byte("twr"),
		"the": []byte("te"),
		"tle": []byte("twle"),
		"to":  []byte("two"),
		"too": []byte("two"),
		"tud": []byte("twud"),
		"two": []byte("too"),
		"ty":  []byte("twy"),
		"ugh": []byte("uwgh"),
		"ul":  []byte("uwl"),
		"un":  []byte("uwn"),
		"ure": []byte("uwre"),
		"urh": []byte("uwgh"),
		"ute": []byte("woot"),
		"v":   []byte("w"),
		"win": []byte("wen"),
		"wor": []byte("wer"),
		"you": []byte("yu"),
	}}
}

func ConvertReplacements(r map[string]string) map[string][]byte {
	replacements := make(map[string][]byte, len(r))
	for pattern, replacement := range r {
		replacements[pattern] = []byte(replacement)
	}
	return replacements
}

func (o *Owofier) Translate(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		input := scanner.Bytes()
		output := make([]byte, 0, len(input)*2)
		for i := 0; i < len(input); {
			if i < len(input)-2 {
				// Get 3 bytes {1,2,3,4,5} => {1,2,3}
				replace, exist := o.Replacements[string(input[i:i+3])]
				if exist {
					output = append(output, replace...)
					i += 3
					continue
				}
			}
			if i < len(input)-1 {
				// Get 2 bytes {1,2,3,4,5} => {1,2}
				replace, exist := o.Replacements[string(input[i:i+2])]
				if exist {
					output = append(output, replace...)
					i += 2
					continue
				}
			}
			// Get 1 byte {1,2,3,4,5} => {1}
			replace, exist := o.Replacements[string(input[i])]
			if exist {
				output = append(output, replace...)
			} else {
				output = append(output, input[i])
			}
			i += 1
		}
		// Append new line to output
		output = append(output, []byte("\n")...)
		// Write output
		_, err := w.Write(output)
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *Owofier) Stats(r io.Reader) map[string]int {
	// Initialize map
	stats := make(map[string]int, len(o.Replacements))
	for key := range o.Replacements {
		stats[key] = 0
	}
	// Process input
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		input := scanner.Bytes()
		for i := 0; i < len(input); {
			if i < len(input)-2 {
				// Get 3 bytes {1,2,3,4,5} => {1,2,3}
				_, exist := o.Replacements[string(input[i:i+3])]
				if exist {
					stats[string(input[i:i+3])] += 1
					i += 3
					continue
				}
			}
			if i < len(input)-1 {
				// Get 2 bytes {1,2,3,4,5} => {1,2}
				_, exist := o.Replacements[string(input[i:i+2])]
				if exist {
					stats[string(input[i:i+2])] += 1
					i += 2
					continue
				}
			}
			// Get 1 byte {1,2,3,4,5} => {1}
			_, exist := o.Replacements[string(input[i])]
			if exist {
				stats[string(input[i])] += 1
			}
			i += 1
		}
	}
	return stats
}
