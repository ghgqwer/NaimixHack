<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Просмотр и редактирование профиля</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 0;
        }
        .main {

            width: 80%;
            margin: 0 auto;
            padding: 20px;
            background-color: #fff;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            border-radius: 8px;
            margin-top: 50px;
        }
        h1 {
            text-align: center;
            margin-bottom: 20px;
        }
        .form-group {
            margin-bottom: 15px;
        }
        .form-group label {
            display: block;
            font-weight: bold;
            margin-bottom: 5px;
        }
        .form-group input {
            width: 100%;
            padding: 8px;
            border: 1px solid #ccc;
            border-radius: 5px;
            font-size: 14px;
        }
        .form-group input[type="date"],
        .form-group input[type="time"] {
            width: 48%;
            display: inline-block;
        }
        .form-group button {
            background-color: #007BFF;
            color: white;
            border: none;
            padding: 10px 20px;
            border-radius: 5px;
            cursor: pointer;
        }
        .form-group button:hover {
            background-color: #0056b3;
        }
        .exit-btn {
            background-color: #dc3545;
            margin-top: 20px;
            padding: 10px 20px;
            border-radius: 5px;
            color: white;
            text-align: center;
            display: block;
            width: 100%;
            text-decoration: none;
        }
        .exit-btn:hover {
            background-color: #c82333;
        }
        .view-mode input {
            background-color: #f4f4f4;
            cursor: not-allowed;
            border: none;
        }
        .edit-mode input {
            cursor: text;
        }
    </style>
</head>
<body>

<div class="main">
    <h1>Просмотр и редактирование профиля</h1>
    
    <!-- Условное отображение профиля (кандидат или компания) -->
    <form id="profileForm">
        <!-- Профиль кандидата -->
        <div id="candidateProfile" class="profileFields view-mode">
            <div class="form-group">
                <label for="firstName">Имя:</label>
                <input type="text" id="firstName" name="firstName" value="Иван" readonly>
            </div>
            <div class="form-group">
                <label for="lastName">Фамилия:</label>
                <input type="text" id="lastName" name="lastName" value="Иванов" readonly>
            </div>
            <div class="form-group">
                <label for="patronymic">Отчество:</label>
                <input type="text" id="patronymic" name="patronymic" value="Иванович" readonly>
            </div>
            <div class="form-group">
                <label for="email">Почта:</label>
                <input type="email" id="email" name="email" value="ivanov@mail.ru" readonly>
            </div>
            <div class="form-group">
                <label for="birthDate">Дата рождения:</label>
                <input type="date" id="birthDate" name="birthDate" value="1990-05-15" readonly>
            </div>
            <div class="form-group">
                <label for="birthTime">Время рождения:</label>
                <input type="time" id="birthTime" name="birthTime" value="08:30" readonly>
            </div>
            <div class="form-group">
                <label for="birthPlace">Город рождения:</label>
                <input type="text" id="birthPlace" name="birthPlace" value="Москва" readonly>
            </div>
            <div class="form-group">
                <label for="specialization">Специализация:</label>
                <input type="text" id="specialization" name="specialization" value="Разработчик" readonly>
            </div>
            <div class="form-group">
                <label for="workExperience">Опыт работы (в годах):</label>
                <input type="number" id="workExperience" name="workExperience" value="5" readonly>
            </div>
        </div>

        <!-- Профиль компании -->
        <div id="companyProfile" class="profileFields view-mode" style="display:none;">
            <div class="form-group">
                <label for="companyName">Название компании:</label>
                <input type="text" id="companyName" name="companyName" value="ООО 'Технополис'" readonly>
            </div>
            <div class="form-group">
                <label for="companyEmail">Почта:</label>
                <input type="email" id="companyEmail" name="companyEmail" value="info@techpolis.ru" readonly>
            </div>
        </div>

        <div class="form-group">
            <button type="button" id="editButton">Редактировать</button>
            <button type="submit" id="saveButton" style="display: none;">Сохранить изменения</button>
        </div>
    </form>
    
    <!-- Кнопка выхода -->
    <a href="/logout" class="exit-btn">Выйти из профиля</a>
</div>

<script>
    const isCandidateProfile = true; // Установить в true для кандидата, в false для компании

    if (isCandidateProfile) {
        document.getElementById('candidateProfile').style.display = 'block';
        document.getElementById('companyProfile').style.display = 'none';
    } else {
        document.getElementById('candidateProfile').style.display = 'none';
        document.getElementById('companyProfile').style.display = 'block';
    }

    // Переключение между режимами "просмотр" и "редактирование"
    document.getElementById('editButton').addEventListener('click', function() {
        const profileFields = document.querySelectorAll('.profileFields input');
        
        // Переключаем все поля на редактируемые
        document.getElementById('candidateProfile').classList.remove('view-mode');
        document.getElementById('candidateProfile').classList.add('edit-mode');
        document.getElementById('companyProfile').classList.remove('view-mode');
        document.getElementById('companyProfile').classList.add('edit-mode');
        
        profileFields.forEach(input => {
            input.removeAttribute('readonly');
        });

        // Показываем кнопку "Сохранить изменения" и скрываем кнопку "Редактировать"
        document.getElementById('saveButton').style.display = 'inline-block';
        document.getElementById('editButton').style.display = 'none';
    });

    // Обработчик отправки формы (сохранение изменений)
    document.getElementById('profileForm').addEventListener('submit', function(event) {
        event.preventDefault();
        
        // После сохранения, восстанавливаем режим "просмотра"
        const profileFields = document.querySelectorAll('.profileFields input');
        document.getElementById('candidateProfile').classList.remove('edit-mode');
        document.getElementById('candidateProfile').classList.add('view-mode');
        document.getElementById('companyProfile').classList.remove('edit-mode');
        document.getElementById('companyProfile').classList.add('view-mode');
        
        // Устанавливаем все поля в режим "только для чтения"
        profileFields.forEach(input => {
            input.setAttribute('readonly', true);
        });

        // Показываем кнопку "Редактировать" и скрываем кнопку "Сохранить"
        document.getElementById('saveButton').style.display = 'none';
        document.getElementById('editButton').style.display = 'inline-block';

        alert('Изменения сохранены!');
    });
</script>

</body>
</html>
