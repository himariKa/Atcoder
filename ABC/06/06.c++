#include <iostream>
#include <vector>
using namespace std;

int main(){
    // 入力データの個数を表す変数Nへの入力
    int N;
    cin >> N;

    // サイズNの配列を宣言（各要素はint型）
    vector<int> A(N);
    // 配列に標準入力を入れる
    for (int i = 0; i <= N; i++){
        cin >> A[i];        
    }

    // 操作回数
    int counter = 0;
    
    // 処理が行えるかどうか（偶数かどうか）を確認
    bool can_do = true;
    while(true){
        // 配列Aの各要素へ入力
        for (int i = 0; i <= N; i++){
        // もし偶数ではなかったらfalseを返して処理を終了する
            if (A[i] % 2 == 1)
            {
                can_do = false;
            }
        }

        // 操作を行えない場合は反復を打ち切る
        if (!can_do){
            break;
        }

        // 操作を行えるならば操作を行う
        for (int i = 0; i<=N;i++){
            A[i] /= 2;
        }

        // counterを1増やす
        ++counter;
    }

    cout << counter << endl;
}