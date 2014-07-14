//
// Author: huangchunping
// Date  : 2014-07-14
//

package Utils

import(
	"fmt"
	"strings"
	"testing"
)

//func isEqualitySign(r rune) bool {
//	return r == '='
//}

func TestIniFile(t *testing.T) {
	var test_str = " abc = 1234"
	equalsign_pos := strings.IndexFunc(test_str, isEqualitySign)
	if equalsign_pos == -1 {
		t.Errorf("can not find equality sign =\n")
		return
	}
	key := test_str[:equalsign_pos]
	value := test_str[equalsign_pos:]
	fmt.Printf("key[%s], value[%s]\n", key, value)
}
