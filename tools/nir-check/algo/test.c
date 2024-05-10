
#include <stdio.h>
#include <limits.h>
#include "./declare.h"



void test_struct_arg_of_awesome_func1 () {
	
	arg_of_awesome_func1 A;


	A.field1 = INT_MAX +1;
	long szl_field1 = INT_MAX +1;

	if ((float)szl_field1 != (float)A.field1){
		printf("field1 sign error");
	}

	szl_field1 <<= 16;
	A.field1 <<= 16;

	if (szl_field1 != A.field1){
		printf("field1 size error");
	}

	A.field2 = INT_MAX +1;
	long szl_field2 = INT_MAX +1;

	if ((float)szl_field2 != (float)A.field2){
		printf("field2 sign error");
	}

	szl_field2 <<= 16;
	A.field2 <<= 16;

	if (szl_field2 != A.field2){
		printf("field2 size error");
	}

	A.field3 = INT_MAX +1;
	char szl_field3 = INT_MAX +1;

	if ((float)szl_field3 != (float)A.field3){
		printf("field3 sign error");
	}

	szl_field3 <<= 16;
	A.field3 <<= 16;

	if (szl_field3 != A.field3){
		printf("field3 size error");
	}
}


void test_struct_arg_of_awesome_func2 () {
	
	arg_of_awesome_func2 A;


	A.field1 = INT_MAX +1;
	char szl_field1 = INT_MAX +1;

	if ((float)szl_field1 != (float)A.field1){
		printf("field1 sign error");
	}

	szl_field1 <<= 16;
	A.field1 <<= 16;

	if (szl_field1 != A.field1){
		printf("field1 size error");
	}

	A.field2 = INT_MAX +1;
	int szl_field2 = INT_MAX +1;

	if ((float)szl_field2 != (float)A.field2){
		printf("field2 sign error");
	}

	szl_field2 <<= 16;
	A.field2 <<= 16;

	if (szl_field2 != A.field2){
		printf("field2 size error");
	}

	A.field3 = INT_MAX +1;
	char szl_field3 = INT_MAX +1;

	if ((float)szl_field3 != (float)A.field3){
		printf("field3 sign error");
	}

	szl_field3 <<= 16;
	A.field3 <<= 16;

	if (szl_field3 != A.field3){
		printf("field3 size error");
	}
}


void test_struct_arg_of_awesome_func3 () {
	
	arg_of_awesome_func3 A;


	A.field1 = INT_MAX +1;
	int szl_field1 = INT_MAX +1;

	if ((float)szl_field1 != (float)A.field1){
		printf("field1 sign error");
	}

	szl_field1 <<= 16;
	A.field1 <<= 16;

	if (szl_field1 != A.field1){
		printf("field1 size error");
	}

	A.field2 = INT_MAX +1;
	int szl_field2 = INT_MAX +1;

	if ((float)szl_field2 != (float)A.field2){
		printf("field2 sign error");
	}

	szl_field2 <<= 16;
	A.field2 <<= 16;

	if (szl_field2 != A.field2){
		printf("field2 size error");
	}

	A.field3 = INT_MAX +1;
	char szl_field3 = INT_MAX +1;

	if ((float)szl_field3 != (float)A.field3){
		printf("field3 sign error");
	}

	szl_field3 <<= 16;
	A.field3 <<= 16;

	if (szl_field3 != A.field3){
		printf("field3 size error");
	}
}

