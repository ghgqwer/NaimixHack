<script>

	import { goto } from "$app/navigation";

</script>

  
<div class="main">
    <form class="registration-form">
        <h2 class="name">Регистрация рекрута</h2>
        
        <div class="form-group">
          <label for="name">Имя</label>
          <input type="text" id="name" name="name" required placeholder="Введите имя">
        </div>
    
        <div class="form-group">
          <label for="surname">Фамилия</label>
          <input type="text" id="surname" name="surname" required placeholder="Введите фамилию">
        </div>
    
        <div class="form-group">
          <label for="secondname">Отчество</label>
          <input type="text" id="secondname" name="secondname" required placeholder="Введите отчество">
        </div>
    
        <div class="form-group">
          <label for="email">Email</label>
          <input type="email" id="email" name="email" required placeholder="Введите email">
        </div>
    
        <div class="form-group">
          <label for="password">Пароль</label>
          <input type="password" id="password" name="password" required placeholder="Введите пароль">
        </div>
    
        <div class="form-group">
          <label for="Date_birth_day">Дата рождения</label>
          <input type="date" id="Date_birth_day" name="Date_birth_day" required>
        </div>

        <div class="form-group">
          <label for="Time_birth">Время рождения</label>
          <input type="time" id="Time_birth" name="Time_birth" required>
        </div>
    
        <div class="form-group">
          <label for="City_of_birth">Город рождения</label>
          <input type="text" id="City_of_birth" name="City_of_birth" required placeholder="Введите город">
        </div>
    
        <div class="form-group">
          <label for="dropdown2">Специализация</label>
          <select id="dropdown2" name="dropdown2" required>
            <option value="">Выберите специализацию</option>
            <option value="developer">Разработчик</option>
            <option value="designer">Дизайнер</option>
            <option value="manager">Менеджер</option>
            <option value="analyst">Аналитик</option>
          </select>
        </div>
    
        <div class="form-group">
          <label for="experience">Опыт работы</label>
          <input type="number" id="Work_experience" name="Work_experience" required min="0" placeholder="Введите количество лет">
          <small>Укажите 0, если у вас нет опыта.</small>
        </div>
    
        <button type="submit" class="submit-button" on:click={() => goto("./main")}>Зарегистрироваться</button>
      </form>
</div>
  


<style>
    body {
      font-family: Arial, sans-serif;
      background-color: #f4f4f4;
      margin: 0;
      padding: 0;
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100vh;
    }


    .main{
      margin-top: 20px;
      margin-bottom: 20px;

        display: flex;
        justify-content: center;
    }

    .registration-form {
    font-size: 15px;
      background-color: #fff;
      padding: 20px;
      border-radius: 8px;
      box-shadow:  6px 6px 6px 6px rgba(0, 0, 0, 0.1);
      width: 400px;
    }

    .registration-form h2 {
      text-align: center;
      margin-bottom: 20px;
    }

    .form-group {
      margin-bottom: 15px;
    }

    .form-group label {
      display: block;
      margin-bottom: 5px;
      font-weight: bold;
    }

    .form-group input,
    .form-group select {
      width: 100%;
      padding: 10px;
      border: 1px solid #ddd;
      border-radius: 4px;
      box-sizing: border-box;
    }

    .form-group input:focus,
    .form-group select:focus {
      border-color:#0E6666;
      outline: none;
    }

    .form-group small {
      color: #666;
      font-size: 12px;
    }

    .submit-button {
      display: block;
      width: 100%;
      padding: 10px;
      background-color: #FF9929;
      color: white;
      border: none;
      border-radius: 4px;
      cursor: pointer;
      font-size: 16px;
    }

    .name {
        margin: 0;
    }
  </style>