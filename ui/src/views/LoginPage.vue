<template>
  <div class="page">
    <div class="card">
      <h1>Sign in</h1>
      <form @submit.prevent="handleLogin">
        <label for="email">Email</label>
        <input
          id="email"
          v-model="email"
          type="email"
          placeholder="you@example.com"
          autocomplete="email"
          required
        />
        <label for="password">Password</label>
        <input
          id="password"
          v-model="password"
          type="password"
          placeholder="••••••••"
          autocomplete="current-password"
          required
        />
        <p v-if="error" class="error">{{ error }}</p>
        <button type="submit" :disabled="submitting">
          {{ submitting ? 'Signing in…' : 'Sign in' }}
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useUsersStore } from '@/stores/users';

const router = useRouter();
const route = useRoute();
const store = useUsersStore();

const email = ref('');
const password = ref('');
const error = ref('');
const submitting = ref(false);

async function handleLogin() {
  error.value = '';
  submitting.value = true;

  const response = await store.logInAPI({ email: email.value, password: password.value });
  submitting.value = false;

  if (response.ok) {
    store.setToken(response.body.token);
    store.logIn(response.body.user);
    router.replace(route.query.next || { name: 'home' });
  } else {
    error.value = response.body?.error ?? 'Login failed';
  }
}
</script>

<style scoped>
.page {
  display: flex;
  align-items: center;
  justify-content: center;
  flex: 1;
  padding: 1rem;
}

.card {
  background: #1a1a1a;
  border: 1px solid #2a2a2a;
  border-radius: 10px;
  padding: 2rem;
  width: 100%;
  max-width: 380px;
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

h1 {
  font-size: 1.5rem;
  font-weight: 600;
}

form {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
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

button {
  margin-top: 0.75rem;
  background: #f0f0f0;
  color: #0d0d0d;
  border: none;
  border-radius: 6px;
  font-size: 0.95rem;
  font-weight: 600;
  padding: 0.65rem;
  cursor: pointer;
  transition: opacity 0.15s;
}

button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.error {
  color: #f87171;
  font-size: 0.85rem;
}
</style>
