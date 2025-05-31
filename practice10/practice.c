#include <stdio.h>
#include <string.h>

int main()
{
  /* 1번 */
  int input_data = getchar(); // Read a character from standard input
  printf("You entered: %c\n", input_data);

  /* 2번 */
  char input_string[100]; // Declare a character array to hold the string
  if (fgets(input_string, sizeof(input_string), stdin) != NULL)
  {
    // fgets는 개행 문자(\n)도 함께 읽어올 수 있으므로, 필요하면 제거합니다.
    input_string[strcspn(input_string, "\n")] = '\0'; // 개행 문자를 찾아 널 문자로 대체

    printf("입력한 내용: %s\n", input_string);
  }
  else
  {
    printf("입력 중 오류가 발생했습니다.\n");
  }
}