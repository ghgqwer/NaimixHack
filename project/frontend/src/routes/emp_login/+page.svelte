<script lang="ts">
    import { goto } from "$app/navigation";
</script>


<style>
    .main {
        font-family: Arial, sans-serif;
        background-color: #f4f4f4;
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100vh;
        margin: 0;
    }
    .login-container {
        margin: 20px;
        background-color: #fff;
        padding: 20px;
        border-radius: 5px;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
        width: 300px;
    }
    h2 {
        text-align: center;
    }
    .form-group {
        margin-bottom: 15px;
    }
    label {
        display: block;
        margin-left: 20px; /* Внешние отступы слева */
        margin-right: 20px; /* Внешние отступы справа */
        margin-bottom: 5px;
    }
    input[type="email"],
    input[type="password"] {
        width: calc(100% - 40px); /* Учитываем отступы */
        padding: 10px; /* Внутренние отступы */
        margin-left: 20px; /* Внешние отступы слева */
        margin-right: 20px; /* Внешние отступы справа */
        border: 1px solid #ccc;
        border-radius: 4px;
        box-sizing: border-box; /* Учитываем ширину границ в общей ширине */width: calc(100% - 40px); /* Учитываем отступы */
    }
    button {
        width: calc(100% - 40px);
        padding: 10px;
        margin-left: 20px; 
        margin-right: 20px; 
        background-color: #FF9929;
        color: white;
        border: none;
        border-radius: 4px;
        cursor: pointer;
    }
    .register-link {
        text-align: center;
        margin-top: 10px;
    }
</style>
<title>Вход в систему</title>
<div class="main">
    <div class="login-container">
        <h2>Вход для рекрута</h2>
        <form action="/employee_authorization_page" method="POST">
            <div class="form-group">
                <label for="email">Электронная почта</label>
                <input type="email" id="email" name="email" required>
            </div>
            <div class="form-group">
                <label for="password">Пароль</label>
                <input type="password" id="password" name="password" required>
            </div>
            <button type="submit" on:click={() => goto("./main")}>Войти</button>
        </form>
        <div class="register-link">
            <p>Нет аккаунта? <a href="/emp_registr">Зарегистрироваться</a></p>
        </div>
    </div>
</div>

