import sys
import random
import string

# Exemplos adicionados para URLs e softwares
URLS = [
    "http://localhost:8080",
    "http://localhost:9090",
    "https://127.0.0.1",
    "https://google.com.br",
    "http://192.145.223.12",
    "http://112.120.233.232",
    "http://192.168.1.1",
    "https://tesla.com",
    "https://volkswagen.com",
    "https://bugcrowd.com",
    "https://tim.com.br",
    "https://nic.ru",
    "https://youtube.com",
    "https://google.com",
    "https://cloudfare.com",
    "https://gmail.com",
    "http://example.com",
    "http://somesite.com",
    "https://freemoney.com",
    "https://microsoft.com",
    "https://xtremecool.com"
]

SOFTWARES = [
    'Google Chrome', 
    'Microsoft Edge', 
    'Mozilla Firefox', 
    'Safari'
]

COMMON_EMAILS = [
    'gmail.com', 
    'protonmail.com', 
    'yahoo.com', 
    'outlook.com', 
    'hotmail.com',
    'proton.me', 
    'wtf.my',
    'yandex.ru',
    'outlook.com',
    'example.com'
]


def generate_random_string(length):
    return ''.join(random.choice(string.ascii_lowercase) for _ in range(length))


def generate_password(length=10):
    characters = string.ascii_letters + string.digits
    return ''.join(random.choice(characters) for _ in range(length))


def generate_stealer_type():
    software = random.choice(SOFTWARES)
    url = random.choice(URLS)
    user = f"{generate_random_string(10)}@example.com"
    password = generate_password()
    return f"SOFT: {software}\nURL:  {url}\nUSER: {user}\nPASS: {password}\n"


def generate_logs_type():
    url = random.choice(URLS)
    user = f"{generate_random_string(10)}@example.com"
    password = generate_password()
    return f'{url}:{user}:{password}'


def generate_mailpass_type():
    mail = f"{generate_random_string(10)}@{random.choice(COMMON_EMAILS)}"
    password = generate_password()
    return f"{mail}:{password}"


def generate_random_type(length=10):
    return generate_random_string(length)


if __name__ == "__main__":
    if len(sys.argv) != 4 or sys.argv[1] != "--number":
        print("Usage: python generate_random_leaks.py --number <num> <type>")
        sys.exit(1)

    try:
        num = int(sys.argv[2])
    except ValueError:
        print("Invalid number.")
        sys.exit(1)

    gen_type = sys.argv[3]

    generators = {
        "stealer": generate_stealer_type,
        "logs": generate_logs_type,
        "mailpass": generate_mailpass_type,
        "random": generate_random_type
    }

    if gen_type not in generators:
        print("Invalid type. Use 'stealer', 'logs', 'mailpass', or 'random'.")
        sys.exit(1)

    generator = generators[gen_type]
    for _ in range(num):
        random_data = generator()
        print(random_data)
