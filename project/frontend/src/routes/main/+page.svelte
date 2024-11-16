<script>

	import { goto } from "$app/navigation";

</script>


<body>
  <main>
    <section class="hero">
      <div class="container hero__content">
        <div class="hero__text">
          <h2 class="hero__title">
            Сервис, где рекрутеры могут использовать новые методы для подбора сотрудников
          </h2>
          <p class="hero__subtitle">Использование эзотерики для определения совместимости сотрудников</p>
        </div>
        <div class="hero__image">
          <img src="/image 2.svg" alt="Собеседование">
        </div>
      </div>
    </section>

    <section class="branding">
      <img src="/naimix-text-orang.svg.svg" alt="Наймикс" class="logo">
    </section>
  </main>
</body>

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

/* Hero Section */
.hero {
  background-color: #0E6666;
  color: #fff;
  padding: 50px 0;
}

.hero__content {
  display: flex;
  align-items: center;
  gap: 20px;
}

.hero__text {
  max-width: 600px;
}

.hero__title {
  font-size: 2.5rem;
  margin-bottom: 10px;
}

.hero__subtitle {
  font-size: 1.2rem;
}

.hero__image img {
  width: 400px;
  height: auto;
  border-radius: 10px;
}

/* Branding Section */
.branding {
  /* background-color: #FF9929; */
  color: #fff;
  text-align: center;
  padding: 20px 0;
}
</style>