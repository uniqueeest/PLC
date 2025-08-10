#include <stdio.h>
#include <string.h>

/* 문자열 숫자 변환 */
int convert_string_to_number(char *str)
{
  int length = strlen(str);
  int number = 0;

  for (int i = 0; i < length; i++)
  {
    int digit = str[i] - '0';
    number = number * 10 + digit;
  }

  return number;
}

/* 문자열 숫자 변환2 */
int convert_string_to_number2(char *str)
{
  int number = 0;
  while (*str)
  {
    number = number * 10 + (*str - '0');
    str++;
  }
  return number;
}

int main()
{
  /* 1번 */
  char *num_string = "123";
  printf("문자열 숫자 변환: %d\n", convert_string_to_number(num_string));

  /* 2번 */
  char *num_string2 = "123123123";
  printf("문자열 숫자 변환2: %d\n", convert_string_to_number2(num_string2));

  /* 3번 */
  int first_num, second_num;
  char first_string[100];
  char second_string[100];

  printf("first: ");
  if (fgets(first_string, sizeof(first_string), stdin) != NULL)
  {
    size_t len = strlen(first_string);
    if (len > 0 && first_string[len - 1] == '\n')
    {
      first_string[len - 1] = '\0';
    }
  }
  else
  {
    printf("입력 오류!\n");
    return 1; // 오류 발생 시 종료
  }

  printf("second: ");
  if (fgets(second_string, sizeof(second_string), stdin) != NULL)
  {
    size_t len = strlen(second_string);
    if (len > 0 && second_string[len - 1] == '\n')
    {
      second_string[len - 1] = '\0';
    }
  }
  else
  {
    printf("입력 오류!\n");
    return 1; // 오류 발생 시 종료
  }

  first_num = convert_string_to_number(first_string);
  second_num = convert_string_to_number(second_string);

  printf("%d + %d = %d\n", first_num, second_num, first_num + second_num);

  return 0;
}