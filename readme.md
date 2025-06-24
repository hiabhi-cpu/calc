
# 🧮 Calc — Command-Line Calculator

**Calc** is a simple command-line calculator implemented in **Go**. It supports infix expressions and uses a **stack** to:

✅ Convert infix expressions to **postfix** notation  
✅ Evaluate the resulting **postfix** expression  

It supports basic operators: `+`, `-`, `*`, and `/`, as well as parenthesis `()`.

---

## ⚡️ Features
- Evaluate expressions passed as command-line arguments
- Support for parenthesis and operator precedence
- Implements the classic **Shunting Yard Algorithm** (`toPostFix.go`) and **Postfix Evaluation** (`solvePostfix.go`)
- Leverages a custom `stack` package from `github.com/hiabhi-cpu/collections`

---

## 📂 Project Structure
```bash
calc/
├─ main.go # Entry point for the calculator
├─ solvePostfix.go # Evaluate the generated postfix expression
├─ toPostFix.go # Convert infix expression to postfix
├─ go.mod
├─ go.sum
```

---

## 🛠️ Installation
Clone the repository:
```bash
git clone https://github.com/hiabhi-cpu/calc.git
cd calc
go mod tidy
```

### ⚡️ Usage
Pass an infix expression wrapped in quotes:

```bash
go run . "3 + 4 * ( 2 - 1 )"
```

Example Output:

```bash
3 4 2 1 - * +           # Postfix
7                       # Result
```

### 🥞 Supported Operators
| Operator  | Description |
| ------------- | ------------- |
| +  | Addition  |
| -  | Subtraction  |
| *  | Multiplication  |
| /  | Integer Division  |
| ()  | Parenthesis Support  |