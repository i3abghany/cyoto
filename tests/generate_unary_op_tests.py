import random

def generate_expression(depth):
    if depth == 0:
        return str(random.randint(1, 10))  # Base case: a random integer

    op = random.choice(['+', '-'])
    sub_expr = generate_expression(depth - 1)
    return f"{op}{sub_expr}"

def generate_code(num_tests):
    code = ""
    for i in range(num_tests):
        expression = generate_expression(random.randint(1, 20))  # Adjust depth range as needed
        ret = eval(expression)
        code += f"""
// NAME unary_op_{i}
// RET {ret}

fn main() i32 {{
    return {expression};
}}

// END
"""
    return code

num_tests = 10
code = generate_code(num_tests)
print(code, file=open('inputs/arith_unary_ops.kyo', 'w+'))
