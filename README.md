# simple-calculator

- simple calculator by Wasm in the Go language

# run on docker

```
docker build -t app . && docker run -it --rm --name app -p 8080:8080 app
```

# Function

- 入力
  - 最大４桁
  - 整数
- 二項演算
  - 足し算(add)
  - 引き算(subtract)
  - 掛け算(multiply)
  - 割り算(divide)
