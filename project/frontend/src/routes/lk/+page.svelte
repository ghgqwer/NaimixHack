<script>
    let profile = {
      firstName: "Иван",
      lastName: "Петров",
      middleName: "Иванович",
      email: "ivan.petrov@example.com",
      birthDate: "1990-05-15",
      specialty: "Программист",
      birthTime: "08:30",
      birthCity: "Москва",
      workCity: "Санкт-Петербург",
      experience: "5 лет работы в IT-индустрии"
    };
  
    let isEditing = false;
  
    const editProfile = () => (isEditing = true);
    const saveProfile = () => (isEditing = false);
    const cancelEdit = () => {
      isEditing = false;
      profile = { ...profile };
    };
  </script>
  
  <style>
    :global(body) {
      background: linear-gradient(to bottom, #f0f4f8, #e8ecf1);
      font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
      margin: 0;
      padding: 0;
      color: #333;
    }
  
    .profile-card {
      max-width: 700px;
      margin: 40px auto;
      background: #ffffff;
      border-radius: 16px;
      padding: 30px;
      box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
    }
  
    .profile-card h2 {
      font-size: 28px;
      font-weight: 600;
      color: #2c3e50;
      text-align: center;
      margin-bottom: 20px;
    }
  
    .form-label {
      font-size: 14px;
      font-weight: bold;
      color: #546e7a;
    }
  
    .form-control {
      font-size: 16px;
      padding: 10px;
      border: 1px solid #ccd6dd;
      border-radius: 8px;
      background-color: #f9f9f9;
      transition: border-color 0.3s, box-shadow 0.3s;
    }
  
    .form-control:focus {
      border-color: #5c9ef8;
      box-shadow: 0 0 6px rgba(92, 158, 248, 0.5);
      background-color: #ffffff;
    }
  
    textarea.form-control {
      resize: none;
    }
  
    .btn {
      font-size: 16px;
      padding: 10px 20px;
      border-radius: 8px;
      cursor: pointer;
      transition: background-color 0.3s, transform 0.2s;
    }
  
    .btn-primary {
      background-color: #5c9ef8;
      color: white;
      border: none;
    }
  
    .btn-primary:hover {
      background-color: #407dcd;
      transform: translateY(-2px);
    }
  
    .btn-success {
      background-color: #2ecc71;
      color: white;
      border: none;
    }
  
    .btn-success:hover {
      background-color: #27ae60;
      transform: translateY(-2px);
    }
  
    .btn-secondary {
      background-color: #95a5a6;
      color: white;
      border: none;
    }
  
    .btn-secondary:hover {
      background-color: #7f8c8d;
      transform: translateY(-2px);
    }
  
    .btn-container {
      display: flex;
      justify-content: center;
      gap: 10px;
      margin-top: 20px;
    }
  </style>
  
  <div class="profile-card">
    <h2>Профиль пользователя</h2>
    <form class="row g-4">
      <!-- Имя -->
      <div class="col-md-4">
        <label for="firstName" class="form-label">Имя</label>
        <input
          id="firstName"
          type="text"
          class="form-control"
          bind:value={profile.firstName}
          disabled={!isEditing}
        />
      </div>
  
      <!-- Фамилия -->
      <div class="col-md-4">
        <label for="lastName" class="form-label">Фамилия</label>
        <input
          id="lastName"
          type="text"
          class="form-control"
          bind:value={profile.lastName}
          disabled={!isEditing}
        />
      </div>
  
      <!-- Отчество -->
      <div class="col-md-4">
        <label for="middleName" class="form-label">Отчество</label>
        <input
          id="middleName"
          type="text"
          class="form-control"
          bind:value={profile.middleName}
          disabled={!isEditing}
        />
      </div>
  
      <!-- Email -->
      <div class="col-md-6">
        <label for="email" class="form-label">Email</label>
        <input
          id="email"
          type="email"
          class="form-control"
          bind:value={profile.email}
          disabled={!isEditing}
        />
      </div>
  
      <!-- Дата рождения -->
      <div class="col-md-6">
        <label for="birthDate" class="form-label">Дата рождения</label>
        <input
          id="birthDate"
          type="date"
          class="form-control"
          bind:value={profile.birthDate}
          disabled={!isEditing}
        />
      </div>
  
      <!-- Специальность -->
      <div class="col-md-6">
        <label for="specialty" class="form-label">Специальность</label>
        <input
          id="specialty"
          type="text"
          class="form-control"
          bind:value={profile.specialty}
          disabled={!isEditing}
        />
      </div>
  
      <!-- Время рождения -->
      <div class="col-md-6">
        <label for="birthTime" class="form-label">Время рождения</label>
        <input
          id="birthTime"
          type="time"
          class="form-control"
          bind:value={profile.birthTime}
          disabled={!isEditing}
        />
      </div>
  
      <!-- Город рождения -->
      <div class="col-md-6">
        <label for="birthCity" class="form-label">Город рождения</label>
        <input
          id="birthCity"
          type="text"
          class="form-control"
          bind:value={profile.birthCity}
          disabled={!isEditing}
        />
      </div>
  
      <!-- Город работы -->
      <div class="col-md-6">
        <label for="workCity" class="form-label">Город работы</label>
        <input
          id="workCity"
          type="text"
          class="form-control"
          bind:value={profile.workCity}
          disabled={!isEditing}
        />
      </div>
  
      <!-- Опыт работы -->
      <div class="col-12">
        <label for="experience" class="form-label">Опыт работы</label>
        <textarea
          id="experience"
          class="form-control"
          rows="3"
          bind:value={profile.experience}
          disabled={!isEditing}
        ></textarea>
      </div>
  
      <!-- Кнопки -->
      <div class="btn-container">
        {#if !isEditing}
          <button type="button" class="btn btn-primary" on:click={editProfile}>
            Редактировать
          </button>
        {:else}
          <button type="button" class="btn btn-success" on:click={saveProfile}>
            Сохранить
          </button>
          <button type="button" class="btn btn-secondary" on:click={cancelEdit}>
            Отмена
          </button>
        {/if}
      </div>
    </form>
  </div>
  