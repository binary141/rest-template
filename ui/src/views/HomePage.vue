<template>
  <div class="page">
    <div class="card">
      <h1>Welcome{{ store.user?.displayName ? ', ' + store.user.displayName : '' }}</h1>
      <p class="muted">{{ store.user?.email }}</p>
      <p class="token-label">Session token</p>
      <code class="token">{{ store.token }}</code>
      <button @click="handleLogout" :disabled="loading">
        {{ loading ? 'Signing out…' : 'Sign out' }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useUsersStore } from '@/stores/users';

const store = useUsersStore();
const router = useRouter();
const loading = ref(false);

async function handleLogout() {
  loading.value = true;
  await store.logOut();
  loading.value = false;
  router.push({ name: 'login' });
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

.muted {
  color: #888;
  font-size: 0.9rem;
}

.token-label {
  font-size: 0.75rem;
  color: #666;
  text-transform: uppercase;
  letter-spacing: 0.05em;
  margin-top: 0.5rem;
}

.token {
  display: block;
  background: #0d0d0d;
  border: 1px solid #2a2a2a;
  border-radius: 6px;
  padding: 0.6rem 0.75rem;
  font-size: 0.75rem;
  color: #aaa;
  word-break: break-all;
}

button {
  background: #2a2a2a;
  color: #f0f0f0;
  border: 1px solid #333;
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
</style>
