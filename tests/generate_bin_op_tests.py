import random

def generate_random_expression(depth=3):
    if depth == 0:
        return str(random.randint(1, 10))

    op = random.choice(['+', '-', '*', '/'])
    left = generate_random_expression(depth-1)
    right = generate_random_expression(depth-1)

    if op == '/' and right == '0':
        right = str(random.randint(1, 10))

    return f"{left} {op} {right}"

def generate_test_cases(n):
    test_cases = []
    for i in range(n):
        expression = generate_random_expression(depth=random.randint(3, 6))
        result = eval(expression.replace('/', '//')) 
        result = result % (2**32)
        if result >= 2**31:
            result -= 2**32

        test_cases.append((expression, result))
    return test_cases

if __name__ == "__main__":
    n = 10
    test_cases = generate_test_cases(n)
    with open("inputs/arith_bin_ops.kyo", "w") as f:
        for i, (expression, result) in enumerate(test_cases):
            print(f"// NAME bin_op_{i}\n// RET {result}\n\nfn main() i32 {{\n\treturn {expression};\n}}\n\n", file=f)
            if i != n - 1:
                print("// END\n", file=f)
