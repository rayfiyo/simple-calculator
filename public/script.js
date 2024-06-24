const go = new Go();

// WebAssembly コードをコンパイル、インスタンス化するためのAPI
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
  (result) => {
    go.run(result.instance);
  }
);

// 入力された数値を表示
function updateDisplay() {
  document.getElementById("currentInput").textContent = getCurrentInput();
  document.getElementById("previousInput").textContent = getPreviousInput();
}

// 数値ボタンが押された時の処理（ハンドラ）
function handleClick(funcName, arg) {
  // 演算結果の出力
  if (arg !== undefined) {
    window[funcName](arg);
  } else {
    window[funcName]();
  }
  updateDisplay();

  // 出力時のアニメーション
  document.getElementById("currentInput").classList.add("animate");
  setTimeout(
    () => document.getElementById("currentInput").classList.remove("animate"),
    500
  );
}
