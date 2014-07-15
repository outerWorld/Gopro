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

type SegmentData struct {
    key string
    values []string     // the key may be mapped to multi-values
}

type IniFile struct {
    sections map[string] []SegmentData
}

func (f *IniFile) Close() {
    fmt.Printf("IniFile object destruct\n")
}

func (f *IniFile) Get(section_name, key_name string)(value string, result bool){
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


// str is utf8-encoded.
// character escaped, 0x00<-->\00,0x01<-->\01 ...., \\<-->\,
// \a<-->0x07,\b<-->0x08, \t<-->0x09, \n-->0x0a, \v<-->0x0b, \f<-->0x0c, \r<-->0x0d
func stringEscape(str string) string {
    hex_digit_map := map[rune]byte {
        '0': 0x00, '1': 0x01, '2': 0x02, '3': 0x03, '4': 0x04, '5': 0x05, '6': 0x06, '7': 0x07, '8': 0x08, '9':0x09,
        'A':0x0a, 'B': 0x0b, 'C': 0x0c, 'D': 0x0d, 'E': 0x0e, 'F': 0x0f, 
        'a':0x0a, 'b': 0x0b, 'c': 0x0c, 'd': 0x0d, 'e': 0x0e, 'f': 0x0f,
    }
    temp_buf := make([]rune, len(str)) 

    slash_sign_flag := false
    var first_hex_char rune = -1
    escaped_len := 0
    for _, rune_char := range str {
        switch rune_char {
            case '\\':
                if slash_sign_flag == true {
                    temp_buf[escaped_len] = rune_char
                    escaped_len += 1
                    // like the case: \a\\
                    if first_hex_char != -1 {
                        temp_buf[escaped_len] = first_hex_char // restore previous character to temp_buf
                        escaped_len += 1
                        first_hex_char = -1
                        slash_sign_flag = true // retain the flag of a new slash.
                    } else {
                        slash_sign_flag = false
                    }
                } else {
                    slash_sign_flag = true
                }
            case '0', '1', '2', '3', '4','5', '6', '7', '8', '9',
                 'A', 'B', 'C', 'D', 'E', 'F',
                 'a', 'b', 'c', 'd', 'e', 'f':
                if slash_sign_flag == false {
                    temp_buf[escaped_len] = rune_char
                    escaped_len += 1
                    continue 
                }
                if first_hex_char != -1 {
                    hex_l4, _ := hex_digit_map[rune_char]
                    hex_h4, _ := hex_digit_map[first_hex_char]
                    temp_buf[escaped_len] = rune(((hex_h4 & 0x0F) << 4) | (hex_l4 & 0x0F))
                    escaped_len += 1
                    slash_sign_flag = false
                    first_hex_char = -1
                } else {
                    first_hex_char = rune_char
                }
            default:
                // like the case: \!
                if slash_sign_flag == true {
                    temp_buf[escaped_len] = '\\'
                    escaped_len += 1
                    slash_sign_flag = false
                }
                // like the case: \a! 
                if first_hex_char != -1 {
                    temp_buf[escaped_len] = first_hex_char
                    escaped_len += 1
                    first_hex_char = -1
                }
                temp_buf[escaped_len] = rune_char
                escaped_len += 1
        }
    }

    return string(temp_buf)
}

//
// Read and Parse the INI file.
func IniFileInit(file_name string) (f *IniFile, r bool) {
    ini_file := &IniFile{123}

    input_file, input_error := os.Open(file_name)
    if input_error != nil {
        return ini_file, false
    }
    defer input_file.Close()
 
    section_name := ""
    reader := bufio.NewReader(input_file)
    for {
        line, reader_err := reader.ReadString('\n')
        if reader_err == io.EOF {
            break
        }
        line = strings.Trim(line, " \t\n")
        // process the comment
        if len(line) <= 0 || checkComment(line) == true {
            continue
        }
        if line[0] == '[' && line[len(line)-1] == ']' {
            section_name = line[1:len(line)-1]
            continue
        }
        
        equalsign_pos := strings.IndexFunc(line, isEqualitySign)
		if equalsign_pos == -1 {
			continue
		}
		key := line[:equalsign_pos]
		value := line[equalsign_pos+1:]
        key = strings.Trim(key, " \t")
        value = stringEscape(strings.Trim(value, " \t"))
		fmt.Printf("section[%s] line[%s], it's: key[%s], value[%s]\n", section_name, line, key, value)
    }

    return ini_file, true
}
