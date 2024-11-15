<script>

	import { goto } from "$app/navigation";

</script>
<div class="main">
    <header>
        <!-- Логотип -->
        <div class="logo">
          <div class="logo-symbol">•</div> <!-- Символ перед логотипом -->
          <div class="logo-text">
            
          </div>
        </div>
      
        <!-- Навигация -->
        <nav>
          <a href="#">ДЛЯ СОИСКАТЕЛЯ</a>
          <a href="#">ДЛЯ БИЗНЕСА</a>
          <a href="#">КОНТАКТЫ</a>
          <a href="#">ПРОДУКТЫ</a>
        </nav>
      
        <!-- Кнопки -->
        <div class="buttons">
          <button class="find-order">НАЙТИ ЗАКАЗ</button>
          <button class="cabinet" on:click={() => goto("../lk")}>КАБИНЕТ</button>
        </div>
      </header>
      <div class="body">

        ;flasf;
      </div>
</div>





<style>
    .body {
      background-color: #0E6666;
      margin: 0;
      font-family: Arial, sans-serif;
    }

    /* Стили для header */
    header {
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 10px 20px;
      border-bottom: 1px solid #ddd;
    }

    /* Стили для логотипа */
    .logo {
      display: flex;
      align-items: center;
      font-size: 24px;
      font-weight: bold;
      color: black;
    }

    .logo-text {
      margin-left: 8px;
    }

    .logo-subtext {
      font-size: 12px;
      color: #555;
    }

    /* Стили для навигации */
    nav {
      display: flex;
      align-items: center;
      gap: 20px;
      font-size: 16px;
      color: #006666;
    }

    nav a {
      color: #006666;
      text-decoration: none;
      font-weight: 500;
    }

    /* Стили для кнопок */
    .buttons {
      display: flex;
      gap: 10px;
    }

    .find-order {
      padding: 8px 16px;
      border: 2px solid #ff6600;
      border-radius: 20px;
      color: black;
      font-weight: bold;
      background: transparent;
      cursor: pointer;
      transition: background-color 0.3s ease;
    }

    .find-order:hover {
      background-color: #ff6600;
      color: white;
    }

    .cabinet {
      padding: 8px 16px;
      border: 1px solid #ddd;
      border-radius: 20px;
      color: #ff6600;
      font-weight: bold;
      background: transparent;
      cursor: pointer;
      transition: background-color 0.3s ease;
    }

    .cabinet:hover {
      background-color: #f9f9f9;
    }
  </style>
