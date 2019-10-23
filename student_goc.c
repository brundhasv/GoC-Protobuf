#include <stdio.h>
#include "student.h"
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

struct get_student_enbuf_return stu;

int main() {
    stu = get_student_enbuf("student_list.json");
    int stu_list_len = stu.r1;
    if(stu.r0!=NULL) {
    	char *decoded_stu = decode_student_enbuf((char*)stu.r0,stu.r1);
    	if(decoded_stu != NULL){
    		printf("%s",decoded_stu);
		free((char*)decoded_stu);
    	}
    free((char*)stu.r0);
    }
    return 0;
}
