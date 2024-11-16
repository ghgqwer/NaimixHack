import requests
import json
import time
import base64
import os
import argparse

api_key = "5867D08E3C13ED7770FFE0C88CFEF03F"
secret_key = "9A760F7EBD2A1D7753DBE5C6C1654FCF"

class Text2ImageAPI:
    def __init__(self, url, api_key, secret_key):
        self.URL = url
        self.AUTH_HEADERS = {
            'X-Key': f'Key {api_key}',
            'X-Secret': f'Secret {secret_key}',
        }

    def get_model(self):
        response = requests.get(self.URL + 'key/api/v1/models', headers=self.AUTH_HEADERS)
        data = response.json()
        return data[0]['id']

    def generate(self, prompt, model, images=1, width=1024, height=1024):
        params = {
            "type": "GENERATE",
            "numImages": images,
            "width": width,
            "height": height,
            "generateParams": {
                "query": f"{prompt}"
            }
        }

        data = {
            'model_id': (None, model),
            'params': (None, json.dumps(params), 'application/json')
        }
        response = requests.post(self.URL + 'key/api/v1/text2image/run', headers=self.AUTH_HEADERS, files=data)
        data = response.json()
        return data['uuid']

    def check_generation(self, request_id, attempts=10, delay=10):
        while attempts > 0:
            response = requests.get(self.URL + 'key/api/v1/text2image/status/' + request_id, headers=self.AUTH_HEADERS)
            data = response.json()
            if data['status'] == 'DONE':
                return data['images']

            attempts -= 1
            time.sleep(delay)

def gen(prom, surname, name, dirr="users"):
    api = Text2ImageAPI('https://api-key.fusionbrain.ai/', api_key, secret_key)
    model_id = api.get_model()
    uuid = api.generate(prom, model_id)
    images = api.check_generation(uuid)

    # Здесь image_base64 - это строка с данными изображения в формате base64
    image_base64 = images[0]

    # Декодируем строку base64 в бинарные данные
    image_data = base64.b64decode(image_base64)

    # Создаем имя файла в формате surname_name_secondname.jpg
    filename = os.path.join(dirr, "icon.jpg")

    # Открываем файл для записи бинарных данных изображения
    with open(filename, "wb") as file:
        file.write(image_data)

def main(prompt, id, surname, name):
    # Получаем путь к директории, где находится скрипт
    script_dir = os.path.dirname(os.path.abspath(__file__))
    users_dir = os.path.join(script_dir, "users")

    # Создаем директорию для сохранения изображений
    os.makedirs(users_dir, exist_ok=True)  # Создает директорию "users", если она не существует

    # Создаем поддиректорию с именем id
    user_dir = os.path.join(users_dir, id)
    os.makedirs(user_dir, exist_ok=True)  # Создает директорию с именем id, если она не существует

    gen(prompt, surname, name, user_dir)
    print(f"Изображение сохранено как {surname}_{name}.jpg в папке {user_dir}")
    print("Завершено")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Generate images from text prompts.")
    parser.add_argument("prompt", type=str, help="Text prompt for image generation")
    parser.add_argument("surname", type=str, help="User  's surname")
    parser.add_argument("name", type=str, help="User  's name")
    parser.add_argument("secondname", type=str, help="User  's second name")
    args = parser.parse_args()

    main(args.prompt, args.surname, args.name, args.secondname)
