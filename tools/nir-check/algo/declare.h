typedef struct {
    int field1;
    short field2;
    float field3;
} arg_of_awesome_func1;

typedef struct {
    int field1;
    short field2;
    float field3;
} arg_of_awesome_func2;

typedef struct {
    int field1;
    short field2;
    float field3;
} arg_of_awesome_func3;

typedef struct{
  int field1;
} arg_of_awesome_func;

typedef struct {
  int field1;
  char field2;
  char field3;
}test_args;

typedef struct {
  unsigned char field1;
  int field2;
  short field3;
}padding_default;

typedef struct __attribute__((packed)) {
  char field1;
  unsigned int field2;
  short field3;
} padding_packed;


