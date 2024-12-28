#include <iostream>
using namespace std;

int main(){
    // AとBとCの型を定義
    int A,B,C;
    // AとBの標準入力を受け付ける
    cin >> A >> B;
    
    // CはAとBの掛け算の結果
    C = A * B;
    if (C % 2 == 0){
        // Cが偶数の時
        cout << "Even" << endl;
    }else{
        // Cが奇数の時
        cout << "Odd" << endl;
    }
}