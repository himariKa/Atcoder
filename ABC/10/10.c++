#include <iostream>
#include <string>
#include <vector>
using namespace std;

// 0:下、1:右、2：上、3:左、4:右下、5:右上、6:左上、7:左下
vector<int> DX = {1,0,-1,0,1,-1,-1,-1,1};
vector<int> DY = {1,0,-1,0,1,-1,-1,-1,1};

int main(){
    // H × Wの盤面を入力
    int H , W;
    cin >> H >> W;
    // 配列SにHの情報を入れる
    vector<string> S(H);
    // rowに入れていく
    for (auto& row:S){
        cin >> row;
    }

    //各ます(i,j)を順に処理
    for(int i = 0;i < H;i++){
        for(int j = 0;j<W;j++){
            // 空きマス以外はそのままにする
            if(S[i][j] != '.'){
                continue;
            }
            //周囲8ますの'#'の個数を数える
            int counter = 0;
            for(int d = 0;d < 8; ++d){
                // マス(i,j)の周囲のマスを(ni,nj)とする
                int ni = i + DX[d];
                int nj = j + DY[j];

                // マス(ni,nj)が盤面外の場合はスキップ
                if (ni < 0 || ni >= H || nj < 0 || nj >= W){
                    continue;
                }

                // #であれば1増やす
                if(S[ni][nj]=='#'){
                    ++counter;
                }
            }

            // マス(i,j)に個数をchar型に変換して記録
            S[i][j] = '0' + counter;
        }
        // 出力
        for(auto row : S){
            cout << row << endl;
        }
    }
    
}