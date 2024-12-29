#include <iostream>
#include <vector>
using namespace std;

int main(){
    // 入力データの個数を表す変数Nへの入力
    int N;
    cin >> N;

    // サイズNの配列Aを宣言（各要素はint型）
    vector<int> A(N);
    // 配列に標準入力を入れる
    for (auto& ai : A){
        cin >> ai;        
    }

    // 配列Aを大きい順に並び替える
    sort(A.begin(),A.end(),greater<int>());

    // 答えを表す変数
    int result = 0;

    // 配列Aの各要素を交互に足し引きする
    for (int i = 0;i < N;++i){
        if (i % 2 == 0){
            //アリスが取得した数
            result += A[i];
        }else{
            // ボブが取得した数
            result -= A[i];
        }
    }

    cout << result << endl;
}