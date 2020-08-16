#include <stdio.h>
#include <unistd.h>

int sum(int num){
    int sum_num = 0;
    for(int i = 0; i < num; i++)
    {
        sum_num+=i;
    }
    return sum_num;
}

int main(int argc, char const *argv[])
{
    /* code */
    printf("%s\n",argv[1]);
    
    int num = (int)argv[1];
    int sum_num = sum(num);

    printf("sum is %d\n",sum_num);
    return 0;
}


    
