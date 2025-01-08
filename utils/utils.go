package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/wishicorp/sdk/helper/errutil"
	"io"
	"io/ioutil"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func JSONMethod(content interface{}) map[string]interface{} {
	var name map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}

// EncodeJSON encodes the given object into a JSON byte array.
func EncodeJSON(in interface{}) ([]byte, error) {
	if in == nil {
		return nil, fmt.Errorf("input for encoding is nil")
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(in); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeJSON(data []byte, out interface{}) error {
	if data == nil || len(data) == 0 {
		return fmt.Errorf("'data' being decoded is nil")
	}
	if out == nil {
		return fmt.Errorf("output parameter 'out' is nil")
	}

	// Decompress the data if it was compressed in the first place
	decompressedBytes, uncompressed, err := Decompress(data)
	if err != nil {
		return errutil.Wrapf("failed to decompress JSON: {{err}}", err)
	}
	if !uncompressed && (decompressedBytes == nil || len(decompressedBytes) == 0) {
		return fmt.Errorf("decompressed data being decoded is invalid")
	}

	// If the input supplied failed to contain the compression canary, it
	// will be notified by the compression utility. Decode the decompressed
	// input.
	if !uncompressed {
		data = decompressedBytes
	}

	return DecodeJSONFromReader(bytes.NewReader(data), out)
}

// DecodeJSONFromReader  the given io.Reader pointing to a JSON, into a desired object
func DecodeJSONFromReader(r io.Reader, out interface{}) error {
	if r == nil {
		return fmt.Errorf("'io.Reader' being decoded is nil")
	}
	if out == nil {
		return fmt.Errorf("output parameter 'out' is nil")
	}

	dec := json.NewDecoder(r)

	// While decoding JSON values, interpret the integer values as `json.Number`s instead of `float64`.
	dec.UseNumber()

	// Since 'out' is an interface representing a pointer, pass it to the decoder without an '&'
	return dec.Decode(out)
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func GetFileContentAsStringLines(filePath string) ([]string, error) {
	fmt.Println("get file content as lines: ", filePath)
	result := []string{}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file error: ", err)
		return result, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, ";") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	fmt.Println("get file content as lines size: ", len(result))
	return result, nil
}

func Print(m map[string]interface{}) {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", float64(vv))
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case nil:
			fmt.Println(k, "is nil", "null")
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			Print(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle ", fmt.Sprintf("%T", v))
		}
	}
}

func GenPK(v interface{}, len int) string {
	str := fmt.Sprintf("%v", v)
	hasher := md5.New()
	hasher.Write([]byte(str))
	a := hex.EncodeToString(hasher.Sum(nil))
	b := a[:len]
	return b
}

func GenDefaultPK(v interface{}) string {
	return GenPK(v, 4)
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func checkIp(ipStr string) bool {
	address := net.ParseIP(ipStr)
	if address == nil {
		// fmt.Println("ip地址格式不正确")
		return false
	} else {
		// fmt.Println("正确的ip地址", address.String())
		return true
	}
}

// ip to int64
func inetAton(ipStr string) int64 {
	bits := strings.Split(ipStr, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

// int64 to IP
func inetNtoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

func IsInnerIp(ipStr string) bool {
	if !checkIp(ipStr) {
		return false
	}
	inputIpNum := inetAton(ipStr)
	innerIpA := inetAton("10.255.255.255")
	innerIpB := inetAton("172.16.255.255")
	innerIpC := inetAton("192.168.255.255")
	innerIpD := inetAton("100.64.255.255")
	innerIpF := inetAton("127.255.255.255")

	return inputIpNum>>24 == innerIpA>>24 || inputIpNum>>20 == innerIpB>>20 ||
		inputIpNum>>16 == innerIpC>>16 || inputIpNum>>22 == innerIpD>>22 ||
		inputIpNum>>24 == innerIpF>>24
}

func IToBool(i int) bool {
	switch i {
	case 0:
		return false
	case 1:
		return true
	default:
		return true
	}
}

func AToBool(str string) (bool, error) {
	switch strings.ToLower(str) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf("unable to parse %s as boolean", str)
	}
}

func ParseTimeString(timeStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

// AssignSimilarFields 使用反射将一个结构体的字段赋值给另一个结构体
func AssignSimilarFields(src interface{}, dest interface{}) {
	srcVal := reflect.ValueOf(src).Elem()   // 获取源结构体值
	destVal := reflect.ValueOf(dest).Elem() // 获取目标结构体值

	for i := 0; i < srcVal.NumField(); i++ {
		fieldName := srcVal.Type().Field(i).Name // 获取字段名称

		// 在目标结构体中查找同名字段
		destField := destVal.FieldByName(fieldName)
		if destField.IsValid() && destField.CanSet() {
			// 赋值
			destField.Set(srcVal.Field(i))
		}
	}
}

// CopyFields 使用反射将一个结构体的字段赋值给另一个结构体，前提是字段具备相同的名字和类型
func CopyFields[T any](src interface{}) *T {
	dst := new(T) // see golang issue #48399
	AssignSimilarFields(src, dst)
	return dst
}

// ValueToPoint 将一个值转换为指针
func ValueToPoint[Type any](v Type) *Type {
	return &v
}

func CheckPointer[T any](v *T, defaultValue T) *T {
	if v == nil {
		return &defaultValue
	}
	return v
}

func TruncateString(s *string, m int) *string {
	_s := CheckPointer[string](s, "") //检查空指针
	if len(*_s) > m {
		*_s = (*_s)[:m]
	}
	return s
}
