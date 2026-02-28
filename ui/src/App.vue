<template>
  <div class="app">
    <header class="header">
      <RouterLink :to="{ name: 'home' }" class="app-name">example</RouterLink>

      <div v-if="store.loggedIn" class="user-menu" ref="menuRef">
        <button class="user-btn" aria-label="User menu" @click="open = !open">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" width="22" height="22">
            <path d="M12 12c2.7 0 4.8-2.1 4.8-4.8S14.7 2.4 12 2.4 7.2 4.5 7.2 7.2 9.3 12 12 12zm0 2.4c-3.2 0-9.6 1.6-9.6 4.8v2.4h19.2v-2.4c0-3.2-6.4-4.8-9.6-4.8z"/>
          </svg>
        </button>

        <div v-if="open" class="dropdown">
          <div class="dropdown-header">
            <p class="dropdown-name">{{ store.user?.displayName || 'Account' }}</p>
            <p class="dropdown-email">{{ store.user?.email }}</p>
          </div>
          <div class="dropdown-divider" />
          <button class="dropdown-item" @click="goToAccount">Account</button>
          <div class="dropdown-divider" />
          <button class="dropdown-item danger" @click="handleLogout">Sign out</button>
        </div>
      </div>
    </header>

    <main class="content">
      <RouterView />
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';
import { useUsersStore } from '@/stores/users';

const store = useUsersStore();
const router = useRouter();
const open = ref(false);
const menuRef = ref(null);

function handleOutsideClick(e) {
  if (menuRef.value && !menuRef.value.contains(e.target)) {
    open.value = false;
  }
}

onMounted(() => document.addEventListener('mousedown', handleOutsideClick));
onUnmounted(() => document.removeEventListener('mousedown', handleOutsideClick));

function goToAccount() {
  open.value = false;
  router.push({ name: 'account' });
}

async function handleLogout() {
  open.value = false;
  await store.logOut();
  router.push({ name: 'login' });
}
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
  text-decoration: none;
}

.user-menu {
  margin-left: auto;
  position: relative;
}

.user-btn {
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

.dropdown {
  position: absolute;
  top: calc(100% + 0.5rem);
  right: 0;
  background: #1a1a1a;
  border: 1px solid #2a2a2a;
  border-radius: 8px;
  min-width: 200px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.4);
  overflow: hidden;
  z-index: 100;
}

.dropdown-header {
  padding: 0.75rem 1rem;
  display: flex;
  flex-direction: column;
  gap: 0.2rem;
}

.dropdown-name {
  font-size: 0.9rem;
  font-weight: 600;
  color: #f0f0f0;
}

.dropdown-email {
  font-size: 0.8rem;
  color: #888;
}

.dropdown-divider {
  height: 1px;
  background: #2a2a2a;
}

.dropdown-item {
  display: block;
  width: 100%;
  padding: 0.65rem 1rem;
  background: none;
  border: none;
  color: #d0d0d0;
  font-size: 0.9rem;
  text-align: left;
  cursor: pointer;
  transition: background 0.15s;
}

.dropdown-item:hover {
  background: #2a2a2a;
}

.dropdown-item.danger {
  color: #f87171;
}

.content {
  flex: 1;
  display: flex;
  flex-direction: column;
}
</style>
