package main

import "C"

import (
    "encoding/json"
    "fmt"
    "unsafe"
    "log"
    "io/ioutil"
    "os"
    "strconv"
    "github.com/golang/protobuf/proto"
)


//export get_student_enbuf
func get_student_enbuf(student_json_file *C.char) (unsafe.Pointer,C.int) {
    jsonFile, err := os.Open(C.GoString(student_json_file))
    if err != nil {
        log.Println("Error in reading json file",err)
        return nil,0
    }
    log.Println("Successfully Opened ",student_json_file)
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened jsonFile as a byte array.
    jsonbytes, _ := ioutil.ReadAll(jsonFile)
    var parsed_proto map[string]interface{}
    json.Unmarshal(jsonbytes, &parsed_proto)
    if parsed_proto==nil {
       log.Println("Error in reading json file",err)
       return nil,0
    }
       
    //Start Encode
    students := &Students{}
    var i int
    //Encode events_per_bundle events in a bundle
    for i=0;i<len(parsed_proto);i++ {
	    selectedjb, err := json.Marshal(parsed_proto[strconv.Itoa(i)])
	    if err != nil {
		log.Println("Error in reading json index",err)
		return nil,0
	    }
	    var parsed_jf map[string]interface{}
	    json.Unmarshal(selectedjb, &parsed_jf)
	    student := &Student{}
	    json.Unmarshal(selectedjb, student)
	    student_en, err := proto.Marshal(student)
	    if err != nil {
		log.Println("marshaling error in student : ", err)
                return nil,0
	    }
	    students.StudentEntry=append(students.StudentEntry,student_en)
    }
    students_en, err := proto.Marshal(students)
    if err != nil {
        log.Println("marshaling error in Students : ", err)
        return nil,0
    }
    return C.CBytes(students_en),C.int(int32(len(students_en)))
}

//export decode_student_enbuf
func decode_student_enbuf(data *C.char,leng C.int) *C.char {
    new_students := &Students{}
    student_en:=C.GoBytes(unsafe.Pointer(data),leng)
    err := proto.Unmarshal(student_en, new_students)
    if err != nil {
        log.Println("unmarshaling error in students : ", err)
        return nil
    }
    var i int
    var decoded_student string
    for i=0;i<len(new_students.StudentEntry);i++ {
	    new_student := &Student{}
	    err = proto.Unmarshal(new_students.StudentEntry[i], new_student)
	    if err != nil {
		log.Println("unmarshaling error in student : ", err)
                return nil
	    }
            decoded_student=decoded_student+"\n"+proto.MarshalTextString(new_student)
    }
    return C.CString(decoded_student)
}


func main() {
    jsonFile, err := os.Open("student_list.json")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened idps_event_input.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()
    // read our opened jsonFile as a byte array.
    jsonbytes, _ := ioutil.ReadAll(jsonFile)
    var parsed_proto map[string]interface{}
    json.Unmarshal(jsonbytes, &parsed_proto)
    if parsed_proto==nil {
       log.Println("Error in reading json file",err)
    }
       
    //Start Encode
    students := &Students{}
    var i int
    //Encode students
    for i=0;i<len(parsed_proto);i++ {
	    selectedjb, err := json.Marshal(parsed_proto[strconv.Itoa(i)])
	    if err != nil {
		log.Println("Error in reading json index",err)
	    }
	    var parsed_jf map[string]interface{}
	    json.Unmarshal(selectedjb, &parsed_jf)
	    student := &Student{}
	    json.Unmarshal(selectedjb, student)
	    student_en, err := proto.Marshal(student)
	    if err != nil {
		log.Println("marshaling error in student : ", err)
	    }
	    students.StudentEntry=append(students.StudentEntry,student_en)
    }
    students_en, err := proto.Marshal(students)
    if err != nil {
        log.Println("marshaling error in students : ", err)
    }

    new_students := &Students{}
    err = proto.Unmarshal(students_en, new_students)
    if err != nil {
        log.Println("unmarshaling error in students : ", err)
    }
    var decoded_student string
    for i=0;i<len(new_students.StudentEntry);i++ {
	    new_student := &Student{}
	    err = proto.Unmarshal(new_students.StudentEntry[i], new_student)
	    if err != nil {
		log.Println("unmarshaling error in student : ", err)
	    }
            decoded_student=decoded_student+"\n"+proto.MarshalTextString(new_student)
    }
    fmt.Println(decoded_student)
}
