<template>
  <div class="page">
    <div class="card">
      <RouterLink :to="{ name: 'home' }" class="back-link">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" width="14" height="14">
          <path d="M20 11H7.83l5.59-5.59L12 4l-8 8 8 8 1.41-1.41L7.83 13H20v-2z"/>
        </svg>
        Home
      </RouterLink>
      <h1>Account</h1>

      <section>
        <h2>Profile</h2>
        <form @submit.prevent="saveProfile">
          <label for="displayName">Display name</label>
          <input id="displayName" v-model="form.displayName" type="text" placeholder="Your name" />

          <label for="email">Email</label>
          <input id="email" v-model="form.email" type="email" placeholder="you@example.com" required />

          <div class="row">
            <p v-if="profileError" class="msg error">{{ profileError }}</p>
            <p v-else-if="profileSuccess" class="msg success">Saved.</p>
            <button type="submit" :disabled="profileSaving || !profileDirty">
              {{ profileSaving ? 'Saving…' : 'Save changes' }}
            </button>
          </div>
        </form>
      </section>

      <div class="divider" />

      <section>
        <h2>Details</h2>
        <div class="details">
          <span class="detail-label">User ID</span>
          <span class="detail-value">{{ store.user?.id }}</span>
          <span class="detail-label">Member since</span>
          <span class="detail-value">{{ joinedDate }}</span>
          <span class="detail-label">Role</span>
          <span class="detail-value">{{ store.user?.isAdmin ? 'Admin' : 'User' }}</span>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import { useUsersStore } from '@/stores/users';

const store = useUsersStore();

const form = ref({
  displayName: store.user?.displayName ?? '',
  email: store.user?.email ?? '',
});

const profileSaving = ref(false);
const profileError = ref('');
const profileSuccess = ref(false);

const profileDirty = computed(() =>
  form.value.displayName !== (store.user?.displayName ?? '') ||
  form.value.email !== (store.user?.email ?? '')
);

const joinedDate = computed(() => {
  if (!store.user?.createdAt) return '—';
  return new Date(store.user.createdAt).toLocaleDateString(undefined, {
    year: 'numeric', month: 'long', day: 'numeric',
  });
});

async function saveProfile() {
  profileError.value = '';
  profileSuccess.value = false;
  profileSaving.value = true;

  const response = await store.updateUser({
    email: form.value.email,
    displayName: form.value.displayName,
  });

  profileSaving.value = false;

  if (response.ok) {
    profileSuccess.value = true;
  } else {
    profileError.value = response.body?.error ?? 'Failed to save changes';
  }
}
</script>

<style scoped>
.page {
  display: flex;
  justify-content: center;
  flex: 1;
  padding: 2rem 1rem;
}

.card {
  background: #1a1a1a;
  border: 1px solid #2a2a2a;
  border-radius: 10px;
  padding: 2rem;
  width: 100%;
  max-width: 480px;
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  height: fit-content;
}

.back-link {
  display: inline-flex;
  align-items: center;
  gap: 0.35rem;
  font-size: 0.85rem;
  color: #888;
  text-decoration: none;
  transition: color 0.15s;
}

.back-link:hover {
  color: #f0f0f0;
}

h1 {
  font-size: 1.5rem;
  font-weight: 600;
}

h2 {
  font-size: 0.8rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.06em;
  color: #666;
  margin-bottom: 0.75rem;
}

section {
  display: flex;
  flex-direction: column;
}

form {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

label {
  font-size: 0.85rem;
  color: #aaa;
  margin-top: 0.5rem;
}

input {
  background: #0d0d0d;
  border: 1px solid #333;
  border-radius: 6px;
  color: #f0f0f0;
  font-size: 0.95rem;
  padding: 0.6rem 0.75rem;
  outline: none;
  transition: border-color 0.15s;
}

input:focus {
  border-color: #555;
}

.row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 0.75rem;
  gap: 0.5rem;
}

button {
  margin-left: auto;
  background: #f0f0f0;
  color: #0d0d0d;
  border: none;
  border-radius: 6px;
  font-size: 0.9rem;
  font-weight: 600;
  padding: 0.55rem 1rem;
  cursor: pointer;
  transition: opacity 0.15s;
  white-space: nowrap;
}

button:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.msg {
  font-size: 0.85rem;
}

.error {
  color: #f87171;
}

.success {
  color: #4ade80;
}

.divider {
  height: 1px;
  background: #2a2a2a;
}

.details {
  display: grid;
  grid-template-columns: auto 1fr;
  gap: 0.4rem 1.5rem;
  align-items: baseline;
}

.detail-label {
  font-size: 0.85rem;
  color: #666;
}

.detail-value {
  font-size: 0.9rem;
  color: #ccc;
}
</style>
