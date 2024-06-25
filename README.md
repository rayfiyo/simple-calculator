# simple-calculator

- simple calculator by Wasm in the Go language

# try

## on docker

```
docker build -t app . && docker run -it --rm --name app -p 8080:8080 app
```

## on render (hosting service)

- The server may be shut down after 2024/06/25

https://simple-calculator-ygrt.onrender.com

# function

- 入力
  - オーバーフローしない範囲（-3037000500 ~ +3037000499）
  - 整数
- 二項演算
  - 足し算(add)
  - 引き算(subtract)
  - 掛け算(multiply)
  - 割り算(divide)
