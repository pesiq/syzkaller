
#include <stdio.h>
#include <limits.h>
#include "./declare.h"



void test_struct_test_args () {
	
	test_args A;


	A.field1 = INT_MAX + 1;
	int szl_field1 = INT_MAX + 1;

	if ((float)szl_field1 != (float)A.field1){
    printf("test_args: field1: sign error\n");
	}

	szl_field1 <<= 16;
	A.field1 <<= 16;

	if (szl_field1 != A.field1){
    printf("test_args: field1: size error\n");
	}

	A.field2 = CHAR_MAX + 1;
	char szl_field2 = CHAR_MAX + 1;

	if ((float)szl_field2 != (float)A.field2){
    printf("test_args: field2: sign error\n");
	}

	szl_field2 <<= 16;
	A.field2 <<= 16;

	if (szl_field2 != A.field2){
    printf("test_args: field2: size error\n");
	}

	A.field3 =  + 1;
	 szl_field3 =  + 1;

	if ((float)szl_field3 != (float)A.field3){
    printf("test_args: field3: sign error\n");
	}

	szl_field3 <<= 16;
	A.field3 <<= 16;

	if (szl_field3 != A.field3){
    printf("test_args: field3: size error\n");
	}
}


void call_tests(){ 
	test_struct_test_args();
}
