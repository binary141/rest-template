<template>
  <div class="app">
    <header class="header">
      <span class="app-name">example</span>
      <button v-if="store.loggedIn" class="user-btn" aria-label="User">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" width="22" height="22">
          <path d="M12 12c2.7 0 4.8-2.1 4.8-4.8S14.7 2.4 12 2.4 7.2 4.5 7.2 7.2 9.3 12 12 12zm0 2.4c-3.2 0-9.6 1.6-9.6 4.8v2.4h19.2v-2.4c0-3.2-6.4-4.8-9.6-4.8z"/>
        </svg>
      </button>
    </header>
    <main class="content">
      <RouterView />
    </main>
  </div>
</template>

<script setup>
import { useUsersStore } from '@/stores/users';

const store = useUsersStore();
</script>

<style scoped>
.app {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.header {
  height: 56px;
  padding: 0 1.5rem;
  background: #1a1a1a;
  border-bottom: 1px solid #2a2a2a;
  display: flex;
  align-items: center;
  flex-shrink: 0;
}

.app-name {
  font-size: 1rem;
  font-weight: 600;
  letter-spacing: 0.03em;
  color: #f0f0f0;
}

.user-btn {
  margin-left: auto;
  background: none;
  border: none;
  color: #aaa;
  cursor: pointer;
  padding: 0.25rem;
  display: flex;
  align-items: center;
  border-radius: 50%;
  transition: color 0.15s, background 0.15s;
}

.user-btn:hover {
  color: #f0f0f0;
  background: #2a2a2a;
}

.content {
  flex: 1;
  display: flex;
  flex-direction: column;
}
</style>
