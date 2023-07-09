package srcmodify

import (
    "regexp"
)

func DeleteLine(src []string) []string {

    re := regexp.MustCompile(`^\#line`)
    var res []string
    var blank bool
    for _, line := range src {
        if re.MatchString(line) {
            blank = true
        }else if blank {
            blank = false
        }else{
            res = append(res, line)
        }
    }
    return res
}
