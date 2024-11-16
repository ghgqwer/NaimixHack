<script lang="ts">
    	import { goto } from "$app/navigation";
</script>

<header class="header">
    <div class="container header__content">
        <img src="/naimix-black2.svg fill.svg" alt="Наймикс" on:click={() => goto('/main')} 
        style="cursor: pointer;" class="logo">
        <nav class="nav">
        <a href="#" class="nav__link">Для рекрутеров</a>
        <a href="#" class="nav__link">Для кандидатов</a>
        <a href="/lk" class="nav__link nav__link--button">Кабинет</a>
        </nav>
    </div>
</header>
<main>
    <slot /> <!-- Контент страницы будет вставлен сюда -->
</main>
<footer class="footer">
    <div class="container footer__content">
      <div class="footer__column">
        <h3 class="footer__title">Наймикс</h3>
        <p>Современные методы анализа совместимости для рекрутеров и команд.</p>
      </div>
      <div class="footer__column">
        <h3 class="footer__title">Исполнителям</h3>
        <ul class="footer__list">
          <li><a href="#" class="footer__link">Открытие</a></li>
          <li><a href="#" class="footer__link">Заявки</a></li>
          <li><a href="#" class="footer__link">Как мы работаем?</a></li>
          <li><a href="#" class="footer__link">Отзывы</a></li>
        </ul>
      </div>
      <div class="footer__column">
        <h3 class="footer__title">Контакты</h3>
        <p>125171, г. Москва, Ленинградское ш., д. 16, стр. 3</p>
        <p>8 800 500 56 54</p>
      </div>
    </div>
  </footer>
<style>
    * {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}
main {
  background-color: #0E6666;
}
body {
  font-family: Arial, sans-serif;
  line-height: 1.6;
  color: #333;
}

/* Utility Classes */
.container {
  width: 90%;
  max-width: 1200px;
  margin: 0 auto;
  justify-content: space-between;
}
.header {
  background-color: #fff;
  border-bottom: 2px solid #eaeaea;
  padding: 10px 0;
}

.header__content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  font-size: 1.5rem;
  font-weight: bold;
  color: #333;
}

.nav {
  display: flex;
  align-items: center;
  gap: 20px;
}

.nav__link {
  text-decoration: none;
  color: #333;
  font-weight: bold;
}

.nav__link--button {
  color: #FF9929;
  font-weight: bold;
  border: 2px solid #FF9929;
  padding: 5px 10px;
  border-radius: 5px;
}
/* Branding Section */
.branding {
  /* background-color: #FF9929; */
  color: #fff;
  text-align: center;
  padding: 20px 0;
}

.branding__text {
  font-size: 3rem;
  text-transform: uppercase;
}

/* Footer */
.footer {
  background-color: #f5f5f5;
  color: #333;
  padding: 20px 0;
}

.footer__content {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  gap: 20px;
}

.footer__column {
  flex: 1;
  min-width: 200px;
}

.footer__title {
  font-size: 1.2rem;
  margin-bottom: 10px;
}

.footer__list {
  list-style: none;
}

.footer__link {
  text-decoration: none;
  color: #333;
  display: block;
  margin: 5px 0;
}
</style>