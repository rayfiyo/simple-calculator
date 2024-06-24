const go = new Go();
let wasmFunctions = {}; // WebAssembly で定義された関数を格納するオブジェクト
let previousInput = ""; // 1つ前の入力値（上段）
let currentInput = "0"; // 現在の入力値（下段）
let currentOperation = null; // 現在の演算種類

// WebAssembly コードをコンパイル、インスタンス化するためのAPI
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
  (result) => {
    go.run(result.instance);
    wasmFunctions.add = add;
    wasmFunctions.subtract = subtract;
    wasmFunctions.multiply = multiply;
    wasmFunctions.divide = divide;
  }
);

// 入力された数値を表示
function updateDisplay() {
  document.getElementById("currentInput").textContent = currentInput;
  document.getElementById("previousInput").textContent = previousInput;
}

// 桁を追加（数字を押した時の処理）
function appendNumber(number) {
  if (currentInput === "0") {
    currentInput = number.toString();
  } else {
    currentInput += number.toString();
  }
  updateDisplay();
}

// 演算記号をセット（演算記号を押した時の処理）
function setOperation(operation) {
  // すでに演算記号が設定されている
  // （演算記号を押す回数が2回以降の）場合は演算を行う
  if (currentOperation !== null) {
    calculate();
  }

  currentOperation = operation;
  previousInput = currentInput + " " + getOperationSymbol(operation);
  currentInput = "0";
  updateDisplay();
}

// 演算記号を取得
function getOperationSymbol(operation) {
  switch (operation) {
    case "add":
      return "+";
    case "subtract":
      return "-";
    case "multiply":
      return "×";
    case "divide":
      return "÷";
  }
}

// 表示のクリア（C ボタン）
function clearCalculator() {
  currentInput = "0";
  previousInput = "";
  currentOperation = null;
  updateDisplay();
}

// 演算（= ボタン）
function calculate() {
  if (currentOperation === null || previousInput === "") return;

  const prev = parseInt(previousInput);
  const current = parseInt(currentInput);
  let result;

  // 演算
  switch (currentOperation) {
    case "add":
      result = wasmFunctions.add(prev, current);
      break;
    case "subtract":
      result = wasmFunctions.subtract(prev, current);
      break;
    case "multiply":
      result = wasmFunctions.multiply(prev, current);
      break;
    case "divide":
      if (current === 0) {
        result = "Error";
      } else {
        result = wasmFunctions.divide(prev, current);
      }
      break;
  }

  // 演算結果の出力
  currentInput = result.toString();
  previousInput = "";
  currentOperation = null;
  updateDisplay();

  // 出力時のアニメーション
  document.getElementById("currentInput").classList.add("animate");
  setTimeout(
    () => document.getElementById("currentInput").classList.remove("animate"),
    500
  );
}
