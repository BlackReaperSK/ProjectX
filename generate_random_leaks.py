import sys
import random
import string

def random_string(length):
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for _ in range(length))

def generate_password(length=10):
    characters = string.ascii_letters + string.digits
    return ''.join(random.choice(characters) for _ in range(length))

def generate_stealer_type():
    software = random.choice(['Google Chrome', 'Microsoft Edge', 'Mozilla Firefox', 'Safari'])
    urls = [
        "http://localhost:8080",
        "http://localhost:9090",
        "https://127.0.0.1",
        "https://google.com.br",
        "http:192.145.223.12",
        "http:112.120.233.232",
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
    url = random.choice(urls)
    user = f"{random_string(10)}@example.com"
    password = generate_password()
    return f"SOFT: {software}\nURL:  {url}\nUSER: {user}\nPASS: {password}\n"

def generate_logs_type():
    logs = [
        "INFO: This is an informational message",
        "ERROR: Something went wrong",
        "WARNING: Be cautious!",
        "DEBUG: Detailed debugging information"
    ]
    return random.choice(logs)

def generate_mailpass_type():
    common_emails = ['gmail.com', 'protonmail.com', 'yahoo.com', 'outlook.com', 'hotmail.com','proton.me', 'wtf.my','yandex.ru','outlook.com','example.com']
    mail = f"{random_string(10)}@{random.choice(common_emails)}"
    user = f"{mail}"
    password = generate_password()
    return f"{user}:{password}"

def generate_random_type(length=10):
    return ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))

if __name__ == "__main__":
    if len(sys.argv) < 3 or sys.argv[1] != "--number":
        print("Usage: python generate_random_leaks.py --number <num> <type>")
        sys.exit(1)

    try:
        num = int(sys.argv[2])
    except ValueError:
        print("Invalid number.")
        sys.exit(1)

    gen_type = sys.argv[3]

    if gen_type == "stealer":
        for _ in range(num):
            random_data = generate_stealer_type()
            print(random_data)
    elif gen_type == "logs":
        for _ in range(num):
            random_data = generate_logs_type()
            print(random_data)
    elif gen_type == "mailpass":
        for _ in range(num):
            random_data = generate_mailpass_type()
            print(random_data)
    elif gen_type == "random":
        for _ in range(num):
            random_data = generate_random_type()
            print(random_data)
    else:
        print("Invalid type. Use 'stealer', 'logs', 'mailpass', or 'random'.")
