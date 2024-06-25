const go = new Go();

// WebAssembly コードをコンパイル、インスタンス化するためのAPI
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
  (result) => {
    go.run(result.instance);
  }
);

// 入力された数値を表示
function updateDisplay() {
  document.getElementById("base").textContent = getBase();
  document.getElementById("log").textContent = getLog();
}

// 数値ボタンが押された時の処理（ハンドラ）
function handleClick(funcName, arg) {
  // windowオブジェクトの関数を呼び出す
  if (arg !== undefined) {
    window[funcName](arg);
  } else {
    window[funcName]();
  }
  updateDisplay();

  // 出力時のアニメーション
  document.getElementById("input").classList.add("animate");
  setTimeout(
    () => document.getElementById("input").classList.remove("animate"),
    500
  );
}
