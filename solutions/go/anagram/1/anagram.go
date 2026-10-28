package anagram

import (
    "fmt"
	"sort"
    "strings"
)

func Delete(index int,candidates []string)[]string{
    candidates=append(candidates[:index],candidates[index+1:]...)
    return candidates
}

func Detect(subject string, candidates []string) []string {
    subjects:=strings.Split(strings.ToLower(subject),"")
    sort.Strings(subjects)
	for i:=0;i<len(candidates);i++{
        if len(subject)!=len(candidates[i]){
            candidates=Delete(i,candidates)
            i--
            continue
        }
        candidate:=strings.Split(strings.ToLower(candidates[i]),"")
        sort.Strings(candidate)
        fmt.Println(strings.Join(subjects,""))
        fmt.Println(strings.Join(candidate,""))
        if ((strings.Join(subjects,""))!=(strings.Join(candidate,""))){
            candidates=Delete(i,candidates)
            i--
            continue
        }
        if strings.ToLower(subject) == strings.ToLower(candidates[i]){
            candidates=Delete(i,candidates)
            i--
            continue
        }
    }
	return candidates
}
