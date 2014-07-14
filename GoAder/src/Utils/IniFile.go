//
// Author: huangchunping
// Date  : 2014-07-14
//
package Utils

import(
    //"errors"
    "bufio"
    //"unicode"
    //"unicode/utf8"
    "fmt"
    "io"
    "os" 
    "strings"
)

type IniFile struct {
    
}

func (f *IniFile) Close() {
}

func (f *IniFile) Get()(s string, result bool){
	return "ok",true
}

// It assumes that the '#' at the front of line (after skip space and tab characters) marks the line COMMENT flag.
func checkComment(data string) bool {
    if data[0] == '#' {
        return true 
	}
    return false
}

func isEqualitySign(r rune) bool{
    return r == '='
}

//
// Read and Parse the INI file.
func IniFileInit(file_name string) (f *IniFile, r bool) {
    ini_file := &IniFile{}

    input_file, input_error := os.Open(file_name)
    if input_error == nil {
        return ini_file, false
    }
    defer input_file.Close()
 
    reader := bufio.NewReader(input_file)
    for {
        line, reader_err := reader.ReadString('\n')
        if reader_err == io.EOF {
            break
        }
        line = strings.Trim(line, " \t")
        // process the comment
        if checkComment(line) == true {
            continue
        }

        equalsign_pos := strings.IndexFunc(line, isEqualitySign)
		if equalsign_pos == -1 {
			continue
		}
		key := line[:equalsign_pos]
		value := line[equalsign_pos:]
		fmt.Printf("line[%s], it's: key[%s], value[%s]\n", line, key, value)
    }

    return ini_file, true
}
