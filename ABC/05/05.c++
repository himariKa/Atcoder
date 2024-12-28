#include <iostream>
using namespace std;

//各桁の合計を求める関数
int calc_sum_digit(int n){
    // 各桁の数の合計を入れる変数
    int sum_digit = 0;

    // Nの各桁の和を足し算する
    while(n > 0){
        sum_digit += n % 10;
        n /= 10;
    }

    return sum_digit;
}

int main(){
    // NとAとBの型を定義
    int N,A,B;
    // KとAとBの標準入力を受け付ける
    cin >> N >> A >> B;

    // 答えを格納する変数
    int result = 0;

    for (int i=1;i<=N;i++){
        // 各桁の合計を求める
        int x = calc_sum_digit(i);

        // もし各桁の合計がA以上B以下なら答えにその数字を足し算する
        if(A <= x && B >= x){
            result =+ i;
        }
    }
    
    cout << result << endl;
}