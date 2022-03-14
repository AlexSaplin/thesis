import string
import random


def generate_random_str(size=20, chars=string.ascii_lowercase + string.digits):
    return f"fn-{''.join(random.choice(chars) for _ in range(size))}"
