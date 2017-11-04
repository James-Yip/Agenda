package util
import "bytes"
import "strings"
// format: yyyy-mm-dd/hh:mm
func Time2str(timeStr string) string {
	var a = ""
	var buff =  bytes.NewBufferString(a)
	
	for i:=0;i<16;i++ {
		if(i!=4&&i!=7&&i!=10&&i!=13) {
			buff.WriteByte(timeStr[i])
		}
	}
	a=buff.String()
	return a;
}

func Str2time(time string) string {
	var a = ""
	var buff =  bytes.NewBufferString(a)
	for i:=0;i<16;i++ {
		switch i {
		case 4,7:
			buff.WriteByte('-')
		case 5,6:
			buff.WriteByte(time[i-1])
		case 8,9:
			buff.WriteByte(time[i-2])
		case 10:
			buff.WriteByte('/')
		case 11,12:
			buff.WriteByte(time[i-3])
		case 13:
			buff.WriteByte(':')
		case 14,15:
			buff.WriteByte(time[i-4])
		default:
			buff.WriteByte(time[i])
		}
	}
	a=buff.String()
	return a;
}

func IsTimeValid(time string) bool {
	a := false
	if (len(time)!=16) {
		return a
	}
	for i:=0;i<16;i++ {
		switch i {
		case 4,7:
			if time[i]!='-' { return a }
		case 10:
			if time[i]!='/'{ return a }
		case 13:
			if time[i]!=':'{ return a }
		default:
			if (time[i]< '0' || time[i]> '9' ) { return a }
		}
	}
	a=true
	return a
}

// input: "a b c"
// output: [a b c]
func Str2slice(participator_str string) []string {
	return strings.Fields(participator_str)
}
