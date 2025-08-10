## 포인터

### ptr과 \*ptr의 차이점

- `ptr = ` : 포인터 변수의 값(가리키는 대상의 주소)이 변경
- `*ptr = ` : 포인터가 가리키는 대상의 값이 변경

### 포인터를 사용하여 간접 주소 방식으로 값을 대입하는 이유

- 모든 변수가 같은 함수에 선언되는 것은 아니기 때문
- 일반 변수는 다른 함수에 있는 변수 사용 불가
- 포인터 변수는 다른 함수에 선언된 변수의 값을 읽거나 변경 가능

### 포인터로 다른 함수에 선언된 변수 사용

```c
void Test(short *ptr)
{
  short soft = 0;
  soft = *ptr;
  *ptr = 3;
}

int main()
{
  short tips = 5;
  Test(&tips);
}
```

### const 키워드로 주소 변경 실수 막기

- const 키워드를 이용하여 피호출자에서 호출자로 전달받은 주소를 변경하는 실수를 방지

```c
void swap(int * const pa, int * const pb)
{
  int temp = *pa;
  pa = pb /* 에러 발생 */
  *pb = temp;
}
```
