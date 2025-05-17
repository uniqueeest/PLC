#include <stdio.h>

/* 지정한 비트 0으로 만듦. AND 연산자 이용 */
unsigned char ResetBit(unsigned char dest_data, unsigned char bit_num)
{
  if (bit_num < 8)
  {
    dest_data = dest_data & ~(1 << bit_num);
  }
  return dest_data;
}

/* 지정한 비트 1로 만듦. OR 연산자 이용 */
unsigned char SetBit(unsigned char dest_data, unsigned char bit_num)
{
  if (bit_num < 8)
  {
    dest_data = dest_data | (1 << bit_num);
  }
  return dest_data;
}

/* 특정 비트의 값 얻기. AND 연산자 이용 */
unsigned char GetBit(unsigned char dest_data, unsigned char bit_num)
{
  unsigned char bit_value = 0;
  if (bit_num < 8)
  {
    bit_value = dest_data & (0x01 << bit_num);
    bit_value = bit_value >> bit_num;
  }
  return bit_value;
}

int main()
{
  unsigned char lamp_state = 0xFF, bit_state;
  printf("%X->", lamp_state);
  for (int i = 0; i < 8; i++)
  {
    bit_state = GetBit(lamp_state, 7 - i);
    printf("%d", bit_state);
  }
  printf("\n");

  unsigned char lamp_state2 = 0xFF;
  printf("%X->", lamp_state2);
  printf("%X\n", ResetBit(lamp_state2, 3));

  unsigned char lamp_state3 = 0x33;
  printf("%X->", lamp_state3);
  printf("%X\n", SetBit(lamp_state3, 3));
}